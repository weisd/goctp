package goctp

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Event int

const (
	onFrontConnected Event = iota
	onRspAuthenticate
	onRspUserLogin
	onRspSettlementInfoConfirm
	onRspQryInvestor
	onRspQryClassifiedInstrument
	onRspQryOrder
	onRspQryTrade
	onRspQryInvestorPosition
	onRspQryInvestorPositionDetail
	onRspQryTradingAccount
	onRspQryAccountregister
	onRspQryTransferBank
)

// TradePro 业务逻辑封装
type TradePro struct {
	*TradeExt

	OnOrder func(pOrder *CThostFtdcOrderField)
	OnTrade func(pTrade *CThostFtdcTradeField)

	// 合约 key: InstrumentID
	Instruments map[string]CThostFtdcInstrumentField
	// 委托 key: OrderLocalID
	Orders map[string]CThostFtdcOrderField
	// 成交 key: OrderLocalID values: []TradeField
	Trades map[string][]CThostFtdcTradeField
	// 投资者 key:InvestorID
	Investors map[string]CThostFtdcInvestorField
	// 持仓 查询时返回
	positions []CThostFtdcInvestorPositionField
	// 持仓明细 查询时返回
	positionDetails []CThostFtdcInvestorPositionDetailField
	// 权益 查询时返回
	accounts map[string]CThostFtdcTradingAccountField
	// 银行开户信息
	AccountRegisters map[string]CThostFtdcAccountregisterField

	// 响应事件
	eventChan chan Event
	// 错误
	errorChan chan CThostFtdcRspInfoField

	// 委托响应 本地报单编号
	orderChan    chan TThostFtdcOrderLocalIDType
	orderErrChan chan CThostFtdcRspInfoField

	// 银转
	inoutChan chan CThostFtdcRspInfoField

	// 用于判断是否此连接的委托
	sessionID TThostFtdcSessionIDType
}

func NewTradePro() *TradePro {
	trd := TradePro{}
	trd.TradeExt = NewTradeExt()

	// 查询相关 chan
	trd.eventChan = make(chan Event)
	trd.errorChan = make(chan CThostFtdcRspInfoField)

	// 银转相关 chan
	trd.inoutChan = make(chan CThostFtdcRspInfoField)

	// 委托相关 chan
	trd.orderChan = make(chan TThostFtdcOrderLocalIDType)
	trd.orderErrChan = make(chan CThostFtdcRspInfoField)

	// 登录过程中查询的信息
	trd.Instruments = make(map[string]CThostFtdcInstrumentField)
	trd.Investors = make(map[string]CThostFtdcInvestorField)
	trd.Orders = make(map[string]CThostFtdcOrderField)
	trd.Trades = make(map[string][]CThostFtdcTradeField)
	trd.AccountRegisters = make(map[string]CThostFtdcAccountregisterField)

	// 用户主动查询得到的数据
	trd.accounts = make(map[string]CThostFtdcTradingAccountField)
	trd.positionDetails = make([]CThostFtdcInvestorPositionDetailField, 0)
	trd.positions = make([]CThostFtdcInvestorPositionField, 0)

	// 持仓
	trd.Trade.OnRspQryInvestorPosition = func(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestorPosition != nil {
			trd.positions = append(trd.positions, *pInvestorPosition)
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryInvestorPosition
			}
		}
	}
	// 持仓明细
	trd.Trade.OnRspQryInvestorPositionDetail = func(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestorPositionDetail != nil {
			trd.positionDetails = append(trd.positionDetails, *pInvestorPositionDetail)
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryInvestorPositionDetail
			}
		}
	}
	// 权益
	trd.Trade.OnRspQryTradingAccount = func(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pTradingAccount != nil {
			trd.accounts[pTradingAccount.AccountID.String()] = *pTradingAccount
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryTradingAccount
			}
		}
	}

	// 委托
	trd.Trade.OnRspOrderInsert = func(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		trd.orderErrChan <- *pRspInfo
	}
	trd.Trade.OnRtnOrder = func(pOrder *CThostFtdcOrderField) {
		if pOrder.SessionID == trd.sessionID { // 此连接的委托
			_, ok := trd.Orders[pOrder.OrderLocalID.String()]
			trd.Orders[pOrder.OrderLocalID.String()] = *pOrder
			if !ok { // 首次响应
				trd.orderChan <- pOrder.OrderLocalID
			} else {
				if trd.OnOrder != nil {
					trd.OnOrder(pOrder)
				}
			}
		} else { // 非本连接响应
			if trd.OnOrder != nil {
				trd.OnOrder(pOrder)
			}
		}
	}
	// 成交
	trd.Trade.OnRtnTrade = func(pTrade *CThostFtdcTradeField) {
		if _, ok := trd.Trades[pTrade.OrderLocalID.String()]; ok {
			trd.Trades[pTrade.OrderLocalID.String()] = append(trd.Trades[pTrade.OrderLocalID.String()], *pTrade)
		} else {
			trd.Trades[pTrade.OrderLocalID.String()] = []CThostFtdcTradeField{*pTrade}
		}
		if trd.OnTrade != nil {
			trd.OnTrade(pTrade)
		}
	}

	// 银转:入金
	trd.Trade.OnRspFromBankToFutureByFuture = func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			trd.inoutChan <- *pRspInfo
		}
	}
	trd.Trade.OnRtnFromBankToFutureByFuture = func(pRspTransfer *CThostFtdcRspTransferField) {
		rsp := CThostFtdcRspInfoField{
			ErrorID: pRspTransfer.ErrorID,
		}
		copy(rsp.ErrorMsg[:], pRspTransfer.ErrorMsg[:])
		trd.inoutChan <- rsp
	}
	// 银转:出金
	trd.Trade.OnRspFromFutureToBankByFuture = func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			trd.inoutChan <- *pRspInfo
		}
	}
	trd.Trade.OnRtnFromFutureToBankByFuture = func(pRspTransfer *CThostFtdcRspTransferField) {
		rsp := CThostFtdcRspInfoField{
			ErrorID: pRspTransfer.ErrorID,
		}
		copy(rsp.ErrorMsg[:], pRspTransfer.ErrorMsg[:])
		trd.inoutChan <- rsp
	}
	return &trd
}

// 行情登录不需要 AppID AuthCode
type LoginConfig struct {
	Front, Broker, UserID, Password, AppID, AuthCode string
}

// Start 接口启动/登录/查询客户基础信息/查询委托/成交/权益
func (trd *TradePro) Start(cfg LoginConfig) (loginInfo CThostFtdcRspUserLoginField, rsp CThostFtdcRspInfoField) {
	trd.Trade.OnFrontConnected = func() {
		trd.eventChan <- onFrontConnected
	}
	trd.Trade.OnRspAuthenticate = func(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			if pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspAuthenticate
			}
		}
	}
	trd.Trade.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			if pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.sessionID = pRspUserLogin.SessionID
				loginInfo = *pRspUserLogin
				trd.eventChan <- onRspUserLogin
			}
		}
	}
	trd.Trade.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			// 交易员无确认结算权限,此处忽略
			// if pRspInfo.ErrorID != 0 {
			// 	trd.errorChan <- *pRspInfo
			// } else {
			trd.eventChan <- onRspSettlementInfoConfirm
			// }
		}
	}
	trd.Trade.OnRspQryInvestor = func(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestor != nil {
			trd.Investors[pInvestor.InvestorID.String()] = *pInvestor
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryInvestor
			}
		}
	}
	trd.Trade.OnRspQryClassifiedInstrument = func(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInstrument != nil {
			trd.Instruments[pInstrument.InstrumentID.String()] = *pInstrument
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryClassifiedInstrument
			}
		}
	}
	trd.Trade.OnRspQryOrder = func(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pOrder != nil {
			trd.Orders[pOrder.OrderLocalID.String()] = *pOrder
		}
		if bIsLast {
			trd.eventChan <- onRspQryOrder
		}
	}
	trd.Trade.OnRspQryTrade = func(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pTrade != nil {
			if _, ok := trd.Trades[pTrade.OrderLocalID.String()]; ok {
				trd.Trades[pTrade.OrderLocalID.String()] = append(trd.Trades[pTrade.OrderLocalID.String()], *pTrade)
			} else {
				trd.Trades[pTrade.OrderLocalID.String()] = []CThostFtdcTradeField{*pTrade}
			}
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryTrade
			}
		}
	}
	// 银期
	trd.OnRspQryAccountregister = func(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pAccountregister != nil {
			trd.AccountRegisters[pAccountregister.BankAccount.String()] = *pAccountregister
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				trd.errorChan <- *pRspInfo
			} else {
				trd.eventChan <- onRspQryAccountregister
			}
		}
	}

	trd.TradeExt.RegisterFront(cfg.Front)
	trd.TradeExt.SubscribePrivateTopic(THOST_TERT_QUICK)
	trd.TradeExt.SubscribePublicTopic(THOST_TERT_RESTART)
	trd.TradeExt.Init()

	// 登录过程
	select {
	case <-trd.eventChan: // 连接
		trd.ReqAuthenticate(cfg.Broker, cfg.UserID, cfg.AppID, cfg.AuthCode) // 认证
	case <-time.NewTimer(5 * time.Second).C:
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("连接超时 5s"))
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	for {
		select {
		case cb := <-trd.eventChan:
			switch cb {
			case onRspAuthenticate:
				trd.TradeExt.ReqUserLogin(cfg.Password) // 登录
			case onRspUserLogin:
				trd.TradeExt.ReqSettlementInfoConfirm() // 确认结算
			case onRspSettlementInfoConfirm:
				time.Sleep(time.Millisecond * 1100)
				trd.TradeExt.ReqQryInvestor() // 查用户
			case onRspQryInvestor:
				// 交易员登录: 跳过查询过程
				if _, exists := trd.Investors[trd.UserID]; !exists {
					time.Sleep(time.Millisecond * 1100)
					trd.TradeExt.ReqQryAccountregister() // 查银期签约
				} else {
					time.Sleep(time.Millisecond * 1100)
					trd.TradeExt.ReqQryClassifiedInstrument() // 查合约
				}
			case onRspQryClassifiedInstrument:
				time.Sleep(time.Millisecond * 1100)
				trd.TradeExt.ReqQryOrder() // 查委托
			case onRspQryOrder:
				time.Sleep(time.Millisecond * 1100)
				trd.TradeExt.ReqQryTrade() // 查成交
			case onRspQryTrade:
				time.Sleep(time.Millisecond * 1100)
				trd.TradeExt.ReqQryAccountregister() // 查银期签约
			case onRspQryAccountregister:
				fmt.Println("登录过程完成")
				bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("正确"))
				copy(rsp.ErrorMsg[:], bs)
				return
			default:
				fmt.Println("未处理标识:", cb)
			}
		case rsp = <-trd.errorChan:
			return
		}
	}
}

// ReqOrderInsertLimit 限价单
// 成功: 返回 localID
// 失败: 返回 rspInfo 包含错误信息
func (trd *TradePro) ReqOrderInsertLimit(buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, instrument string, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	inst, exists := trd.Instruments[instrument]
	if !exists {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此合约:" + instrument))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	exchange := inst.ExchangeID.String()
	// 最小变动的倍数
	limitPrice := math.Round(price/float64(inst.PriceTick)) * float64(inst.PriceTick)
	trd.TradeExt.ReqOrderInsert(buySell, openClose, instrument, exchange, limitPrice, volume, trd.InvestorID, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_GFD, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)

	select {
	case id := <-trd.orderChan:
		localID = id.String()
	case rsp = <-trd.orderErrChan:
	case <-time.NewTimer(1 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout 1s")
	}
	return
}

// ReqOrderInsertFAK FAK 全成全撤
// 成功: 返回 localID
// 失败: 返回 rspInfo 包含错误信息
func (trd *TradePro) ReqOrderInsertFAK(buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, instrument string, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	inst, exists := trd.Instruments[instrument]
	if !exists {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此合约:" + instrument))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	exchange := inst.ExchangeID.String()
	// 最小变动的倍数
	limitPrice := math.Round(price/float64(inst.PriceTick)) * float64(inst.PriceTick)
	trd.TradeExt.ReqOrderInsert(buySell, openClose, instrument, exchange, limitPrice, volume, trd.InvestorID, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_IOC, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)

	select {
	case id := <-trd.orderChan:
		localID = id.String()
	case rsp = <-trd.orderErrChan:
	case <-time.NewTimer(1 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout 1s")
	}
	return
}

// ReqOrderInsertFOK FOK 部成撤单
// 成功: 返回 localID
// 失败: 返回 rspInfo 包含错误信息
func (trd *TradePro) ReqOrderInsertFOK(buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, instrument string, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	inst, exists := trd.Instruments[instrument]
	if !exists {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此合约:" + instrument))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	exchange := inst.ExchangeID.String()
	// 最小变动的倍数
	limitPrice := math.Round(price/float64(inst.PriceTick)) * float64(inst.PriceTick)
	trd.TradeExt.ReqOrderInsert(buySell, openClose, instrument, exchange, limitPrice, volume, trd.InvestorID, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_IOC, THOST_FTDC_VC_CV, THOST_FTDC_CC_Immediately)

	select {
	case id := <-trd.orderChan:
		localID = id.String()
	case rsp = <-trd.orderErrChan:
	case <-time.NewTimer(1 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout 1s")
	}
	return
}

// ReqOrderInsertMarket 市价单(不是所有交易所都支持)
// 成功: 返回 localID
// 失败: 返回 rspInfo 包含错误信息
func (trd *TradePro) ReqOrderInsertMarket(buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, instrument string, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	inst, exists := trd.Instruments[instrument]
	if !exists {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此合约:" + instrument))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	exchange := inst.ExchangeID.String()
	// 最小变动的倍数
	limitPrice := math.Round(price/float64(inst.PriceTick)) * float64(inst.PriceTick)
	trd.TradeExt.ReqOrderInsert(buySell, openClose, instrument, exchange, limitPrice, volume, trd.InvestorID, THOST_FTDC_OPT_AnyPrice, THOST_FTDC_TC_IOC, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)

	select {
	case id := <-trd.orderChan:
		localID = id.String()
	case rsp = <-trd.orderErrChan:
	case <-time.NewTimer(1 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout 1s")
	}
	return
}

// ReqFromBankToFutureByFuture 入金
func (trd *TradePro) ReqFromBankToFutureByFuture(bankAccount, bankPwd, accountPwd string, amount float64) (rsp CThostFtdcRspInfoField) {
	regInfo, ok := trd.AccountRegisters[bankAccount]
	if !ok {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此帐号:" + bankAccount))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	trd.TradeExt.ReqFromBankToFutureByFuture(regInfo, bankPwd, accountPwd, amount)
	rsp = <-trd.inoutChan
	return
}

// ReqFromFutureToBankByFuture 出金
func (trd *TradePro) ReqFromFutureToBankByFuture(bankAccount, accountPwd string, amount float64) (rsp CThostFtdcRspInfoField) {
	regInfo, ok := trd.AccountRegisters[bankAccount]
	if !ok {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此帐号:" + bankAccount))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	trd.TradeExt.ReqFromFutureToBankByFuture(regInfo, accountPwd, amount)
	rsp = <-trd.inoutChan
	return
}

// ReqQryPosition 查持仓
// 返回 nil 时注意流控
func (trd *TradePro) ReqQryPosition() []CThostFtdcInvestorPositionField {
	trd.positions = make([]CThostFtdcInvestorPositionField, 0)
	trd.TradeExt.ReqQryPosition()
	select {
	case <-trd.eventChan:
		return trd.positions
	case <-time.NewTimer(3 * time.Second).C:
		return nil
	}
}

// ReqQryPositionDetail 查持仓
// 返回 nil 时注意流控
func (trd *TradePro) ReqQryPositionDetail() []CThostFtdcInvestorPositionDetailField {
	trd.positionDetails = make([]CThostFtdcInvestorPositionDetailField, 0)
	trd.TradeExt.ReqQryPositionDetail()
	select {
	case <-trd.eventChan:
		return trd.positionDetails
	case <-time.NewTimer(3 * time.Second).C:
		return nil
	}
}

// ReqQryTradingAccount 查持仓
// 返回 nil 时注意流控
func (trd *TradePro) ReqQryTradingAccount() map[string]CThostFtdcTradingAccountField {
	trd.TradeExt.ReqQryTradingAccount()
	select {
	case <-trd.eventChan:
		return trd.accounts
	case <-time.NewTimer(3 * time.Second).C:
		return nil
	}
}