package def

type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2
)

// 交易所交易员代码类型
type TThostFtdcTraderIDType [21]byte

// 投资者代码类型
type TThostFtdcInvestorIDType [13]byte

// 经纪公司代码类型
type TThostFtdcBrokerIDType [11]byte

// 经纪公司简称类型
type TThostFtdcBrokerAbbrType [9]byte

// 经纪公司名称类型
type TThostFtdcBrokerNameType [81]byte

// 合约在交易所的代码类型
type TThostFtdcOldExchangeInstIDType [31]byte

// 合约在交易所的代码类型
type TThostFtdcExchangeInstIDType [81]byte

// 报单引用类型
type TThostFtdcOrderRefType [13]byte

// 会员代码类型
type TThostFtdcParticipantIDType [11]byte

// 用户代码类型
type TThostFtdcUserIDType [16]byte

// 密码类型
type TThostFtdcPasswordType [41]byte

// 交易编码类型
type TThostFtdcClientIDType [11]byte

// 合约代码类型
type TThostFtdcInstrumentIDType [81]byte

// 合约代码类型
type TThostFtdcOldInstrumentIDType [31]byte

// 合约标识码类型
type TThostFtdcInstrumentCodeType [31]byte

// 市场代码类型
type TThostFtdcMarketIDType [31]byte

// 产品名称类型
type TThostFtdcProductNameType [21]byte

// 交易所代码类型
type TThostFtdcExchangeIDType [9]byte

// 交易所名称类型
type TThostFtdcExchangeNameType [61]byte

// 交易所简称类型
type TThostFtdcExchangeAbbrType [9]byte

// 交易所标志类型
type TThostFtdcExchangeFlagType [2]byte

// Mac地址类型
type TThostFtdcMacAddressType [21]byte

// 系统编号类型
type TThostFtdcSystemIDType [21]byte

// 客户登录备注2类型
type TThostFtdcClientLoginRemarkType [151]byte

// 交易所属性类型
type TThostFtdcExchangePropertyType byte

const THOST_FTDC_EXP_Normal TThostFtdcExchangePropertyType = '0'          // 正常
const THOST_FTDC_EXP_GenOrderByTrade TThostFtdcExchangePropertyType = '1' // 根据成交生成报单

// 日期类型
type TThostFtdcDateType [9]byte

// 时间类型
type TThostFtdcTimeType [9]byte

// 长时间类型
type TThostFtdcLongTimeType [13]byte

// 合约名称类型
type TThostFtdcInstrumentNameType [21]byte

// 结算组代码类型
type TThostFtdcSettlementGroupIDType [9]byte

// 报单编号类型
type TThostFtdcOrderSysIDType [21]byte

// 成交编号类型
type TThostFtdcTradeIDType [21]byte

// DB命令类型类型
type TThostFtdcCommandTypeType [65]byte

// IP地址类型
type TThostFtdcOldIPAddressType [16]byte

// IP地址类型
type TThostFtdcIPAddressType [33]byte

// IP端口类型
type TThostFtdcIPPortType int32

// 产品信息类型
type TThostFtdcProductInfoType [11]byte

// 协议信息类型
type TThostFtdcProtocolInfoType [11]byte

// 业务单元类型
type TThostFtdcBusinessUnitType [21]byte

// 出入金流水号类型
type TThostFtdcDepositSeqNoType [15]byte

// 证件号码类型
type TThostFtdcIdentifiedCardNoType [51]byte

// 证件类型类型
type TThostFtdcIdCardTypeType byte

const THOST_FTDC_ICT_EID TThostFtdcIdCardTypeType = '0'                     // 组织机构代码
const THOST_FTDC_ICT_IDCard TThostFtdcIdCardTypeType = '1'                  // 中国公民身份证
const THOST_FTDC_ICT_OfficerIDCard TThostFtdcIdCardTypeType = '2'           // 军官证
const THOST_FTDC_ICT_PoliceIDCard TThostFtdcIdCardTypeType = '3'            // 警官证
const THOST_FTDC_ICT_SoldierIDCard TThostFtdcIdCardTypeType = '4'           // 士兵证
const THOST_FTDC_ICT_HouseholdRegister TThostFtdcIdCardTypeType = '5'       // 户口簿
const THOST_FTDC_ICT_Passport TThostFtdcIdCardTypeType = '6'                // 护照
const THOST_FTDC_ICT_TaiwanCompatriotIDCard TThostFtdcIdCardTypeType = '7'  // 台胞证
const THOST_FTDC_ICT_HomeComingCard TThostFtdcIdCardTypeType = '8'          // 回乡证
const THOST_FTDC_ICT_LicenseNo TThostFtdcIdCardTypeType = '9'               // 营业执照号
const THOST_FTDC_ICT_TaxNo TThostFtdcIdCardTypeType = 'A'                   // 税务登记号/当地纳税ID
const THOST_FTDC_ICT_HMMainlandTravelPermit TThostFtdcIdCardTypeType = 'B'  // 港澳居民来往内地通行证
const THOST_FTDC_ICT_TwMainlandTravelPermit TThostFtdcIdCardTypeType = 'C'  // 台湾居民来往大陆通行证
const THOST_FTDC_ICT_DrivingLicense TThostFtdcIdCardTypeType = 'D'          // 驾照
const THOST_FTDC_ICT_SocialID TThostFtdcIdCardTypeType = 'F'                // 当地社保ID
const THOST_FTDC_ICT_LocalID TThostFtdcIdCardTypeType = 'G'                 // 当地身份证
const THOST_FTDC_ICT_BusinessRegistration TThostFtdcIdCardTypeType = 'H'    // 商业登记证
const THOST_FTDC_ICT_HKMCIDCard TThostFtdcIdCardTypeType = 'I'              // 港澳永久性居民身份证
const THOST_FTDC_ICT_AccountsPermits TThostFtdcIdCardTypeType = 'J'         // 人行开户许可证
const THOST_FTDC_ICT_FrgPrmtRdCard TThostFtdcIdCardTypeType = 'K'           // 外国人永久居留证
const THOST_FTDC_ICT_CptMngPrdLetter TThostFtdcIdCardTypeType = 'L'         // 资管产品备案函
const THOST_FTDC_ICT_HKMCTwResidencePermit TThostFtdcIdCardTypeType = 'M'   // 港澳台居民居住证
const THOST_FTDC_ICT_UniformSocialCreditCode TThostFtdcIdCardTypeType = 'N' // 统一社会信用代码
const THOST_FTDC_ICT_CorporationCertNo TThostFtdcIdCardTypeType = 'O'       // 机构成立证明文件
const THOST_FTDC_ICT_OtherCard TThostFtdcIdCardTypeType = 'x'               // 其他证件

// 本地报单编号类型
type TThostFtdcOrderLocalIDType [13]byte

// 用户名称类型
type TThostFtdcUserNameType [81]byte

// 参与人名称类型
type TThostFtdcPartyNameType [81]byte

// 错误信息类型
type TThostFtdcErrorMsgType [81]byte

// 字段名类型
type TThostFtdcFieldNameType [2049]byte

// 字段内容类型
type TThostFtdcFieldContentType [2049]byte

// 系统名称类型
type TThostFtdcSystemNameType [41]byte

// 消息正文类型
type TThostFtdcContentType [501]byte

// 投资者范围类型
type TThostFtdcInvestorRangeType byte

const THOST_FTDC_IR_All TThostFtdcInvestorRangeType = '1'    // 所有
const THOST_FTDC_IR_Group TThostFtdcInvestorRangeType = '2'  // 投资者组
const THOST_FTDC_IR_Single TThostFtdcInvestorRangeType = '3' // 单一投资者

// 投资者范围类型
type TThostFtdcDepartmentRangeType byte

const THOST_FTDC_DR_All TThostFtdcDepartmentRangeType = '1'    // 所有
const THOST_FTDC_DR_Group TThostFtdcDepartmentRangeType = '2'  // 组织架构
const THOST_FTDC_DR_Single TThostFtdcDepartmentRangeType = '3' // 单一投资者

// 数据同步状态类型
type TThostFtdcDataSyncStatusType byte

const THOST_FTDC_DS_Asynchronous TThostFtdcDataSyncStatusType = '1'  // 未同步
const THOST_FTDC_DS_Synchronizing TThostFtdcDataSyncStatusType = '2' // 同步中
const THOST_FTDC_DS_Synchronized TThostFtdcDataSyncStatusType = '3'  // 已同步

// 经纪公司数据同步状态类型
type TThostFtdcBrokerDataSyncStatusType byte

const THOST_FTDC_BDS_Synchronized TThostFtdcBrokerDataSyncStatusType = '1'  // 已同步
const THOST_FTDC_BDS_Synchronizing TThostFtdcBrokerDataSyncStatusType = '2' // 同步中

// 交易所连接状态类型
type TThostFtdcExchangeConnectStatusType byte

const THOST_FTDC_ECS_NoConnection TThostFtdcExchangeConnectStatusType = '1'      // 没有任何连接
const THOST_FTDC_ECS_QryInstrumentSent TThostFtdcExchangeConnectStatusType = '2' // 已经发出合约查询请求
const THOST_FTDC_ECS_GotInformation TThostFtdcExchangeConnectStatusType = '9'    // 已经获取信息

// 交易所交易员连接状态类型
type TThostFtdcTraderConnectStatusType byte

const THOST_FTDC_TCS_NotConnected TThostFtdcTraderConnectStatusType = '1'      // 没有任何连接
const THOST_FTDC_TCS_Connected TThostFtdcTraderConnectStatusType = '2'         // 已经连接
const THOST_FTDC_TCS_QryInstrumentSent TThostFtdcTraderConnectStatusType = '3' // 已经发出合约查询请求
const THOST_FTDC_TCS_SubPrivateFlow TThostFtdcTraderConnectStatusType = '4'    // 订阅私有流

// 功能代码类型
type TThostFtdcFunctionCodeType byte

const THOST_FTDC_FC_DataAsync TThostFtdcFunctionCodeType = '1'              // 数据异步化
const THOST_FTDC_FC_ForceUserLogout TThostFtdcFunctionCodeType = '2'        // 强制用户登出
const THOST_FTDC_FC_UserPasswordUpdate TThostFtdcFunctionCodeType = '3'     // 变更管理用户口令
const THOST_FTDC_FC_BrokerPasswordUpdate TThostFtdcFunctionCodeType = '4'   // 变更经纪公司口令
const THOST_FTDC_FC_InvestorPasswordUpdate TThostFtdcFunctionCodeType = '5' // 变更投资者口令
const THOST_FTDC_FC_OrderInsert TThostFtdcFunctionCodeType = '6'            // 报单插入
const THOST_FTDC_FC_OrderAction TThostFtdcFunctionCodeType = '7'            // 报单操作
const THOST_FTDC_FC_SyncSystemData TThostFtdcFunctionCodeType = '8'         // 同步系统数据
const THOST_FTDC_FC_SyncBrokerData TThostFtdcFunctionCodeType = '9'         // 同步经纪公司数据
const THOST_FTDC_FC_BachSyncBrokerData TThostFtdcFunctionCodeType = 'A'     // 批量同步经纪公司数据
const THOST_FTDC_FC_SuperQuery TThostFtdcFunctionCodeType = 'B'             // 超级查询
const THOST_FTDC_FC_ParkedOrderInsert TThostFtdcFunctionCodeType = 'C'      // 预埋报单插入
const THOST_FTDC_FC_ParkedOrderAction TThostFtdcFunctionCodeType = 'D'      // 预埋报单操作
const THOST_FTDC_FC_SyncOTP TThostFtdcFunctionCodeType = 'E'                // 同步动态令牌
const THOST_FTDC_FC_DeleteOrder TThostFtdcFunctionCodeType = 'F'            // 删除未知单

// 经纪公司功能代码类型
type TThostFtdcBrokerFunctionCodeType byte

const THOST_FTDC_BFC_ForceUserLogout TThostFtdcBrokerFunctionCodeType = '1'    // 强制用户登出
const THOST_FTDC_BFC_UserPasswordUpdate TThostFtdcBrokerFunctionCodeType = '2' // 变更用户口令
const THOST_FTDC_BFC_SyncBrokerData TThostFtdcBrokerFunctionCodeType = '3'     // 同步经纪公司数据
const THOST_FTDC_BFC_BachSyncBrokerData TThostFtdcBrokerFunctionCodeType = '4' // 批量同步经纪公司数据
const THOST_FTDC_BFC_OrderInsert TThostFtdcBrokerFunctionCodeType = '5'        // 报单插入
const THOST_FTDC_BFC_OrderAction TThostFtdcBrokerFunctionCodeType = '6'        // 报单操作
const THOST_FTDC_BFC_AllQuery TThostFtdcBrokerFunctionCodeType = '7'           // 全部查询
const THOST_FTDC_BFC_log TThostFtdcBrokerFunctionCodeType = 'a'                // 系统功能：登入/登出/修改密码等
const THOST_FTDC_BFC_BaseQry TThostFtdcBrokerFunctionCodeType = 'b'            // 基本查询：查询基础数据，如合约，交易所等常量
const THOST_FTDC_BFC_TradeQry TThostFtdcBrokerFunctionCodeType = 'c'           // 交易查询：如查成交，委托
const THOST_FTDC_BFC_Trade TThostFtdcBrokerFunctionCodeType = 'd'              // 交易功能：报单，撤单
const THOST_FTDC_BFC_Virement TThostFtdcBrokerFunctionCodeType = 'e'           // 银期转账
const THOST_FTDC_BFC_Risk TThostFtdcBrokerFunctionCodeType = 'f'               // 风险监控
const THOST_FTDC_BFC_Session TThostFtdcBrokerFunctionCodeType = 'g'            // 查询/管理：查询会话，踢人等
const THOST_FTDC_BFC_RiskNoticeCtl TThostFtdcBrokerFunctionCodeType = 'h'      // 风控通知控制
const THOST_FTDC_BFC_RiskNotice TThostFtdcBrokerFunctionCodeType = 'i'         // 风控通知发送
const THOST_FTDC_BFC_BrokerDeposit TThostFtdcBrokerFunctionCodeType = 'j'      // 察看经纪公司资金权限
const THOST_FTDC_BFC_QueryFund TThostFtdcBrokerFunctionCodeType = 'k'          // 资金查询
const THOST_FTDC_BFC_QueryOrder TThostFtdcBrokerFunctionCodeType = 'l'         // 报单查询
const THOST_FTDC_BFC_QueryTrade TThostFtdcBrokerFunctionCodeType = 'm'         // 成交查询
const THOST_FTDC_BFC_QueryPosition TThostFtdcBrokerFunctionCodeType = 'n'      // 持仓查询
const THOST_FTDC_BFC_QueryMarketData TThostFtdcBrokerFunctionCodeType = 'o'    // 行情查询
const THOST_FTDC_BFC_QueryUserEvent TThostFtdcBrokerFunctionCodeType = 'p'     // 用户事件查询
const THOST_FTDC_BFC_QueryRiskNotify TThostFtdcBrokerFunctionCodeType = 'q'    // 风险通知查询
const THOST_FTDC_BFC_QueryFundChange TThostFtdcBrokerFunctionCodeType = 'r'    // 出入金查询
const THOST_FTDC_BFC_QueryInvestor TThostFtdcBrokerFunctionCodeType = 's'      // 投资者信息查询
const THOST_FTDC_BFC_QueryTradingCode TThostFtdcBrokerFunctionCodeType = 't'   // 交易编码查询
const THOST_FTDC_BFC_ForceClose TThostFtdcBrokerFunctionCodeType = 'u'         // 强平
const THOST_FTDC_BFC_PressTest TThostFtdcBrokerFunctionCodeType = 'v'          // 压力测试
const THOST_FTDC_BFC_RemainCalc TThostFtdcBrokerFunctionCodeType = 'w'         // 权益反算
const THOST_FTDC_BFC_NetPositionInd TThostFtdcBrokerFunctionCodeType = 'x'     // 净持仓保证金指标
const THOST_FTDC_BFC_RiskPredict TThostFtdcBrokerFunctionCodeType = 'y'        // 风险预算
const THOST_FTDC_BFC_DataExport TThostFtdcBrokerFunctionCodeType = 'z'         // 数据导出
const THOST_FTDC_BFC_RiskTargetSetup TThostFtdcBrokerFunctionCodeType = 'A'    // 风控指标设置
const THOST_FTDC_BFC_MarketDataWarn TThostFtdcBrokerFunctionCodeType = 'B'     // 行情预警
const THOST_FTDC_BFC_QryBizNotice TThostFtdcBrokerFunctionCodeType = 'C'       // 业务通知查询
const THOST_FTDC_BFC_CfgBizNotice TThostFtdcBrokerFunctionCodeType = 'D'       // 业务通知模板设置
const THOST_FTDC_BFC_SyncOTP TThostFtdcBrokerFunctionCodeType = 'E'            // 同步动态令牌
const THOST_FTDC_BFC_SendBizNotice TThostFtdcBrokerFunctionCodeType = 'F'      // 发送业务通知
const THOST_FTDC_BFC_CfgRiskLevelStd TThostFtdcBrokerFunctionCodeType = 'G'    // 风险级别标准设置
const THOST_FTDC_BFC_TbCommand TThostFtdcBrokerFunctionCodeType = 'H'          // 交易终端应急功能
const THOST_FTDC_BFC_DeleteOrder TThostFtdcBrokerFunctionCodeType = 'J'        // 删除未知单
const THOST_FTDC_BFC_ParkedOrderInsert TThostFtdcBrokerFunctionCodeType = 'K'  // 预埋报单插入
const THOST_FTDC_BFC_ParkedOrderAction TThostFtdcBrokerFunctionCodeType = 'L'  // 预埋报单操作
const THOST_FTDC_BFC_ExecOrderNoCheck TThostFtdcBrokerFunctionCodeType = 'M'   // 资金不够仍允许行权
const THOST_FTDC_BFC_Designate TThostFtdcBrokerFunctionCodeType = 'N'          // 指定
const THOST_FTDC_BFC_StockDisposal TThostFtdcBrokerFunctionCodeType = 'O'      // 证券处置
const THOST_FTDC_BFC_BrokerDepositWarn TThostFtdcBrokerFunctionCodeType = 'Q'  // 席位资金预警
const THOST_FTDC_BFC_CoverWarn TThostFtdcBrokerFunctionCodeType = 'S'          // 备兑不足预警
const THOST_FTDC_BFC_PreExecOrder TThostFtdcBrokerFunctionCodeType = 'T'       // 行权试算
const THOST_FTDC_BFC_ExecOrderRisk TThostFtdcBrokerFunctionCodeType = 'P'      // 行权交收风险
const THOST_FTDC_BFC_PosiLimitWarn TThostFtdcBrokerFunctionCodeType = 'U'      // 持仓限额预警
const THOST_FTDC_BFC_QryPosiLimit TThostFtdcBrokerFunctionCodeType = 'V'       // 持仓限额查询
const THOST_FTDC_BFC_FBSign TThostFtdcBrokerFunctionCodeType = 'W'             // 银期签到签退
const THOST_FTDC_BFC_FBAccount TThostFtdcBrokerFunctionCodeType = 'X'          // 银期签约解约

// 报单操作状态类型
type TThostFtdcOrderActionStatusType byte

const THOST_FTDC_OAS_Submitted TThostFtdcOrderActionStatusType = 'a' // 已经提交
const THOST_FTDC_OAS_Accepted TThostFtdcOrderActionStatusType = 'b'  // 已经接受
const THOST_FTDC_OAS_Rejected TThostFtdcOrderActionStatusType = 'c'  // 已经被拒绝

// 报单状态类型
type TThostFtdcOrderStatusType byte

const THOST_FTDC_OST_AllTraded TThostFtdcOrderStatusType = '0'             // 全部成交
const THOST_FTDC_OST_PartTradedQueueing TThostFtdcOrderStatusType = '1'    // 部分成交还在队列中
const THOST_FTDC_OST_PartTradedNotQueueing TThostFtdcOrderStatusType = '2' // 部分成交不在队列中
const THOST_FTDC_OST_NoTradeQueueing TThostFtdcOrderStatusType = '3'       // 未成交还在队列中
const THOST_FTDC_OST_NoTradeNotQueueing TThostFtdcOrderStatusType = '4'    // 未成交不在队列中
const THOST_FTDC_OST_Canceled TThostFtdcOrderStatusType = '5'              // 撤单
const THOST_FTDC_OST_Unknown TThostFtdcOrderStatusType = 'a'               // 未知
const THOST_FTDC_OST_NotTouched TThostFtdcOrderStatusType = 'b'            // 尚未触发
const THOST_FTDC_OST_Touched TThostFtdcOrderStatusType = 'c'               // 已触发

// 报单提交状态类型
type TThostFtdcOrderSubmitStatusType byte

const THOST_FTDC_OSS_InsertSubmitted TThostFtdcOrderSubmitStatusType = '0' // 已经提交
const THOST_FTDC_OSS_CancelSubmitted TThostFtdcOrderSubmitStatusType = '1' // 撤单已经提交
const THOST_FTDC_OSS_ModifySubmitted TThostFtdcOrderSubmitStatusType = '2' // 修改已经提交
const THOST_FTDC_OSS_Accepted TThostFtdcOrderSubmitStatusType = '3'        // 已经接受
const THOST_FTDC_OSS_InsertRejected TThostFtdcOrderSubmitStatusType = '4'  // 报单已经被拒绝
const THOST_FTDC_OSS_CancelRejected TThostFtdcOrderSubmitStatusType = '5'  // 撤单已经被拒绝
const THOST_FTDC_OSS_ModifyRejected TThostFtdcOrderSubmitStatusType = '6'  // 改单已经被拒绝

// 持仓日期类型
type TThostFtdcPositionDateType byte

const THOST_FTDC_PSD_Today TThostFtdcPositionDateType = '1'   // 今日持仓
const THOST_FTDC_PSD_History TThostFtdcPositionDateType = '2' // 历史持仓

// 持仓日期类型类型
type TThostFtdcPositionDateTypeType byte

const THOST_FTDC_PDT_UseHistory TThostFtdcPositionDateTypeType = '1'   // 使用历史持仓
const THOST_FTDC_PDT_NoUseHistory TThostFtdcPositionDateTypeType = '2' // 不使用历史持仓

// 交易角色类型
type TThostFtdcTradingRoleType byte

const THOST_FTDC_ER_Broker TThostFtdcTradingRoleType = '1' // 代理
const THOST_FTDC_ER_Host TThostFtdcTradingRoleType = '2'   // 自营
const THOST_FTDC_ER_Maker TThostFtdcTradingRoleType = '3'  // 做市商

// 产品类型类型
type TThostFtdcProductClassType byte

const THOST_FTDC_PC_Futures TThostFtdcProductClassType = '1'     // 期货
const THOST_FTDC_PC_Options TThostFtdcProductClassType = '2'     // 期货期权
const THOST_FTDC_PC_Combination TThostFtdcProductClassType = '3' // 组合
const THOST_FTDC_PC_Spot TThostFtdcProductClassType = '4'        // 即期
const THOST_FTDC_PC_EFP TThostFtdcProductClassType = '5'         // 期转现
const THOST_FTDC_PC_SpotOption TThostFtdcProductClassType = '6'  // 现货期权
const THOST_FTDC_PC_TAS TThostFtdcProductClassType = '7'         // TAS合约
const THOST_FTDC_PC_MI TThostFtdcProductClassType = 'I'          // 金属指数

// 产品类型类型
type TThostFtdcAPIProductClassType byte

const THOST_FTDC_APC_FutureSingle TThostFtdcAPIProductClassType = '1'  // 期货单一合约
const THOST_FTDC_APC_OptionSingle TThostFtdcAPIProductClassType = '2'  // 期权单一合约
const THOST_FTDC_APC_Futures TThostFtdcAPIProductClassType = '3'       // 可交易期货(含期货组合和期货单一合约)
const THOST_FTDC_APC_Options TThostFtdcAPIProductClassType = '4'       // 可交易期权(含期权组合和期权单一合约)
const THOST_FTDC_APC_TradingComb TThostFtdcAPIProductClassType = '5'   // 可下单套利组合
const THOST_FTDC_APC_UnTradingComb TThostFtdcAPIProductClassType = '6' // 可申请的组合（可以申请的组合合约 包含可以交易的合约）
const THOST_FTDC_APC_AllTrading TThostFtdcAPIProductClassType = '7'    // 所有可以交易合约
const THOST_FTDC_APC_All TThostFtdcAPIProductClassType = '8'           // 所有合约（包含不能交易合约 慎用）

// 合约生命周期状态类型
type TThostFtdcInstLifePhaseType byte

const THOST_FTDC_IP_NotStart TThostFtdcInstLifePhaseType = '0' // 未上市
const THOST_FTDC_IP_Started TThostFtdcInstLifePhaseType = '1'  // 上市
const THOST_FTDC_IP_Pause TThostFtdcInstLifePhaseType = '2'    // 停牌
const THOST_FTDC_IP_Expired TThostFtdcInstLifePhaseType = '3'  // 到期

// 买卖方向类型
type TThostFtdcDirectionType byte

const THOST_FTDC_D_Buy TThostFtdcDirectionType = '0'  // 买
const THOST_FTDC_D_Sell TThostFtdcDirectionType = '1' // 卖

// 持仓类型类型
type TThostFtdcPositionTypeType byte

const THOST_FTDC_PT_Net TThostFtdcPositionTypeType = '1'   // 净持仓
const THOST_FTDC_PT_Gross TThostFtdcPositionTypeType = '2' // 综合持仓

// 持仓多空方向类型
type TThostFtdcPosiDirectionType byte

const THOST_FTDC_PD_Net TThostFtdcPosiDirectionType = '1'   // 净
const THOST_FTDC_PD_Long TThostFtdcPosiDirectionType = '2'  // 多头
const THOST_FTDC_PD_Short TThostFtdcPosiDirectionType = '3' // 空头

// 系统结算状态类型
type TThostFtdcSysSettlementStatusType byte

const THOST_FTDC_SS_NonActive TThostFtdcSysSettlementStatusType = '1'          // 不活跃
const THOST_FTDC_SS_Startup TThostFtdcSysSettlementStatusType = '2'            // 启动
const THOST_FTDC_SS_Operating TThostFtdcSysSettlementStatusType = '3'          // 操作
const THOST_FTDC_SS_Settlement TThostFtdcSysSettlementStatusType = '4'         // 结算
const THOST_FTDC_SS_SettlementFinished TThostFtdcSysSettlementStatusType = '5' // 结算完成

// 费率属性类型
type TThostFtdcRatioAttrType byte

const THOST_FTDC_RA_Trade TThostFtdcRatioAttrType = '0'      // 交易费率
const THOST_FTDC_RA_Settlement TThostFtdcRatioAttrType = '1' // 结算费率

// 投机套保标志类型
type TThostFtdcHedgeFlagType byte

const THOST_FTDC_HF_Speculation TThostFtdcHedgeFlagType = '1' // 投机
const THOST_FTDC_HF_Arbitrage TThostFtdcHedgeFlagType = '2'   // 套利
const THOST_FTDC_HF_Hedge TThostFtdcHedgeFlagType = '3'       // 套保
const THOST_FTDC_HF_MarketMaker TThostFtdcHedgeFlagType = '5' // 做市商
const THOST_FTDC_HF_SpecHedge TThostFtdcHedgeFlagType = '6'   // 第一腿投机第二腿套保
const THOST_FTDC_HF_HedgeSpec TThostFtdcHedgeFlagType = '7'   // 第一腿套保第二腿投机

// 投机套保标志类型
type TThostFtdcBillHedgeFlagType byte

const THOST_FTDC_BHF_Speculation TThostFtdcBillHedgeFlagType = '1' // 投机
const THOST_FTDC_BHF_Arbitrage TThostFtdcBillHedgeFlagType = '2'   // 套利
const THOST_FTDC_BHF_Hedge TThostFtdcBillHedgeFlagType = '3'       // 套保

// 交易编码类型类型
type TThostFtdcClientIDTypeType byte

const THOST_FTDC_CIDT_Speculation TThostFtdcClientIDTypeType = '1' // 投机
const THOST_FTDC_CIDT_Arbitrage TThostFtdcClientIDTypeType = '2'   // 套利
const THOST_FTDC_CIDT_Hedge TThostFtdcClientIDTypeType = '3'       // 套保
const THOST_FTDC_CIDT_MarketMaker TThostFtdcClientIDTypeType = '5' // 做市商

// 报单价格条件类型
type TThostFtdcOrderPriceTypeType byte

const THOST_FTDC_OPT_AnyPrice TThostFtdcOrderPriceTypeType = '1'                // 任意价
const THOST_FTDC_OPT_LimitPrice TThostFtdcOrderPriceTypeType = '2'              // 限价
const THOST_FTDC_OPT_BestPrice TThostFtdcOrderPriceTypeType = '3'               // 最优价
const THOST_FTDC_OPT_LastPrice TThostFtdcOrderPriceTypeType = '4'               // 最新价
const THOST_FTDC_OPT_LastPricePlusOneTicks TThostFtdcOrderPriceTypeType = '5'   // 最新价浮动上浮1个ticks
const THOST_FTDC_OPT_LastPricePlusTwoTicks TThostFtdcOrderPriceTypeType = '6'   // 最新价浮动上浮2个ticks
const THOST_FTDC_OPT_LastPricePlusThreeTicks TThostFtdcOrderPriceTypeType = '7' // 最新价浮动上浮3个ticks
const THOST_FTDC_OPT_AskPrice1 TThostFtdcOrderPriceTypeType = '8'               // 卖一价
const THOST_FTDC_OPT_AskPrice1PlusOneTicks TThostFtdcOrderPriceTypeType = '9'   // 卖一价浮动上浮1个ticks
const THOST_FTDC_OPT_AskPrice1PlusTwoTicks TThostFtdcOrderPriceTypeType = 'A'   // 卖一价浮动上浮2个ticks
const THOST_FTDC_OPT_AskPrice1PlusThreeTicks TThostFtdcOrderPriceTypeType = 'B' // 卖一价浮动上浮3个ticks
const THOST_FTDC_OPT_BidPrice1 TThostFtdcOrderPriceTypeType = 'C'               // 买一价
const THOST_FTDC_OPT_BidPrice1PlusOneTicks TThostFtdcOrderPriceTypeType = 'D'   // 买一价浮动上浮1个ticks
const THOST_FTDC_OPT_BidPrice1PlusTwoTicks TThostFtdcOrderPriceTypeType = 'E'   // 买一价浮动上浮2个ticks
const THOST_FTDC_OPT_BidPrice1PlusThreeTicks TThostFtdcOrderPriceTypeType = 'F' // 买一价浮动上浮3个ticks
const THOST_FTDC_OPT_FiveLevelPrice TThostFtdcOrderPriceTypeType = 'G'          // 五档价

// 开平标志类型
type TThostFtdcOffsetFlagType byte

const THOST_FTDC_OF_Open TThostFtdcOffsetFlagType = '0'            // 开仓
const THOST_FTDC_OF_Close TThostFtdcOffsetFlagType = '1'           // 平仓
const THOST_FTDC_OF_ForceClose TThostFtdcOffsetFlagType = '2'      // 强平
const THOST_FTDC_OF_CloseToday TThostFtdcOffsetFlagType = '3'      // 平今
const THOST_FTDC_OF_CloseYesterday TThostFtdcOffsetFlagType = '4'  // 平昨
const THOST_FTDC_OF_ForceOff TThostFtdcOffsetFlagType = '5'        // 强减
const THOST_FTDC_OF_LocalForceClose TThostFtdcOffsetFlagType = '6' // 本地强平

// 强平原因类型
type TThostFtdcForceCloseReasonType byte

const THOST_FTDC_FCC_NotForceClose TThostFtdcForceCloseReasonType = '0'           // 非强平
const THOST_FTDC_FCC_LackDeposit TThostFtdcForceCloseReasonType = '1'             // 资金不足
const THOST_FTDC_FCC_ClientOverPositionLimit TThostFtdcForceCloseReasonType = '2' // 客户超仓
const THOST_FTDC_FCC_MemberOverPositionLimit TThostFtdcForceCloseReasonType = '3' // 会员超仓
const THOST_FTDC_FCC_NotMultiple TThostFtdcForceCloseReasonType = '4'             // 持仓非整数倍
const THOST_FTDC_FCC_Violation TThostFtdcForceCloseReasonType = '5'               // 违规
const THOST_FTDC_FCC_Other TThostFtdcForceCloseReasonType = '6'                   // 其它
const THOST_FTDC_FCC_PersonDeliv TThostFtdcForceCloseReasonType = '7'             // 自然人临近交割
const THOST_FTDC_FCC_Notverifycapital TThostFtdcForceCloseReasonType = '8'        // 风控强平不验证资金

// 报单类型类型
type TThostFtdcOrderTypeType byte

const THOST_FTDC_ORDT_Normal TThostFtdcOrderTypeType = '0'                // 正常
const THOST_FTDC_ORDT_DeriveFromQuote TThostFtdcOrderTypeType = '1'       // 报价衍生
const THOST_FTDC_ORDT_DeriveFromCombination TThostFtdcOrderTypeType = '2' // 组合衍生
const THOST_FTDC_ORDT_Combination TThostFtdcOrderTypeType = '3'           // 组合报单
const THOST_FTDC_ORDT_ConditionalOrder TThostFtdcOrderTypeType = '4'      // 条件单
const THOST_FTDC_ORDT_Swap TThostFtdcOrderTypeType = '5'                  // 互换单
const THOST_FTDC_ORDT_DeriveFromBlockTrade TThostFtdcOrderTypeType = '6'  // 大宗交易成交衍生
const THOST_FTDC_ORDT_DeriveFromEFPTrade TThostFtdcOrderTypeType = '7'    // 期转现成交衍生

// 有效期类型类型
type TThostFtdcTimeConditionType byte

const THOST_FTDC_TC_IOC TThostFtdcTimeConditionType = '1' // 立即完成，否则撤销
const THOST_FTDC_TC_GFS TThostFtdcTimeConditionType = '2' // 本节有效
const THOST_FTDC_TC_GFD TThostFtdcTimeConditionType = '3' // 当日有效
const THOST_FTDC_TC_GTD TThostFtdcTimeConditionType = '4' // 指定日期前有效
const THOST_FTDC_TC_GTC TThostFtdcTimeConditionType = '5' // 撤销前有效
const THOST_FTDC_TC_GFA TThostFtdcTimeConditionType = '6' // 集合竞价有效

// 成交量类型类型
type TThostFtdcVolumeConditionType byte

const THOST_FTDC_VC_AV TThostFtdcVolumeConditionType = '1' // 任何数量
const THOST_FTDC_VC_MV TThostFtdcVolumeConditionType = '2' // 最小数量
const THOST_FTDC_VC_CV TThostFtdcVolumeConditionType = '3' // 全部数量

// 触发条件类型
type TThostFtdcContingentConditionType byte

const THOST_FTDC_CC_Immediately TThostFtdcContingentConditionType = '1'                    // 立即
const THOST_FTDC_CC_Touch TThostFtdcContingentConditionType = '2'                          // 止损
const THOST_FTDC_CC_TouchProfit TThostFtdcContingentConditionType = '3'                    // 止赢
const THOST_FTDC_CC_ParkedOrder TThostFtdcContingentConditionType = '4'                    // 预埋单
const THOST_FTDC_CC_LastPriceGreaterThanStopPrice TThostFtdcContingentConditionType = '5'  // 最新价大于条件价
const THOST_FTDC_CC_LastPriceGreaterEqualStopPrice TThostFtdcContingentConditionType = '6' // 最新价大于等于条件价
const THOST_FTDC_CC_LastPriceLesserThanStopPrice TThostFtdcContingentConditionType = '7'   // 最新价小于条件价
const THOST_FTDC_CC_LastPriceLesserEqualStopPrice TThostFtdcContingentConditionType = '8'  // 最新价小于等于条件价
const THOST_FTDC_CC_AskPriceGreaterThanStopPrice TThostFtdcContingentConditionType = '9'   // 卖一价大于条件价
const THOST_FTDC_CC_AskPriceGreaterEqualStopPrice TThostFtdcContingentConditionType = 'A'  // 卖一价大于等于条件价
const THOST_FTDC_CC_AskPriceLesserThanStopPrice TThostFtdcContingentConditionType = 'B'    // 卖一价小于条件价
const THOST_FTDC_CC_AskPriceLesserEqualStopPrice TThostFtdcContingentConditionType = 'C'   // 卖一价小于等于条件价
const THOST_FTDC_CC_BidPriceGreaterThanStopPrice TThostFtdcContingentConditionType = 'D'   // 买一价大于条件价
const THOST_FTDC_CC_BidPriceGreaterEqualStopPrice TThostFtdcContingentConditionType = 'E'  // 买一价大于等于条件价
const THOST_FTDC_CC_BidPriceLesserThanStopPrice TThostFtdcContingentConditionType = 'F'    // 买一价小于条件价
const THOST_FTDC_CC_BidPriceLesserEqualStopPrice TThostFtdcContingentConditionType = 'H'   // 买一价小于等于条件价

// 操作标志类型
type TThostFtdcActionFlagType byte

const THOST_FTDC_AF_Delete TThostFtdcActionFlagType = '0' // 删除
const THOST_FTDC_AF_Modify TThostFtdcActionFlagType = '3' // 修改

// 交易权限类型
type TThostFtdcTradingRightType byte

const THOST_FTDC_TR_Allow TThostFtdcTradingRightType = '0'     // 可以交易
const THOST_FTDC_TR_CloseOnly TThostFtdcTradingRightType = '1' // 只能平仓
const THOST_FTDC_TR_Forbidden TThostFtdcTradingRightType = '2' // 不能交易

// 报单来源类型
type TThostFtdcOrderSourceType byte

const THOST_FTDC_OSRC_Participant TThostFtdcOrderSourceType = '0'   // 来自参与者
const THOST_FTDC_OSRC_Administrator TThostFtdcOrderSourceType = '1' // 来自管理员

// 成交类型类型
type TThostFtdcTradeTypeType byte

const THOST_FTDC_TRDT_SplitCombination TThostFtdcTradeTypeType = '#'   // 组合持仓拆分为单一持仓,初始化不应包含该类型的持仓
const THOST_FTDC_TRDT_Common TThostFtdcTradeTypeType = '0'             // 普通成交
const THOST_FTDC_TRDT_OptionsExecution TThostFtdcTradeTypeType = '1'   // 期权执行
const THOST_FTDC_TRDT_OTC TThostFtdcTradeTypeType = '2'                // OTC成交
const THOST_FTDC_TRDT_EFPDerived TThostFtdcTradeTypeType = '3'         // 期转现衍生成交
const THOST_FTDC_TRDT_CombinationDerived TThostFtdcTradeTypeType = '4' // 组合衍生成交
const THOST_FTDC_TRDT_BlockTrade TThostFtdcTradeTypeType = '5'         // 大宗交易成交

// 特殊持仓明细标识类型
type TThostFtdcSpecPosiTypeType byte

const THOST_FTDC_SPOST_Common TThostFtdcSpecPosiTypeType = '#' // 普通持仓明细
const THOST_FTDC_SPOST_Tas TThostFtdcSpecPosiTypeType = '0'    // TAS合约成交产生的标的合约持仓明细

// 成交价来源类型
type TThostFtdcPriceSourceType byte

const THOST_FTDC_PSRC_LastPrice TThostFtdcPriceSourceType = '0' // 前成交价
const THOST_FTDC_PSRC_Buy TThostFtdcPriceSourceType = '1'       // 买委托价
const THOST_FTDC_PSRC_Sell TThostFtdcPriceSourceType = '2'      // 卖委托价
const THOST_FTDC_PSRC_OTC TThostFtdcPriceSourceType = '3'       // 场外成交价

// 合约交易状态类型
type TThostFtdcInstrumentStatusType byte

const THOST_FTDC_IS_BeforeTrading TThostFtdcInstrumentStatusType = '0'   // 开盘前
const THOST_FTDC_IS_NoTrading TThostFtdcInstrumentStatusType = '1'       // 非交易
const THOST_FTDC_IS_Continous TThostFtdcInstrumentStatusType = '2'       // 连续交易
const THOST_FTDC_IS_AuctionOrdering TThostFtdcInstrumentStatusType = '3' // 集合竞价报单
const THOST_FTDC_IS_AuctionBalance TThostFtdcInstrumentStatusType = '4'  // 集合竞价价格平衡
const THOST_FTDC_IS_AuctionMatch TThostFtdcInstrumentStatusType = '5'    // 集合竞价撮合
const THOST_FTDC_IS_Closed TThostFtdcInstrumentStatusType = '6'          // 收盘

// 品种进入交易状态原因类型
type TThostFtdcInstStatusEnterReasonType byte

const THOST_FTDC_IER_Automatic TThostFtdcInstStatusEnterReasonType = '1' // 自动切换
const THOST_FTDC_IER_Manual TThostFtdcInstStatusEnterReasonType = '2'    // 手动切换
const THOST_FTDC_IER_Fuse TThostFtdcInstStatusEnterReasonType = '3'      // 熔断

// 报单操作引用类型
type TThostFtdcOrderActionRefType int32

// 安装数量类型
type TThostFtdcInstallCountType int32

// 安装编号类型
type TThostFtdcInstallIDType int32

// 错误代码类型
type TThostFtdcErrorIDType int32

// 结算编号类型
type TThostFtdcSettlementIDType int32

// 数量类型
type TThostFtdcVolumeType int32

// 前置编号类型
type TThostFtdcFrontIDType int32

// 会话编号类型
type TThostFtdcSessionIDType int32

// 序号类型
type TThostFtdcSequenceNoType int32

// DB命令序号类型
type TThostFtdcCommandNoType int32

// 时间（毫秒）类型
type TThostFtdcMillisecType int32

// 时间（秒）类型
type TThostFtdcSecType int32

// 合约数量乘数类型
type TThostFtdcVolumeMultipleType int32

// 交易阶段编号类型
type TThostFtdcTradingSegmentSNType int32

// 请求编号类型
type TThostFtdcRequestIDType int32

// 年份类型
type TThostFtdcYearType int32

// 月份类型
type TThostFtdcMonthType int32

// 布尔型类型
type TThostFtdcBoolType int32

// 价格类型
type TThostFtdcPriceType float64

// 组合开平标志类型
type TThostFtdcCombOffsetFlagType [5]byte

// 组合投机套保标志类型
type TThostFtdcCombHedgeFlagType [5]byte

// 比率类型
type TThostFtdcRatioType float64

// 资金类型
type TThostFtdcMoneyType float64

// 大额数量类型
type TThostFtdcLargeVolumeType float64

// 序列系列号类型
type TThostFtdcSequenceSeriesType int16

// 通讯时段编号类型
type TThostFtdcCommPhaseNoType int16

// 序列编号类型
type TThostFtdcSequenceLabelType [2]byte

// 基础商品乘数类型
type TThostFtdcUnderlyingMultipleType float64

// 优先级类型
type TThostFtdcPriorityType int32

// 合同编号类型
type TThostFtdcContractCodeType [41]byte

// 市类型
type TThostFtdcCityType [51]byte

// 是否股民类型
type TThostFtdcIsStockType [11]byte

// 渠道类型
type TThostFtdcChannelType [51]byte

// 通讯地址类型
type TThostFtdcAddressType [101]byte

// 邮政编码类型
type TThostFtdcZipCodeType [7]byte

// 联系电话类型
type TThostFtdcTelephoneType [41]byte

// 传真类型
type TThostFtdcFaxType [41]byte

// 手机类型
type TThostFtdcMobileType [41]byte

// 电子邮件类型
type TThostFtdcEMailType [41]byte

// 备注类型
type TThostFtdcMemoType [161]byte

// 企业代码类型
type TThostFtdcCompanyCodeType [51]byte

// 网站地址类型
type TThostFtdcWebsiteType [51]byte

// 税务登记号类型
type TThostFtdcTaxNoType [31]byte

// 处理状态类型
type TThostFtdcBatchStatusType byte

const THOST_FTDC_BS_NoUpload TThostFtdcBatchStatusType = '1' // 未上传
const THOST_FTDC_BS_Uploaded TThostFtdcBatchStatusType = '2' // 已上传
const THOST_FTDC_BS_Failed TThostFtdcBatchStatusType = '3'   // 审核失败

// 属性代码类型
type TThostFtdcPropertyIDType [33]byte

// 属性名称类型
type TThostFtdcPropertyNameType [65]byte

// 营业执照号类型
type TThostFtdcLicenseNoType [51]byte

// 经纪人代码类型
type TThostFtdcAgentIDType [13]byte

// 经纪人名称类型
type TThostFtdcAgentNameType [41]byte

// 经纪人组代码类型
type TThostFtdcAgentGroupIDType [13]byte

// 经纪人组名称类型
type TThostFtdcAgentGroupNameType [41]byte

// 按品种返还方式类型
type TThostFtdcReturnStyleType byte

const THOST_FTDC_RS_All TThostFtdcReturnStyleType = '1'       // 按所有品种
const THOST_FTDC_RS_ByProduct TThostFtdcReturnStyleType = '2' // 按品种

// 返还模式类型
type TThostFtdcReturnPatternType byte

const THOST_FTDC_RP_ByVolume TThostFtdcReturnPatternType = '1'    // 按成交手数
const THOST_FTDC_RP_ByFeeOnHand TThostFtdcReturnPatternType = '2' // 按留存手续费

// 返还级别类型
type TThostFtdcReturnLevelType byte

const THOST_FTDC_RL_Level1 TThostFtdcReturnLevelType = '1' // 级别1
const THOST_FTDC_RL_Level2 TThostFtdcReturnLevelType = '2' // 级别2
const THOST_FTDC_RL_Level3 TThostFtdcReturnLevelType = '3' // 级别3
const THOST_FTDC_RL_Level4 TThostFtdcReturnLevelType = '4' // 级别4
const THOST_FTDC_RL_Level5 TThostFtdcReturnLevelType = '5' // 级别5
const THOST_FTDC_RL_Level6 TThostFtdcReturnLevelType = '6' // 级别6
const THOST_FTDC_RL_Level7 TThostFtdcReturnLevelType = '7' // 级别7
const THOST_FTDC_RL_Level8 TThostFtdcReturnLevelType = '8' // 级别8
const THOST_FTDC_RL_Level9 TThostFtdcReturnLevelType = '9' // 级别9

// 返还标准类型
type TThostFtdcReturnStandardType byte

const THOST_FTDC_RSD_ByPeriod TThostFtdcReturnStandardType = '1'   // 分阶段返还
const THOST_FTDC_RSD_ByStandard TThostFtdcReturnStandardType = '2' // 按某一标准

// 质押类型类型
type TThostFtdcMortgageTypeType byte

const THOST_FTDC_MT_Out TThostFtdcMortgageTypeType = '0' // 质出
const THOST_FTDC_MT_In TThostFtdcMortgageTypeType = '1'  // 质入

// 投资者结算参数代码类型
type TThostFtdcInvestorSettlementParamIDType byte

const THOST_FTDC_ISPI_MortgageRatio TThostFtdcInvestorSettlementParamIDType = '4' // 质押比例
const THOST_FTDC_ISPI_MarginWay TThostFtdcInvestorSettlementParamIDType = '5'     // 保证金算法
const THOST_FTDC_ISPI_BillDeposit TThostFtdcInvestorSettlementParamIDType = '9'   // 结算单结存是否包含质押

// 交易所结算参数代码类型
type TThostFtdcExchangeSettlementParamIDType byte

const THOST_FTDC_ESPI_MortgageRatio TThostFtdcExchangeSettlementParamIDType = '1'      // 质押比例
const THOST_FTDC_ESPI_OtherFundItem TThostFtdcExchangeSettlementParamIDType = '2'      // 分项资金导入项
const THOST_FTDC_ESPI_OtherFundImport TThostFtdcExchangeSettlementParamIDType = '3'    // 分项资金入交易所出入金
const THOST_FTDC_ESPI_CFFEXMinPrepa TThostFtdcExchangeSettlementParamIDType = '6'      // 中金所开户最低可用金额
const THOST_FTDC_ESPI_CZCESettlementType TThostFtdcExchangeSettlementParamIDType = '7' // 郑商所结算方式
const THOST_FTDC_ESPI_ExchDelivFeeMode TThostFtdcExchangeSettlementParamIDType = '9'   // 交易所交割手续费收取方式
const THOST_FTDC_ESPI_DelivFeeMode TThostFtdcExchangeSettlementParamIDType = '0'       // 投资者交割手续费收取方式
const THOST_FTDC_ESPI_CZCEComMarginType TThostFtdcExchangeSettlementParamIDType = 'A'  // 郑商所组合持仓保证金收取方式
const THOST_FTDC_ESPI_DceComMarginType TThostFtdcExchangeSettlementParamIDType = 'B'   // 大商所套利保证金是否优惠
const THOST_FTDC_ESPI_OptOutDisCountRate TThostFtdcExchangeSettlementParamIDType = 'a' // 虚值期权保证金优惠比率
const THOST_FTDC_ESPI_OptMiniGuarantee TThostFtdcExchangeSettlementParamIDType = 'b'   // 最低保障系数

// 系统参数代码类型
type TThostFtdcSystemParamIDType byte

const THOST_FTDC_SPI_InvestorIDMinLength TThostFtdcSystemParamIDType = '1'     // 投资者代码最小长度
const THOST_FTDC_SPI_AccountIDMinLength TThostFtdcSystemParamIDType = '2'      // 投资者帐号代码最小长度
const THOST_FTDC_SPI_UserRightLogon TThostFtdcSystemParamIDType = '3'          // 投资者开户默认登录权限
const THOST_FTDC_SPI_SettlementBillTrade TThostFtdcSystemParamIDType = '4'     // 投资者交易结算单成交汇总方式
const THOST_FTDC_SPI_TradingCode TThostFtdcSystemParamIDType = '5'             // 统一开户更新交易编码方式
const THOST_FTDC_SPI_CheckFund TThostFtdcSystemParamIDType = '6'               // 结算是否判断存在未复核的出入金和分项资金
const THOST_FTDC_SPI_CommModelRight TThostFtdcSystemParamIDType = '7'          // 是否启用手续费模板数据权限
const THOST_FTDC_SPI_MarginModelRight TThostFtdcSystemParamIDType = '9'        // 是否启用保证金率模板数据权限
const THOST_FTDC_SPI_IsStandardActive TThostFtdcSystemParamIDType = '8'        // 是否规范用户才能激活
const THOST_FTDC_SPI_UploadSettlementFile TThostFtdcSystemParamIDType = 'U'    // 上传的交易所结算文件路径
const THOST_FTDC_SPI_DownloadCSRCFile TThostFtdcSystemParamIDType = 'D'        // 上报保证金监控中心文件路径
const THOST_FTDC_SPI_SettlementBillFile TThostFtdcSystemParamIDType = 'S'      // 生成的结算单文件路径
const THOST_FTDC_SPI_CSRCOthersFile TThostFtdcSystemParamIDType = 'C'          // 证监会文件标识
const THOST_FTDC_SPI_InvestorPhoto TThostFtdcSystemParamIDType = 'P'           // 投资者照片路径
const THOST_FTDC_SPI_CSRCData TThostFtdcSystemParamIDType = 'R'                // 全结经纪公司上传文件路径
const THOST_FTDC_SPI_InvestorPwdModel TThostFtdcSystemParamIDType = 'I'        // 开户密码录入方式
const THOST_FTDC_SPI_CFFEXInvestorSettleFile TThostFtdcSystemParamIDType = 'F' // 投资者中金所结算文件下载路径
const THOST_FTDC_SPI_InvestorIDType TThostFtdcSystemParamIDType = 'a'          // 投资者代码编码方式
const THOST_FTDC_SPI_FreezeMaxReMain TThostFtdcSystemParamIDType = 'r'         // 休眠户最高权益
const THOST_FTDC_SPI_IsSync TThostFtdcSystemParamIDType = 'A'                  // 手续费相关操作实时上场开关
const THOST_FTDC_SPI_RelieveOpenLimit TThostFtdcSystemParamIDType = 'O'        // 解除开仓权限限制
const THOST_FTDC_SPI_IsStandardFreeze TThostFtdcSystemParamIDType = 'X'        // 是否规范用户才能休眠
const THOST_FTDC_SPI_CZCENormalProductHedge TThostFtdcSystemParamIDType = 'B'  // 郑商所是否开放所有品种套保交易

// 交易系统参数代码类型
type TThostFtdcTradeParamIDType byte

const THOST_FTDC_TPID_EncryptionStandard TThostFtdcTradeParamIDType = 'E'      // 系统加密算法
const THOST_FTDC_TPID_RiskMode TThostFtdcTradeParamIDType = 'R'                // 系统风险算法
const THOST_FTDC_TPID_RiskModeGlobal TThostFtdcTradeParamIDType = 'G'          // 系统风险算法是否全局 0-否 1-是
const THOST_FTDC_TPID_modeEncode TThostFtdcTradeParamIDType = 'P'              // 密码加密算法
const THOST_FTDC_TPID_tickMode TThostFtdcTradeParamIDType = 'T'                // 价格小数位数参数
const THOST_FTDC_TPID_SingleUserSessionMaxNum TThostFtdcTradeParamIDType = 'S' // 用户最大会话数
const THOST_FTDC_TPID_LoginFailMaxNum TThostFtdcTradeParamIDType = 'L'         // 最大连续登录失败数
const THOST_FTDC_TPID_IsAuthForce TThostFtdcTradeParamIDType = 'A'             // 是否强制认证
const THOST_FTDC_TPID_IsPosiFreeze TThostFtdcTradeParamIDType = 'F'            // 是否冻结证券持仓
const THOST_FTDC_TPID_IsPosiLimit TThostFtdcTradeParamIDType = 'M'             // 是否限仓
const THOST_FTDC_TPID_ForQuoteTimeInterval TThostFtdcTradeParamIDType = 'Q'    // 郑商所询价时间间隔
const THOST_FTDC_TPID_IsFuturePosiLimit TThostFtdcTradeParamIDType = 'B'       // 是否期货限仓
const THOST_FTDC_TPID_IsFutureOrderFreq TThostFtdcTradeParamIDType = 'C'       // 是否期货下单频率限制
const THOST_FTDC_TPID_IsExecOrderProfit TThostFtdcTradeParamIDType = 'H'       // 行权冻结是否计算盈利
const THOST_FTDC_TPID_IsCheckBankAcc TThostFtdcTradeParamIDType = 'I'          // 银期开户是否验证开户银行卡号是否是预留银行账户
const THOST_FTDC_TPID_PasswordDeadLine TThostFtdcTradeParamIDType = 'J'        // 弱密码最后修改日期
const THOST_FTDC_TPID_IsStrongPassword TThostFtdcTradeParamIDType = 'K'        // 强密码校验
const THOST_FTDC_TPID_BalanceMorgage TThostFtdcTradeParamIDType = 'a'          // 自有资金质押比
const THOST_FTDC_TPID_MinPwdLen TThostFtdcTradeParamIDType = 'O'               // 最小密码长度
const THOST_FTDC_TPID_LoginFailMaxNumForIP TThostFtdcTradeParamIDType = 'U'    // IP当日最大登陆失败次数
const THOST_FTDC_TPID_PasswordPeriod TThostFtdcTradeParamIDType = 'V'          // 密码有效期

// 参数代码值类型
type TThostFtdcSettlementParamValueType [256]byte

// 计数器代码类型
type TThostFtdcCounterIDType [33]byte

// 投资者分组名称类型
type TThostFtdcInvestorGroupNameType [41]byte

// 牌号类型
type TThostFtdcBrandCodeType [257]byte

// 仓库类型
type TThostFtdcWarehouseType [257]byte

// 产期类型
type TThostFtdcProductDateType [41]byte

// 等级类型
type TThostFtdcGradeType [41]byte

// 类别类型
type TThostFtdcClassifyType [41]byte

// 货位类型
type TThostFtdcPositionType [41]byte

// 产地类型
type TThostFtdcYieldlyType [41]byte

// 公定重量类型
type TThostFtdcWeightType [41]byte

// 分项资金流水号类型
type TThostFtdcSubEntryFundNoType int32

// 文件标识类型
type TThostFtdcFileIDType byte

const THOST_FTDC_FI_SettlementFund TThostFtdcFileIDType = 'F'            // 资金数据
const THOST_FTDC_FI_Trade TThostFtdcFileIDType = 'T'                     // 成交数据
const THOST_FTDC_FI_InvestorPosition TThostFtdcFileIDType = 'P'          // 投资者持仓数据
const THOST_FTDC_FI_SubEntryFund TThostFtdcFileIDType = 'O'              // 投资者分项资金数据
const THOST_FTDC_FI_CZCECombinationPos TThostFtdcFileIDType = 'C'        // 组合持仓数据
const THOST_FTDC_FI_CSRCData TThostFtdcFileIDType = 'R'                  // 上报保证金监控中心数据
const THOST_FTDC_FI_CZCEClose TThostFtdcFileIDType = 'L'                 // 郑商所平仓了结数据
const THOST_FTDC_FI_CZCENoClose TThostFtdcFileIDType = 'N'               // 郑商所非平仓了结数据
const THOST_FTDC_FI_PositionDtl TThostFtdcFileIDType = 'D'               // 持仓明细数据
const THOST_FTDC_FI_OptionStrike TThostFtdcFileIDType = 'S'              // 期权执行文件
const THOST_FTDC_FI_SettlementPriceComparison TThostFtdcFileIDType = 'M' // 结算价比对文件
const THOST_FTDC_FI_NonTradePosChange TThostFtdcFileIDType = 'B'         // 上期所非持仓变动明细

// 文件名称类型
type TThostFtdcFileNameType [257]byte

// 文件上传类型类型
type TThostFtdcFileTypeType byte

const THOST_FTDC_FUT_Settlement TThostFtdcFileTypeType = '0' // 结算
const THOST_FTDC_FUT_Check TThostFtdcFileTypeType = '1'      // 核对

// 文件格式类型
type TThostFtdcFileFormatType byte

const THOST_FTDC_FFT_Txt TThostFtdcFileFormatType = '0' // 文本文件(.txt)
const THOST_FTDC_FFT_Zip TThostFtdcFileFormatType = '1' // 压缩文件(.zip)
const THOST_FTDC_FFT_DBF TThostFtdcFileFormatType = '2' // DBF文件(.dbf)

// 文件状态类型
type TThostFtdcFileUploadStatusType byte

const THOST_FTDC_FUS_SucceedUpload TThostFtdcFileUploadStatusType = '1'   // 上传成功
const THOST_FTDC_FUS_FailedUpload TThostFtdcFileUploadStatusType = '2'    // 上传失败
const THOST_FTDC_FUS_SucceedLoad TThostFtdcFileUploadStatusType = '3'     // 导入成功
const THOST_FTDC_FUS_PartSucceedLoad TThostFtdcFileUploadStatusType = '4' // 导入部分成功
const THOST_FTDC_FUS_FailedLoad TThostFtdcFileUploadStatusType = '5'      // 导入失败

// 移仓方向类型
type TThostFtdcTransferDirectionType byte

const THOST_FTDC_TD_Out TThostFtdcTransferDirectionType = '0' // 移出
const THOST_FTDC_TD_In TThostFtdcTransferDirectionType = '1'  // 移入

// 上传文件类型类型
type TThostFtdcUploadModeType [21]byte

// 投资者帐号类型
type TThostFtdcAccountIDType [13]byte

// 银行统一标识类型类型
type TThostFtdcBankFlagType [4]byte

// 银行账户类型
type TThostFtdcBankAccountType [41]byte

// 银行账户的开户人名称类型
type TThostFtdcOpenNameType [61]byte

// 银行账户的开户行类型
type TThostFtdcOpenBankType [101]byte

// 银行名称类型
type TThostFtdcBankNameType [101]byte

// 发布路径类型
type TThostFtdcPublishPathType [257]byte

// 操作员代码类型
type TThostFtdcOperatorIDType [65]byte

// 月份数量类型
type TThostFtdcMonthCountType int32

// 月份提前数组类型
type TThostFtdcAdvanceMonthArrayType [13]byte

// 日期表达式类型
type TThostFtdcDateExprType [1025]byte

// 合约代码表达式类型
type TThostFtdcInstrumentIDExprType [41]byte

// 合约名称表达式类型
type TThostFtdcInstrumentNameExprType [41]byte

// 特殊的创建规则类型
type TThostFtdcSpecialCreateRuleType byte

const THOST_FTDC_SC_NoSpecialRule TThostFtdcSpecialCreateRuleType = '0'    // 没有特殊创建规则
const THOST_FTDC_SC_NoSpringFestival TThostFtdcSpecialCreateRuleType = '1' // 不包含春节

// 挂牌基准价类型类型
type TThostFtdcBasisPriceTypeType byte

const THOST_FTDC_IPT_LastSettlement TThostFtdcBasisPriceTypeType = '1' // 上一合约结算价
const THOST_FTDC_IPT_LaseClose TThostFtdcBasisPriceTypeType = '2'      // 上一合约收盘价

// 产品生命周期状态类型
type TThostFtdcProductLifePhaseType byte

const THOST_FTDC_PLP_Active TThostFtdcProductLifePhaseType = '1'    // 活跃
const THOST_FTDC_PLP_NonActive TThostFtdcProductLifePhaseType = '2' // 不活跃
const THOST_FTDC_PLP_Canceled TThostFtdcProductLifePhaseType = '3'  // 注销

// 交割方式类型
type TThostFtdcDeliveryModeType byte

const THOST_FTDC_DM_CashDeliv TThostFtdcDeliveryModeType = '1'      // 现金交割
const THOST_FTDC_DM_CommodityDeliv TThostFtdcDeliveryModeType = '2' // 实物交割

// 日志级别类型
type TThostFtdcLogLevelType [33]byte

// 存储过程名称类型
type TThostFtdcProcessNameType [257]byte

// 操作摘要类型
type TThostFtdcOperationMemoType [1025]byte

// 出入金类型类型
type TThostFtdcFundIOTypeType byte

const THOST_FTDC_FIOT_FundIO TThostFtdcFundIOTypeType = '1'       // 出入金
const THOST_FTDC_FIOT_Transfer TThostFtdcFundIOTypeType = '2'     // 银期转帐
const THOST_FTDC_FIOT_SwapCurrency TThostFtdcFundIOTypeType = '3' // 银期换汇

// 资金类型类型
type TThostFtdcFundTypeType byte

const THOST_FTDC_FT_Deposite TThostFtdcFundTypeType = '1'      // 银行存款
const THOST_FTDC_FT_ItemFund TThostFtdcFundTypeType = '2'      // 分项资金
const THOST_FTDC_FT_Company TThostFtdcFundTypeType = '3'       // 公司调整
const THOST_FTDC_FT_InnerTransfer TThostFtdcFundTypeType = '4' // 资金内转

// 出入金方向类型
type TThostFtdcFundDirectionType byte

const THOST_FTDC_FD_In TThostFtdcFundDirectionType = '1'  // 入金
const THOST_FTDC_FD_Out TThostFtdcFundDirectionType = '2' // 出金

// 资金状态类型
type TThostFtdcFundStatusType byte

const THOST_FTDC_FS_Record TThostFtdcFundStatusType = '1' // 已录入
const THOST_FTDC_FS_Check TThostFtdcFundStatusType = '2'  // 已复核
const THOST_FTDC_FS_Charge TThostFtdcFundStatusType = '3' // 已冲销

// 票据号类型
type TThostFtdcBillNoType [15]byte

// 票据名称类型
type TThostFtdcBillNameType [33]byte

// 发布状态类型
type TThostFtdcPublishStatusType byte

const THOST_FTDC_PS_None TThostFtdcPublishStatusType = '1'       // 未发布
const THOST_FTDC_PS_Publishing TThostFtdcPublishStatusType = '2' // 正在发布
const THOST_FTDC_PS_Published TThostFtdcPublishStatusType = '3'  // 已发布

// 枚举值代码类型
type TThostFtdcEnumValueIDType [65]byte

// 枚举值类型类型
type TThostFtdcEnumValueTypeType [33]byte

// 枚举值名称类型
type TThostFtdcEnumValueLabelType [65]byte

// 枚举值结果类型
type TThostFtdcEnumValueResultType [33]byte

// 系统状态类型
type TThostFtdcSystemStatusType byte

const THOST_FTDC_ES_NonActive TThostFtdcSystemStatusType = '1'   // 不活跃
const THOST_FTDC_ES_Startup TThostFtdcSystemStatusType = '2'     // 启动
const THOST_FTDC_ES_Initialize TThostFtdcSystemStatusType = '3'  // 交易开始初始化
const THOST_FTDC_ES_Initialized TThostFtdcSystemStatusType = '4' // 交易完成初始化
const THOST_FTDC_ES_Close TThostFtdcSystemStatusType = '5'       // 收市开始
const THOST_FTDC_ES_Closed TThostFtdcSystemStatusType = '6'      // 收市完成
const THOST_FTDC_ES_Settlement TThostFtdcSystemStatusType = '7'  // 结算

// 结算状态类型
type TThostFtdcSettlementStatusType byte

const THOST_FTDC_STS_Initialize TThostFtdcSettlementStatusType = '0'    // 初始
const THOST_FTDC_STS_Settlementing TThostFtdcSettlementStatusType = '1' // 结算中
const THOST_FTDC_STS_Settlemented TThostFtdcSettlementStatusType = '2'  // 已结算
const THOST_FTDC_STS_Finished TThostFtdcSettlementStatusType = '3'      // 结算完成

// 限定值类型类型
type TThostFtdcRangeIntTypeType [33]byte

// 限定值下限类型
type TThostFtdcRangeIntFromType [33]byte

// 限定值上限类型
type TThostFtdcRangeIntToType [33]byte

// 功能代码类型
type TThostFtdcFunctionIDType [25]byte

// 功能编码类型
type TThostFtdcFunctionValueCodeType [257]byte

// 功能名称类型
type TThostFtdcFunctionNameType [65]byte

// 角色编号类型
type TThostFtdcRoleIDType [11]byte

// 角色名称类型
type TThostFtdcRoleNameType [41]byte

// 描述类型
type TThostFtdcDescriptionType [401]byte

// 组合编号类型
type TThostFtdcCombineIDType [25]byte

// 组合类型类型
type TThostFtdcCombineTypeType [25]byte

// 投资者类型类型
type TThostFtdcInvestorTypeType byte

const THOST_FTDC_CT_Person TThostFtdcInvestorTypeType = '0'       // 自然人
const THOST_FTDC_CT_Company TThostFtdcInvestorTypeType = '1'      // 法人
const THOST_FTDC_CT_Fund TThostFtdcInvestorTypeType = '2'         // 投资基金
const THOST_FTDC_CT_SpecialOrgan TThostFtdcInvestorTypeType = '3' // 特殊法人
const THOST_FTDC_CT_Asset TThostFtdcInvestorTypeType = '4'        // 资管户

// 经纪公司类型类型
type TThostFtdcBrokerTypeType byte

const THOST_FTDC_BT_Trade TThostFtdcBrokerTypeType = '0'       // 交易会员
const THOST_FTDC_BT_TradeSettle TThostFtdcBrokerTypeType = '1' // 交易结算会员

// 风险等级类型
type TThostFtdcRiskLevelType byte

const THOST_FTDC_FAS_Low TThostFtdcRiskLevelType = '1'    // 低风险客户
const THOST_FTDC_FAS_Normal TThostFtdcRiskLevelType = '2' // 普通客户
const THOST_FTDC_FAS_Focus TThostFtdcRiskLevelType = '3'  // 关注客户
const THOST_FTDC_FAS_Risk TThostFtdcRiskLevelType = '4'   // 风险客户

// 手续费收取方式类型
type TThostFtdcFeeAcceptStyleType byte

const THOST_FTDC_FAS_ByTrade TThostFtdcFeeAcceptStyleType = '1' // 按交易收取
const THOST_FTDC_FAS_ByDeliv TThostFtdcFeeAcceptStyleType = '2' // 按交割收取
const THOST_FTDC_FAS_None TThostFtdcFeeAcceptStyleType = '3'    // 不收
const THOST_FTDC_FAS_FixFee TThostFtdcFeeAcceptStyleType = '4'  // 按指定手续费收取

// 密码类型类型
type TThostFtdcPasswordTypeType byte

const THOST_FTDC_PWDT_Trade TThostFtdcPasswordTypeType = '1'   // 交易密码
const THOST_FTDC_PWDT_Account TThostFtdcPasswordTypeType = '2' // 资金密码

// 盈亏算法类型
type TThostFtdcAlgorithmType byte

const THOST_FTDC_AG_All TThostFtdcAlgorithmType = '1'      // 浮盈浮亏都计算
const THOST_FTDC_AG_OnlyLost TThostFtdcAlgorithmType = '2' // 浮盈不计，浮亏计
const THOST_FTDC_AG_OnlyGain TThostFtdcAlgorithmType = '3' // 浮盈计，浮亏不计
const THOST_FTDC_AG_None TThostFtdcAlgorithmType = '4'     // 浮盈浮亏都不计算

// 是否包含平仓盈利类型
type TThostFtdcIncludeCloseProfitType byte

const THOST_FTDC_ICP_Include TThostFtdcIncludeCloseProfitType = '0'    // 包含平仓盈利
const THOST_FTDC_ICP_NotInclude TThostFtdcIncludeCloseProfitType = '2' // 不包含平仓盈利

// 是否受可提比例限制类型
type TThostFtdcAllWithoutTradeType byte

const THOST_FTDC_AWT_Enable TThostFtdcAllWithoutTradeType = '0'       // 无仓无成交不受可提比例限制
const THOST_FTDC_AWT_Disable TThostFtdcAllWithoutTradeType = '2'      // 受可提比例限制
const THOST_FTDC_AWT_NoHoldEnable TThostFtdcAllWithoutTradeType = '3' // 无仓不受可提比例限制

// 盈亏算法说明类型
type TThostFtdcCommentType [31]byte

// 版本号类型
type TThostFtdcVersionType [4]byte

// 交易代码类型
type TThostFtdcTradeCodeType [7]byte

// 交易日期类型
type TThostFtdcTradeDateType [9]byte

// 交易时间类型
type TThostFtdcTradeTimeType [9]byte

// 发起方流水号类型
type TThostFtdcTradeSerialType [9]byte

// 发起方流水号类型
type TThostFtdcTradeSerialNoType int32

// 期货公司代码类型
type TThostFtdcFutureIDType [11]byte

// 银行代码类型
type TThostFtdcBankIDType [4]byte

// 银行分中心代码类型
type TThostFtdcBankBrchIDType [5]byte

// 分中心代码类型
type TThostFtdcBankBranchIDType [11]byte

// 交易柜员类型
type TThostFtdcOperNoType [17]byte

// 渠道标志类型
type TThostFtdcDeviceIDType [3]byte

// 记录数类型
type TThostFtdcRecordNumType [7]byte

// 期货资金账号类型
type TThostFtdcFutureAccountType [22]byte

// 资金密码核对标志类型
type TThostFtdcFuturePwdFlagType byte

const THOST_FTDC_FPWD_UnCheck TThostFtdcFuturePwdFlagType = '0' // 不核对
const THOST_FTDC_FPWD_Check TThostFtdcFuturePwdFlagType = '1'   // 核对

// 银期转账类型类型
type TThostFtdcTransferTypeType byte

const THOST_FTDC_TT_BankToFuture TThostFtdcTransferTypeType = '0' // 银行转期货
const THOST_FTDC_TT_FutureToBank TThostFtdcTransferTypeType = '1' // 期货转银行

// 期货资金密码类型
type TThostFtdcFutureAccPwdType [17]byte

// 币种类型
type TThostFtdcCurrencyCodeType [4]byte

// 响应代码类型
type TThostFtdcRetCodeType [5]byte

// 响应信息类型
type TThostFtdcRetInfoType [129]byte

// 银行总余额类型
type TThostFtdcTradeAmtType [20]byte

// 银行可用余额类型
type TThostFtdcUseAmtType [20]byte

// 银行可取余额类型
type TThostFtdcFetchAmtType [20]byte

// 转账有效标志类型
type TThostFtdcTransferValidFlagType byte

const THOST_FTDC_TVF_Invalid TThostFtdcTransferValidFlagType = '0' // 无效或失败
const THOST_FTDC_TVF_Valid TThostFtdcTransferValidFlagType = '1'   // 有效
const THOST_FTDC_TVF_Reverse TThostFtdcTransferValidFlagType = '2' // 冲正

// 证件号码类型
type TThostFtdcCertCodeType [21]byte

// 事由类型
type TThostFtdcReasonType byte

const THOST_FTDC_RN_CD TThostFtdcReasonType = '0' // 错单
const THOST_FTDC_RN_ZT TThostFtdcReasonType = '1' // 资金在途
const THOST_FTDC_RN_QT TThostFtdcReasonType = '2' // 其它

// 资金项目编号类型
type TThostFtdcFundProjectIDType [5]byte

// 性别类型
type TThostFtdcSexType byte

const THOST_FTDC_SEX_None TThostFtdcSexType = '0'  // 未知
const THOST_FTDC_SEX_Man TThostFtdcSexType = '1'   // 男
const THOST_FTDC_SEX_Woman TThostFtdcSexType = '2' // 女

// 职业类型
type TThostFtdcProfessionType [101]byte

// 国籍类型
type TThostFtdcNationalType [31]byte

// 省类型
type TThostFtdcProvinceType [51]byte

// 区类型
type TThostFtdcRegionType [16]byte

// 国家类型
type TThostFtdcCountryType [16]byte

// 营业执照类型
type TThostFtdcLicenseNOType [33]byte

// 企业性质类型
type TThostFtdcCompanyTypeType [16]byte

// 经营范围类型
type TThostFtdcBusinessScopeType [1001]byte

// 注册资本币种类型
type TThostFtdcCapitalCurrencyType [4]byte

// 用户类型类型
type TThostFtdcUserTypeType byte

const THOST_FTDC_UT_Investor TThostFtdcUserTypeType = '0'  // 投资者
const THOST_FTDC_UT_Operator TThostFtdcUserTypeType = '1'  // 操作员
const THOST_FTDC_UT_SuperUser TThostFtdcUserTypeType = '2' // 管理员

// 营业部编号类型
type TThostFtdcBranchIDType [9]byte

// 费率类型类型
type TThostFtdcRateTypeType byte

const THOST_FTDC_RATETYPE_MarginRate TThostFtdcRateTypeType = '2' // 保证金率

// 通知类型类型
type TThostFtdcNoteTypeType byte

const THOST_FTDC_NOTETYPE_TradeSettleBill TThostFtdcNoteTypeType = '1'  // 交易结算单
const THOST_FTDC_NOTETYPE_TradeSettleMonth TThostFtdcNoteTypeType = '2' // 交易结算月报
const THOST_FTDC_NOTETYPE_CallMarginNotes TThostFtdcNoteTypeType = '3'  // 追加保证金通知书
const THOST_FTDC_NOTETYPE_ForceCloseNotes TThostFtdcNoteTypeType = '4'  // 强行平仓通知书
const THOST_FTDC_NOTETYPE_TradeNotes TThostFtdcNoteTypeType = '5'       // 成交通知书
const THOST_FTDC_NOTETYPE_DelivNotes TThostFtdcNoteTypeType = '6'       // 交割通知书

// 结算单方式类型
type TThostFtdcSettlementStyleType byte

const THOST_FTDC_SBS_Day TThostFtdcSettlementStyleType = '1'    // 逐日盯市
const THOST_FTDC_SBS_Volume TThostFtdcSettlementStyleType = '2' // 逐笔对冲

// 域名类型
type TThostFtdcBrokerDNSType [256]byte

// 语句类型
type TThostFtdcSentenceType [501]byte

// 结算单类型类型
type TThostFtdcSettlementBillTypeType byte

const THOST_FTDC_ST_Day TThostFtdcSettlementBillTypeType = '0'   // 日报
const THOST_FTDC_ST_Month TThostFtdcSettlementBillTypeType = '1' // 月报

// 客户权限类型类型
type TThostFtdcUserRightTypeType byte

const THOST_FTDC_URT_Logon TThostFtdcUserRightTypeType = '1'          // 登录
const THOST_FTDC_URT_Transfer TThostFtdcUserRightTypeType = '2'       // 银期转帐
const THOST_FTDC_URT_EMail TThostFtdcUserRightTypeType = '3'          // 邮寄结算单
const THOST_FTDC_URT_Fax TThostFtdcUserRightTypeType = '4'            // 传真结算单
const THOST_FTDC_URT_ConditionOrder TThostFtdcUserRightTypeType = '5' // 条件单

// 保证金价格类型类型
type TThostFtdcMarginPriceTypeType byte

const THOST_FTDC_MPT_PreSettlementPrice TThostFtdcMarginPriceTypeType = '1' // 昨结算价
const THOST_FTDC_MPT_SettlementPrice TThostFtdcMarginPriceTypeType = '2'    // 最新价
const THOST_FTDC_MPT_AveragePrice TThostFtdcMarginPriceTypeType = '3'       // 成交均价
const THOST_FTDC_MPT_OpenPrice TThostFtdcMarginPriceTypeType = '4'          // 开仓价

// 结算单生成状态类型
type TThostFtdcBillGenStatusType byte

const THOST_FTDC_BGS_None TThostFtdcBillGenStatusType = '0'        // 未生成
const THOST_FTDC_BGS_NoGenerated TThostFtdcBillGenStatusType = '1' // 生成中
const THOST_FTDC_BGS_Generated TThostFtdcBillGenStatusType = '2'   // 已生成

// 算法类型类型
type TThostFtdcAlgoTypeType byte

const THOST_FTDC_AT_HandlePositionAlgo TThostFtdcAlgoTypeType = '1' // 持仓处理算法
const THOST_FTDC_AT_FindMarginRateAlgo TThostFtdcAlgoTypeType = '2' // 寻找保证金率算法

// 持仓处理算法编号类型
type TThostFtdcHandlePositionAlgoIDType byte

const THOST_FTDC_HPA_Base TThostFtdcHandlePositionAlgoIDType = '1' // 基本
const THOST_FTDC_HPA_DCE TThostFtdcHandlePositionAlgoIDType = '2'  // 大连商品交易所
const THOST_FTDC_HPA_CZCE TThostFtdcHandlePositionAlgoIDType = '3' // 郑州商品交易所

// 寻找保证金率算法编号类型
type TThostFtdcFindMarginRateAlgoIDType byte

const THOST_FTDC_FMRA_Base TThostFtdcFindMarginRateAlgoIDType = '1' // 基本
const THOST_FTDC_FMRA_DCE TThostFtdcFindMarginRateAlgoIDType = '2'  // 大连商品交易所
const THOST_FTDC_FMRA_CZCE TThostFtdcFindMarginRateAlgoIDType = '3' // 郑州商品交易所

// 资金处理算法编号类型
type TThostFtdcHandleTradingAccountAlgoIDType byte

const THOST_FTDC_HTAA_Base TThostFtdcHandleTradingAccountAlgoIDType = '1' // 基本
const THOST_FTDC_HTAA_DCE TThostFtdcHandleTradingAccountAlgoIDType = '2'  // 大连商品交易所
const THOST_FTDC_HTAA_CZCE TThostFtdcHandleTradingAccountAlgoIDType = '3' // 郑州商品交易所

// 联系人类型类型
type TThostFtdcPersonTypeType byte

const THOST_FTDC_PST_Order TThostFtdcPersonTypeType = '1'              // 指定下单人
const THOST_FTDC_PST_Open TThostFtdcPersonTypeType = '2'               // 开户授权人
const THOST_FTDC_PST_Fund TThostFtdcPersonTypeType = '3'               // 资金调拨人
const THOST_FTDC_PST_Settlement TThostFtdcPersonTypeType = '4'         // 结算单确认人
const THOST_FTDC_PST_Company TThostFtdcPersonTypeType = '5'            // 法人
const THOST_FTDC_PST_Corporation TThostFtdcPersonTypeType = '6'        // 法人代表
const THOST_FTDC_PST_LinkMan TThostFtdcPersonTypeType = '7'            // 投资者联系人
const THOST_FTDC_PST_Ledger TThostFtdcPersonTypeType = '8'             // 分户管理资产负责人
const THOST_FTDC_PST_Trustee TThostFtdcPersonTypeType = '9'            // 托（保）管人
const THOST_FTDC_PST_TrusteeCorporation TThostFtdcPersonTypeType = 'A' // 托（保）管机构法人代表
const THOST_FTDC_PST_TrusteeOpen TThostFtdcPersonTypeType = 'B'        // 托（保）管机构开户授权人
const THOST_FTDC_PST_TrusteeContact TThostFtdcPersonTypeType = 'C'     // 托（保）管机构联系人
const THOST_FTDC_PST_ForeignerRefer TThostFtdcPersonTypeType = 'D'     // 境外自然人参考证件
const THOST_FTDC_PST_CorporationRefer TThostFtdcPersonTypeType = 'E'   // 法人代表参考证件

// 查询范围类型
type TThostFtdcQueryInvestorRangeType byte

const THOST_FTDC_QIR_All TThostFtdcQueryInvestorRangeType = '1'    // 所有
const THOST_FTDC_QIR_Group TThostFtdcQueryInvestorRangeType = '2'  // 查询分类
const THOST_FTDC_QIR_Single TThostFtdcQueryInvestorRangeType = '3' // 单一投资者

// 投资者风险状态类型
type TThostFtdcInvestorRiskStatusType byte

const THOST_FTDC_IRS_Normal TThostFtdcInvestorRiskStatusType = '1'    // 正常
const THOST_FTDC_IRS_Warn TThostFtdcInvestorRiskStatusType = '2'      // 警告
const THOST_FTDC_IRS_Call TThostFtdcInvestorRiskStatusType = '3'      // 追保
const THOST_FTDC_IRS_Force TThostFtdcInvestorRiskStatusType = '4'     // 强平
const THOST_FTDC_IRS_Exception TThostFtdcInvestorRiskStatusType = '5' // 异常

// 单腿编号类型
type TThostFtdcLegIDType int32

// 单腿乘数类型
type TThostFtdcLegMultipleType int32

// 派生层数类型
type TThostFtdcImplyLevelType int32

// 结算账户类型
type TThostFtdcClearAccountType [33]byte

// 结算账户类型
type TThostFtdcOrganNOType [6]byte

// 结算账户联行号类型
type TThostFtdcClearbarchIDType [6]byte

// 用户事件类型类型
type TThostFtdcUserEventTypeType byte

const THOST_FTDC_UET_Login TThostFtdcUserEventTypeType = '1'          // 登录
const THOST_FTDC_UET_Logout TThostFtdcUserEventTypeType = '2'         // 登出
const THOST_FTDC_UET_Trading TThostFtdcUserEventTypeType = '3'        // 交易成功
const THOST_FTDC_UET_TradingError TThostFtdcUserEventTypeType = '4'   // 交易失败
const THOST_FTDC_UET_UpdatePassword TThostFtdcUserEventTypeType = '5' // 修改密码
const THOST_FTDC_UET_Authenticate TThostFtdcUserEventTypeType = '6'   // 客户端认证
const THOST_FTDC_UET_SubmitSysInfo TThostFtdcUserEventTypeType = '7'  // 终端信息上报
const THOST_FTDC_UET_Transfer TThostFtdcUserEventTypeType = '8'       // 转账
const THOST_FTDC_UET_Other TThostFtdcUserEventTypeType = '9'          // 其他

// 用户事件信息类型
type TThostFtdcUserEventInfoType [1025]byte

// 平仓方式类型
type TThostFtdcCloseStyleType byte

const THOST_FTDC_ICS_Close TThostFtdcCloseStyleType = '0'      // 先开先平
const THOST_FTDC_ICS_CloseToday TThostFtdcCloseStyleType = '1' // 先平今再平昨

// 统计方式类型
type TThostFtdcStatModeType byte

const THOST_FTDC_SM_Non TThostFtdcStatModeType = '0'        // ----
const THOST_FTDC_SM_Instrument TThostFtdcStatModeType = '1' // 按合约统计
const THOST_FTDC_SM_Product TThostFtdcStatModeType = '2'    // 按产品统计
const THOST_FTDC_SM_Investor TThostFtdcStatModeType = '3'   // 按投资者统计

// 预埋单状态类型
type TThostFtdcParkedOrderStatusType byte

const THOST_FTDC_PAOS_NotSend TThostFtdcParkedOrderStatusType = '1' // 未发送
const THOST_FTDC_PAOS_Send TThostFtdcParkedOrderStatusType = '2'    // 已发送
const THOST_FTDC_PAOS_Deleted TThostFtdcParkedOrderStatusType = '3' // 已删除

// 预埋报单编号类型
type TThostFtdcParkedOrderIDType [13]byte

// 预埋撤单编号类型
type TThostFtdcParkedOrderActionIDType [13]byte

// 处理状态类型
type TThostFtdcVirDealStatusType byte

const THOST_FTDC_VDS_Dealing TThostFtdcVirDealStatusType = '1'      // 正在处理
const THOST_FTDC_VDS_DeaclSucceed TThostFtdcVirDealStatusType = '2' // 处理成功

// 原有系统代码类型
type TThostFtdcOrgSystemIDType byte

const THOST_FTDC_ORGS_Standard TThostFtdcOrgSystemIDType = '0'   // 综合交易平台
const THOST_FTDC_ORGS_ESunny TThostFtdcOrgSystemIDType = '1'     // 易盛系统
const THOST_FTDC_ORGS_KingStarV6 TThostFtdcOrgSystemIDType = '2' // 金仕达V6系统

// 交易状态类型
type TThostFtdcVirTradeStatusType byte

const THOST_FTDC_VTS_NaturalDeal TThostFtdcVirTradeStatusType = '0'  // 正常处理中
const THOST_FTDC_VTS_SucceedEnd TThostFtdcVirTradeStatusType = '1'   // 成功结束
const THOST_FTDC_VTS_FailedEND TThostFtdcVirTradeStatusType = '2'    // 失败结束
const THOST_FTDC_VTS_Exception TThostFtdcVirTradeStatusType = '3'    // 异常中
const THOST_FTDC_VTS_ManualDeal TThostFtdcVirTradeStatusType = '4'   // 已人工异常处理
const THOST_FTDC_VTS_MesException TThostFtdcVirTradeStatusType = '5' // 通讯异常 ，请人工处理
const THOST_FTDC_VTS_SysException TThostFtdcVirTradeStatusType = '6' // 系统出错，请人工处理

// 银行帐户类型类型
type TThostFtdcVirBankAccTypeType byte

const THOST_FTDC_VBAT_BankBook TThostFtdcVirBankAccTypeType = '1'   // 存折
const THOST_FTDC_VBAT_BankCard TThostFtdcVirBankAccTypeType = '2'   // 储蓄卡
const THOST_FTDC_VBAT_CreditCard TThostFtdcVirBankAccTypeType = '3' // 信用卡

// 银行帐户类型类型
type TThostFtdcVirementStatusType byte

const THOST_FTDC_VMS_Natural TThostFtdcVirementStatusType = '0'  // 正常
const THOST_FTDC_VMS_Canceled TThostFtdcVirementStatusType = '9' // 销户

// 有效标志类型
type TThostFtdcVirementAvailAbilityType byte

const THOST_FTDC_VAA_NoAvailAbility TThostFtdcVirementAvailAbilityType = '0' // 未确认
const THOST_FTDC_VAA_AvailAbility TThostFtdcVirementAvailAbilityType = '1'   // 有效
const THOST_FTDC_VAA_Repeal TThostFtdcVirementAvailAbilityType = '2'         // 冲正

// 交易代码类型
type TThostFtdcVirementTradeCodeType string

const THOST_FTDC_VTC_BankBankToFuture TThostFtdcVirementTradeCodeType = "102001"   // 银行发起银行资金转期货
const THOST_FTDC_VTC_BankFutureToBank TThostFtdcVirementTradeCodeType = "102002"   // 银行发起期货资金转银行
const THOST_FTDC_VTC_FutureBankToFuture TThostFtdcVirementTradeCodeType = "202001" // 期货发起银行资金转期货
const THOST_FTDC_VTC_FutureFutureToBank TThostFtdcVirementTradeCodeType = "202002" // 期货发起期货资金转银行

// 影像类型名称类型
type TThostFtdcPhotoTypeNameType [41]byte

// 影像类型代码类型
type TThostFtdcPhotoTypeIDType [5]byte

// 影像名称类型
type TThostFtdcPhotoNameType [161]byte

// 主题代码类型
type TThostFtdcTopicIDType int32

// 交易报告类型标识类型
type TThostFtdcReportTypeIDType [3]byte

// 交易特征代码类型
type TThostFtdcCharacterIDType [5]byte

// 参数代码类型
type TThostFtdcAMLParamIDType [21]byte

// 投资者类型类型
type TThostFtdcAMLInvestorTypeType [3]byte

// 证件类型类型
type TThostFtdcAMLIdCardTypeType [3]byte

// 资金进出方向类型
type TThostFtdcAMLTradeDirectType [3]byte

// 资金进出方式类型
type TThostFtdcAMLTradeModelType [3]byte

// 业务参数代码值类型
type TThostFtdcAMLOpParamValueType float64

// 客户身份证件/证明文件类型类型
type TThostFtdcAMLCustomerCardTypeType [81]byte

// 金融机构网点名称类型
type TThostFtdcAMLInstitutionNameType [65]byte

// 金融机构网点所在地区行政区划代码类型
type TThostFtdcAMLDistrictIDType [7]byte

// 金融机构网点与大额交易的关系类型
type TThostFtdcAMLRelationShipType [3]byte

// 金融机构网点代码类型类型
type TThostFtdcAMLInstitutionTypeType [3]byte

// 金融机构网点代码类型
type TThostFtdcAMLInstitutionIDType [13]byte

// 账户类型类型
type TThostFtdcAMLAccountTypeType [5]byte

// 交易方式类型
type TThostFtdcAMLTradingTypeType [7]byte

// 涉外收支交易分类与代码类型
type TThostFtdcAMLTransactClassType [7]byte

// 资金收付标识类型
type TThostFtdcAMLCapitalIOType [3]byte

// 交易地点类型
type TThostFtdcAMLSiteType [10]byte

// 资金用途类型
type TThostFtdcAMLCapitalPurposeType [129]byte

// 报文类型类型
type TThostFtdcAMLReportTypeType [2]byte

// 编号类型
type TThostFtdcAMLSerialNoType [5]byte

// 状态类型
type TThostFtdcAMLStatusType [2]byte

// Aml生成方式类型
type TThostFtdcAMLGenStatusType byte

const THOST_FTDC_GEN_Program TThostFtdcAMLGenStatusType = '0'  // 程序生成
const THOST_FTDC_GEN_HandWork TThostFtdcAMLGenStatusType = '1' // 人工生成

// 业务标识号类型
type TThostFtdcAMLSeqCodeType [65]byte

// AML文件名类型
type TThostFtdcAMLFileNameType [257]byte

// 反洗钱资金类型
type TThostFtdcAMLMoneyType float64

// 反洗钱资金类型
type TThostFtdcAMLFileAmountType int32

// 密钥类型(保证金监管)类型
type TThostFtdcCFMMCKeyType [21]byte

// 令牌类型(保证金监管)类型
type TThostFtdcCFMMCTokenType [21]byte

// 动态密钥类别(保证金监管)类型
type TThostFtdcCFMMCKeyKindType byte

const THOST_FTDC_CFMMCKK_REQUEST TThostFtdcCFMMCKeyKindType = 'R' // 主动请求更新
const THOST_FTDC_CFMMCKK_AUTO TThostFtdcCFMMCKeyKindType = 'A'    // CFMMC自动更新
const THOST_FTDC_CFMMCKK_MANUAL TThostFtdcCFMMCKeyKindType = 'M'  // CFMMC手动更新

// 报文名称类型
type TThostFtdcAMLReportNameType [81]byte

// 个人姓名类型
type TThostFtdcIndividualNameType [51]byte

// 币种代码类型
type TThostFtdcCurrencyIDType [4]byte

// 客户编号类型
type TThostFtdcCustNumberType [36]byte

// 机构编码类型
type TThostFtdcOrganCodeType [36]byte

// 机构名称类型
type TThostFtdcOrganNameType [71]byte

// 上级机构编码,即期货公司总部、银行总行类型
type TThostFtdcSuperOrganCodeType [12]byte

// 分支机构类型
type TThostFtdcSubBranchIDType [31]byte

// 分支机构名称类型
type TThostFtdcSubBranchNameType [71]byte

// 机构网点号类型
type TThostFtdcBranchNetCodeType [31]byte

// 机构网点名称类型
type TThostFtdcBranchNetNameType [71]byte

// 机构标识类型
type TThostFtdcOrganFlagType [2]byte

// 银行对期货公司的编码类型
type TThostFtdcBankCodingForFutureType [33]byte

// 银行对返回码的定义类型
type TThostFtdcBankReturnCodeType [7]byte

// 银期转帐平台对返回码的定义类型
type TThostFtdcPlateReturnCodeType [5]byte

// 银行分支机构编码类型
type TThostFtdcBankSubBranchIDType [31]byte

// 期货分支机构编码类型
type TThostFtdcFutureBranchIDType [31]byte

// 返回代码类型
type TThostFtdcReturnCodeType [7]byte

// 操作员类型
type TThostFtdcOperatorCodeType [17]byte

// 机构结算帐户机构号类型
type TThostFtdcClearDepIDType [6]byte

// 机构结算帐户联行号类型
type TThostFtdcClearBrchIDType [6]byte

// 机构结算帐户名称类型
type TThostFtdcClearNameType [71]byte

// 银行帐户名称类型
type TThostFtdcBankAccountNameType [71]byte

// 机构投资人账号机构号类型
type TThostFtdcInvDepIDType [6]byte

// 机构投资人联行号类型
type TThostFtdcInvBrchIDType [6]byte

// 信息格式版本类型
type TThostFtdcMessageFormatVersionType [36]byte

// 摘要类型
type TThostFtdcDigestType [36]byte

// 认证数据类型
type TThostFtdcAuthenticDataType [129]byte

// 密钥类型
type TThostFtdcPasswordKeyType [129]byte

// 期货帐户名称类型
type TThostFtdcFutureAccountNameType [129]byte

// 手机类型
type TThostFtdcMobilePhoneType [21]byte

// 期货公司主密钥类型
type TThostFtdcFutureMainKeyType [129]byte

// 期货公司工作密钥类型
type TThostFtdcFutureWorkKeyType [129]byte

// 期货公司传输密钥类型
type TThostFtdcFutureTransKeyType [129]byte

// 银行主密钥类型
type TThostFtdcBankMainKeyType [129]byte

// 银行工作密钥类型
type TThostFtdcBankWorkKeyType [129]byte

// 银行传输密钥类型
type TThostFtdcBankTransKeyType [129]byte

// 银行服务器描述信息类型
type TThostFtdcBankServerDescriptionType [129]byte

// 附加信息类型
type TThostFtdcAddInfoType [129]byte

// 返回码描述类型
type TThostFtdcDescrInfoForReturnCodeType [129]byte

// 国家代码类型
type TThostFtdcCountryCodeType [21]byte

// 流水号类型
type TThostFtdcSerialType int32

// 平台流水号类型
type TThostFtdcPlateSerialType int32

// 银行流水号类型
type TThostFtdcBankSerialType [13]byte

// 被冲正交易流水号类型
type TThostFtdcCorrectSerialType int32

// 期货公司流水号类型
type TThostFtdcFutureSerialType int32

// 应用标识类型
type TThostFtdcApplicationIDType int32

// 银行代理标识类型
type TThostFtdcBankProxyIDType int32

// 银期转帐核心系统标识类型
type TThostFtdcFBTCoreIDType int32

// 服务端口号类型
type TThostFtdcServerPortType int32

// 已经冲正次数类型
type TThostFtdcRepealedTimesType int32

// 冲正时间间隔类型
type TThostFtdcRepealTimeIntervalType int32

// 每日累计转帐次数类型
type TThostFtdcTotalTimesType int32

// 请求ID类型
type TThostFtdcFBTRequestIDType int32

// 交易ID类型
type TThostFtdcTIDType int32

// 交易金额（元）类型
type TThostFtdcTradeAmountType float64

// 应收客户费用（元）类型
type TThostFtdcCustFeeType float64

// 应收期货公司费用（元）类型
type TThostFtdcFutureFeeType float64

// 单笔最高限额类型
type TThostFtdcSingleMaxAmtType float64

// 单笔最低限额类型
type TThostFtdcSingleMinAmtType float64

// 每日累计转帐额度类型
type TThostFtdcTotalAmtType float64

// 证件类型类型
type TThostFtdcCertificationTypeType byte

const THOST_FTDC_CFT_IDCard TThostFtdcCertificationTypeType = '0'                // 身份证
const THOST_FTDC_CFT_Passport TThostFtdcCertificationTypeType = '1'              // 护照
const THOST_FTDC_CFT_OfficerIDCard TThostFtdcCertificationTypeType = '2'         // 军官证
const THOST_FTDC_CFT_SoldierIDCard TThostFtdcCertificationTypeType = '3'         // 士兵证
const THOST_FTDC_CFT_HomeComingCard TThostFtdcCertificationTypeType = '4'        // 回乡证
const THOST_FTDC_CFT_HouseholdRegister TThostFtdcCertificationTypeType = '5'     // 户口簿
const THOST_FTDC_CFT_LicenseNo TThostFtdcCertificationTypeType = '6'             // 营业执照号
const THOST_FTDC_CFT_InstitutionCodeCard TThostFtdcCertificationTypeType = '7'   // 组织机构代码证
const THOST_FTDC_CFT_TempLicenseNo TThostFtdcCertificationTypeType = '8'         // 临时营业执照号
const THOST_FTDC_CFT_NoEnterpriseLicenseNo TThostFtdcCertificationTypeType = '9' // 民办非企业登记证书
const THOST_FTDC_CFT_OtherCard TThostFtdcCertificationTypeType = 'x'             // 其他证件
const THOST_FTDC_CFT_SuperDepAgree TThostFtdcCertificationTypeType = 'a'         // 主管部门批文

// 文件业务功能类型
type TThostFtdcFileBusinessCodeType byte

const THOST_FTDC_FBC_Others TThostFtdcFileBusinessCodeType = '0'                         // 其他
const THOST_FTDC_FBC_TransferDetails TThostFtdcFileBusinessCodeType = '1'                // 转账交易明细对账
const THOST_FTDC_FBC_CustAccStatus TThostFtdcFileBusinessCodeType = '2'                  // 客户账户状态对账
const THOST_FTDC_FBC_AccountTradeDetails TThostFtdcFileBusinessCodeType = '3'            // 账户类交易明细对账
const THOST_FTDC_FBC_FutureAccountChangeInfoDetails TThostFtdcFileBusinessCodeType = '4' // 期货账户信息变更明细对账
const THOST_FTDC_FBC_CustMoneyDetail TThostFtdcFileBusinessCodeType = '5'                // 客户资金台账余额明细对账
const THOST_FTDC_FBC_CustCancelAccountInfo TThostFtdcFileBusinessCodeType = '6'          // 客户销户结息明细对账
const THOST_FTDC_FBC_CustMoneyResult TThostFtdcFileBusinessCodeType = '7'                // 客户资金余额对账结果
const THOST_FTDC_FBC_OthersExceptionResult TThostFtdcFileBusinessCodeType = '8'          // 其它对账异常结果文件
const THOST_FTDC_FBC_CustInterestNetMoneyDetails TThostFtdcFileBusinessCodeType = '9'    // 客户结息净额明细
const THOST_FTDC_FBC_CustMoneySendAndReceiveDetails TThostFtdcFileBusinessCodeType = 'a' // 客户资金交收明细
const THOST_FTDC_FBC_CorporationMoneyTotal TThostFtdcFileBusinessCodeType = 'b'          // 法人存管银行资金交收汇总
const THOST_FTDC_FBC_MainbodyMoneyTotal TThostFtdcFileBusinessCodeType = 'c'             // 主体间资金交收汇总
const THOST_FTDC_FBC_MainPartMonitorData TThostFtdcFileBusinessCodeType = 'd'            // 总分平衡监管数据
const THOST_FTDC_FBC_PreparationMoney TThostFtdcFileBusinessCodeType = 'e'               // 存管银行备付金余额
const THOST_FTDC_FBC_BankMoneyMonitorData TThostFtdcFileBusinessCodeType = 'f'           // 协办存管银行资金监管数据

// 汇钞标志类型
type TThostFtdcCashExchangeCodeType byte

const THOST_FTDC_CEC_Exchange TThostFtdcCashExchangeCodeType = '1' // 汇
const THOST_FTDC_CEC_Cash TThostFtdcCashExchangeCodeType = '2'     // 钞

// 是或否标识类型
type TThostFtdcYesNoIndicatorType byte

const THOST_FTDC_YNI_Yes TThostFtdcYesNoIndicatorType = '0' // 是
const THOST_FTDC_YNI_No TThostFtdcYesNoIndicatorType = '1'  // 否

// 余额类型类型
type TThostFtdcBanlanceTypeType byte

const THOST_FTDC_BLT_CurrentMoney TThostFtdcBanlanceTypeType = '0'   // 当前余额
const THOST_FTDC_BLT_UsableMoney TThostFtdcBanlanceTypeType = '1'    // 可用余额
const THOST_FTDC_BLT_FetchableMoney TThostFtdcBanlanceTypeType = '2' // 可取余额
const THOST_FTDC_BLT_FreezeMoney TThostFtdcBanlanceTypeType = '3'    // 冻结余额

// 性别类型
type TThostFtdcGenderType byte

const THOST_FTDC_GD_Unknown TThostFtdcGenderType = '0' // 未知状态
const THOST_FTDC_GD_Male TThostFtdcGenderType = '1'    // 男
const THOST_FTDC_GD_Female TThostFtdcGenderType = '2'  // 女

// 费用支付标志类型
type TThostFtdcFeePayFlagType byte

const THOST_FTDC_FPF_BEN TThostFtdcFeePayFlagType = '0' // 由受益方支付费用
const THOST_FTDC_FPF_OUR TThostFtdcFeePayFlagType = '1' // 由发送方支付费用
const THOST_FTDC_FPF_SHA TThostFtdcFeePayFlagType = '2' // 由发送方支付发起的费用，受益方支付接受的费用

// 密钥类型类型
type TThostFtdcPassWordKeyTypeType byte

const THOST_FTDC_PWKT_ExchangeKey TThostFtdcPassWordKeyTypeType = '0' // 交换密钥
const THOST_FTDC_PWKT_PassWordKey TThostFtdcPassWordKeyTypeType = '1' // 密码密钥
const THOST_FTDC_PWKT_MACKey TThostFtdcPassWordKeyTypeType = '2'      // MAC密钥
const THOST_FTDC_PWKT_MessageKey TThostFtdcPassWordKeyTypeType = '3'  // 报文密钥

// 密码类型类型
type TThostFtdcFBTPassWordTypeType byte

const THOST_FTDC_PWT_Query TThostFtdcFBTPassWordTypeType = '0'    // 查询
const THOST_FTDC_PWT_Fetch TThostFtdcFBTPassWordTypeType = '1'    // 取款
const THOST_FTDC_PWT_Transfer TThostFtdcFBTPassWordTypeType = '2' // 转帐
const THOST_FTDC_PWT_Trade TThostFtdcFBTPassWordTypeType = '3'    // 交易

// 加密方式类型
type TThostFtdcFBTEncryModeType byte

const THOST_FTDC_EM_NoEncry TThostFtdcFBTEncryModeType = '0' // 不加密
const THOST_FTDC_EM_DES TThostFtdcFBTEncryModeType = '1'     // DES
const THOST_FTDC_EM_3DES TThostFtdcFBTEncryModeType = '2'    // 3DES

// 银行冲正标志类型
type TThostFtdcBankRepealFlagType byte

const THOST_FTDC_BRF_BankNotNeedRepeal TThostFtdcBankRepealFlagType = '0' // 银行无需自动冲正
const THOST_FTDC_BRF_BankWaitingRepeal TThostFtdcBankRepealFlagType = '1' // 银行待自动冲正
const THOST_FTDC_BRF_BankBeenRepealed TThostFtdcBankRepealFlagType = '2'  // 银行已自动冲正

// 期商冲正标志类型
type TThostFtdcBrokerRepealFlagType byte

const THOST_FTDC_BRORF_BrokerNotNeedRepeal TThostFtdcBrokerRepealFlagType = '0' // 期商无需自动冲正
const THOST_FTDC_BRORF_BrokerWaitingRepeal TThostFtdcBrokerRepealFlagType = '1' // 期商待自动冲正
const THOST_FTDC_BRORF_BrokerBeenRepealed TThostFtdcBrokerRepealFlagType = '2'  // 期商已自动冲正

// 机构类别类型
type TThostFtdcInstitutionTypeType byte

const THOST_FTDC_TS_Bank TThostFtdcInstitutionTypeType = '0'   // 银行
const THOST_FTDC_TS_Future TThostFtdcInstitutionTypeType = '1' // 期商
const THOST_FTDC_TS_Store TThostFtdcInstitutionTypeType = '2'  // 券商

// 最后分片标志类型
type TThostFtdcLastFragmentType byte

const THOST_FTDC_LF_Yes TThostFtdcLastFragmentType = '0' // 是最后分片
const THOST_FTDC_LF_No TThostFtdcLastFragmentType = '1'  // 不是最后分片

// 银行账户状态类型
type TThostFtdcBankAccStatusType byte

const THOST_FTDC_BAS_Normal TThostFtdcBankAccStatusType = '0'     // 正常
const THOST_FTDC_BAS_Freeze TThostFtdcBankAccStatusType = '1'     // 冻结
const THOST_FTDC_BAS_ReportLoss TThostFtdcBankAccStatusType = '2' // 挂失

// 资金账户状态类型
type TThostFtdcMoneyAccountStatusType byte

const THOST_FTDC_MAS_Normal TThostFtdcMoneyAccountStatusType = '0' // 正常
const THOST_FTDC_MAS_Cancel TThostFtdcMoneyAccountStatusType = '1' // 销户

// 存管状态类型
type TThostFtdcManageStatusType byte

const THOST_FTDC_MSS_Point TThostFtdcManageStatusType = '0'       // 指定存管
const THOST_FTDC_MSS_PrePoint TThostFtdcManageStatusType = '1'    // 预指定
const THOST_FTDC_MSS_CancelPoint TThostFtdcManageStatusType = '2' // 撤销指定

// 应用系统类型类型
type TThostFtdcSystemTypeType byte

const THOST_FTDC_SYT_FutureBankTransfer TThostFtdcSystemTypeType = '0' // 银期转帐
const THOST_FTDC_SYT_StockBankTransfer TThostFtdcSystemTypeType = '1'  // 银证转帐
const THOST_FTDC_SYT_TheThirdPartStore TThostFtdcSystemTypeType = '2'  // 第三方存管

// 银期转帐划转结果标志类型
type TThostFtdcTxnEndFlagType byte

const THOST_FTDC_TEF_NormalProcessing TThostFtdcTxnEndFlagType = '0'             // 正常处理中
const THOST_FTDC_TEF_Success TThostFtdcTxnEndFlagType = '1'                      // 成功结束
const THOST_FTDC_TEF_Failed TThostFtdcTxnEndFlagType = '2'                       // 失败结束
const THOST_FTDC_TEF_Abnormal TThostFtdcTxnEndFlagType = '3'                     // 异常中
const THOST_FTDC_TEF_ManualProcessedForException TThostFtdcTxnEndFlagType = '4'  // 已人工异常处理
const THOST_FTDC_TEF_CommuFailedNeedManualProcess TThostFtdcTxnEndFlagType = '5' // 通讯异常 ，请人工处理
const THOST_FTDC_TEF_SysErrorNeedManualProcess TThostFtdcTxnEndFlagType = '6'    // 系统出错，请人工处理

// 银期转帐服务处理状态类型
type TThostFtdcProcessStatusType byte

const THOST_FTDC_PSS_NotProcess TThostFtdcProcessStatusType = '0'   // 未处理
const THOST_FTDC_PSS_StartProcess TThostFtdcProcessStatusType = '1' // 开始处理
const THOST_FTDC_PSS_Finished TThostFtdcProcessStatusType = '2'     // 处理完成

// 客户类型类型
type TThostFtdcCustTypeType byte

const THOST_FTDC_CUSTT_Person TThostFtdcCustTypeType = '0'      // 自然人
const THOST_FTDC_CUSTT_Institution TThostFtdcCustTypeType = '1' // 机构户

// 银期转帐方向类型
type TThostFtdcFBTTransferDirectionType byte

const THOST_FTDC_FBTTD_FromBankToFuture TThostFtdcFBTTransferDirectionType = '1' // 入金，银行转期货
const THOST_FTDC_FBTTD_FromFutureToBank TThostFtdcFBTTransferDirectionType = '2' // 出金，期货转银行

// 开销户类别类型
type TThostFtdcOpenOrDestroyType byte

const THOST_FTDC_OOD_Open TThostFtdcOpenOrDestroyType = '1'    // 开户
const THOST_FTDC_OOD_Destroy TThostFtdcOpenOrDestroyType = '0' // 销户

// 有效标志类型
type TThostFtdcAvailabilityFlagType byte

const THOST_FTDC_AVAF_Invalid TThostFtdcAvailabilityFlagType = '0' // 未确认
const THOST_FTDC_AVAF_Valid TThostFtdcAvailabilityFlagType = '1'   // 有效
const THOST_FTDC_AVAF_Repeal TThostFtdcAvailabilityFlagType = '2'  // 冲正

// 机构类型类型
type TThostFtdcOrganTypeType byte

const THOST_FTDC_OT_Bank TThostFtdcOrganTypeType = '1'      // 银行代理
const THOST_FTDC_OT_Future TThostFtdcOrganTypeType = '2'    // 交易前置
const THOST_FTDC_OT_PlateForm TThostFtdcOrganTypeType = '9' // 银期转帐平台管理

// 机构级别类型
type TThostFtdcOrganLevelType byte

const THOST_FTDC_OL_HeadQuarters TThostFtdcOrganLevelType = '1' // 银行总行或期商总部
const THOST_FTDC_OL_Branch TThostFtdcOrganLevelType = '2'       // 银行分中心或期货公司营业部

// 协议类型类型
type TThostFtdcProtocalIDType byte

const THOST_FTDC_PID_FutureProtocal TThostFtdcProtocalIDType = '0'       // 期商协议
const THOST_FTDC_PID_ICBCProtocal TThostFtdcProtocalIDType = '1'         // 工行协议
const THOST_FTDC_PID_ABCProtocal TThostFtdcProtocalIDType = '2'          // 农行协议
const THOST_FTDC_PID_CBCProtocal TThostFtdcProtocalIDType = '3'          // 中国银行协议
const THOST_FTDC_PID_CCBProtocal TThostFtdcProtocalIDType = '4'          // 建行协议
const THOST_FTDC_PID_BOCOMProtocal TThostFtdcProtocalIDType = '5'        // 交行协议
const THOST_FTDC_PID_FBTPlateFormProtocal TThostFtdcProtocalIDType = 'X' // 银期转帐平台协议

// 套接字连接方式类型
type TThostFtdcConnectModeType byte

const THOST_FTDC_CM_ShortConnect TThostFtdcConnectModeType = '0' // 短连接
const THOST_FTDC_CM_LongConnect TThostFtdcConnectModeType = '1'  // 长连接

// 套接字通信方式类型
type TThostFtdcSyncModeType byte

const THOST_FTDC_SRM_ASync TThostFtdcSyncModeType = '0' // 异步
const THOST_FTDC_SRM_Sync TThostFtdcSyncModeType = '1'  // 同步

// 银行帐号类型类型
type TThostFtdcBankAccTypeType byte

const THOST_FTDC_BAT_BankBook TThostFtdcBankAccTypeType = '1'   // 银行存折
const THOST_FTDC_BAT_SavingCard TThostFtdcBankAccTypeType = '2' // 储蓄卡
const THOST_FTDC_BAT_CreditCard TThostFtdcBankAccTypeType = '3' // 信用卡

// 期货公司帐号类型类型
type TThostFtdcFutureAccTypeType byte

const THOST_FTDC_FAT_BankBook TThostFtdcFutureAccTypeType = '1'   // 银行存折
const THOST_FTDC_FAT_SavingCard TThostFtdcFutureAccTypeType = '2' // 储蓄卡
const THOST_FTDC_FAT_CreditCard TThostFtdcFutureAccTypeType = '3' // 信用卡

// 接入机构状态类型
type TThostFtdcOrganStatusType byte

const THOST_FTDC_OS_Ready TThostFtdcOrganStatusType = '0'            // 启用
const THOST_FTDC_OS_CheckIn TThostFtdcOrganStatusType = '1'          // 签到
const THOST_FTDC_OS_CheckOut TThostFtdcOrganStatusType = '2'         // 签退
const THOST_FTDC_OS_CheckFileArrived TThostFtdcOrganStatusType = '3' // 对帐文件到达
const THOST_FTDC_OS_CheckDetail TThostFtdcOrganStatusType = '4'      // 对帐
const THOST_FTDC_OS_DayEndClean TThostFtdcOrganStatusType = '5'      // 日终清理
const THOST_FTDC_OS_Invalid TThostFtdcOrganStatusType = '9'          // 注销

// 建行收费模式类型
type TThostFtdcCCBFeeModeType byte

const THOST_FTDC_CCBFM_ByAmount TThostFtdcCCBFeeModeType = '1' // 按金额扣收
const THOST_FTDC_CCBFM_ByMonth TThostFtdcCCBFeeModeType = '2'  // 按月扣收

// 通讯API类型类型
type TThostFtdcCommApiTypeType byte

const THOST_FTDC_CAPIT_Client TThostFtdcCommApiTypeType = '1'  // 客户端
const THOST_FTDC_CAPIT_Server TThostFtdcCommApiTypeType = '2'  // 服务端
const THOST_FTDC_CAPIT_UserApi TThostFtdcCommApiTypeType = '3' // 交易系统的UserApi

// 服务编号类型
type TThostFtdcServiceIDType int32

// 服务线路编号类型
type TThostFtdcServiceLineNoType int32

// 服务名类型
type TThostFtdcServiceNameType [61]byte

// 连接状态类型
type TThostFtdcLinkStatusType byte

const THOST_FTDC_LS_Connected TThostFtdcLinkStatusType = '1'    // 已经连接
const THOST_FTDC_LS_Disconnected TThostFtdcLinkStatusType = '2' // 没有连接

// 通讯API指针类型
type TThostFtdcCommApiPointerType int32

// 密码核对标志类型
type TThostFtdcPwdFlagType byte

const THOST_FTDC_BPWDF_NoCheck TThostFtdcPwdFlagType = '0'      // 不核对
const THOST_FTDC_BPWDF_BlankCheck TThostFtdcPwdFlagType = '1'   // 明文核对
const THOST_FTDC_BPWDF_EncryptCheck TThostFtdcPwdFlagType = '2' // 密文核对

// 期货帐号类型类型
type TThostFtdcSecuAccTypeType byte

const THOST_FTDC_SAT_AccountID TThostFtdcSecuAccTypeType = '1'       // 资金帐号
const THOST_FTDC_SAT_CardID TThostFtdcSecuAccTypeType = '2'          // 资金卡号
const THOST_FTDC_SAT_SHStockholderID TThostFtdcSecuAccTypeType = '3' // 上海股东帐号
const THOST_FTDC_SAT_SZStockholderID TThostFtdcSecuAccTypeType = '4' // 深圳股东帐号

// 转账交易状态类型
type TThostFtdcTransferStatusType byte

const THOST_FTDC_TRFS_Normal TThostFtdcTransferStatusType = '0'   // 正常
const THOST_FTDC_TRFS_Repealed TThostFtdcTransferStatusType = '1' // 被冲正

// 发起方类型
type TThostFtdcSponsorTypeType byte

const THOST_FTDC_SPTYPE_Broker TThostFtdcSponsorTypeType = '0' // 期商
const THOST_FTDC_SPTYPE_Bank TThostFtdcSponsorTypeType = '1'   // 银行

// 请求响应类别类型
type TThostFtdcReqRspTypeType byte

const THOST_FTDC_REQRSP_Request TThostFtdcReqRspTypeType = '0'  // 请求
const THOST_FTDC_REQRSP_Response TThostFtdcReqRspTypeType = '1' // 响应

// 银期转帐用户事件类型类型
type TThostFtdcFBTUserEventTypeType byte

const THOST_FTDC_FBTUET_SignIn TThostFtdcFBTUserEventTypeType = '0'                    // 签到
const THOST_FTDC_FBTUET_FromBankToFuture TThostFtdcFBTUserEventTypeType = '1'          // 银行转期货
const THOST_FTDC_FBTUET_FromFutureToBank TThostFtdcFBTUserEventTypeType = '2'          // 期货转银行
const THOST_FTDC_FBTUET_OpenAccount TThostFtdcFBTUserEventTypeType = '3'               // 开户
const THOST_FTDC_FBTUET_CancelAccount TThostFtdcFBTUserEventTypeType = '4'             // 销户
const THOST_FTDC_FBTUET_ChangeAccount TThostFtdcFBTUserEventTypeType = '5'             // 变更银行账户
const THOST_FTDC_FBTUET_RepealFromBankToFuture TThostFtdcFBTUserEventTypeType = '6'    // 冲正银行转期货
const THOST_FTDC_FBTUET_RepealFromFutureToBank TThostFtdcFBTUserEventTypeType = '7'    // 冲正期货转银行
const THOST_FTDC_FBTUET_QueryBankAccount TThostFtdcFBTUserEventTypeType = '8'          // 查询银行账户
const THOST_FTDC_FBTUET_QueryFutureAccount TThostFtdcFBTUserEventTypeType = '9'        // 查询期货账户
const THOST_FTDC_FBTUET_SignOut TThostFtdcFBTUserEventTypeType = 'A'                   // 签退
const THOST_FTDC_FBTUET_SyncKey TThostFtdcFBTUserEventTypeType = 'B'                   // 密钥同步
const THOST_FTDC_FBTUET_ReserveOpenAccount TThostFtdcFBTUserEventTypeType = 'C'        // 预约开户
const THOST_FTDC_FBTUET_CancelReserveOpenAccount TThostFtdcFBTUserEventTypeType = 'D'  // 撤销预约开户
const THOST_FTDC_FBTUET_ReserveOpenAccountConfirm TThostFtdcFBTUserEventTypeType = 'E' // 预约开户确认
const THOST_FTDC_FBTUET_Other TThostFtdcFBTUserEventTypeType = 'Z'                     // 其他

// 银行自己的编码类型
type TThostFtdcBankIDByBankType [21]byte

// 银行操作员号类型
type TThostFtdcBankOperNoType [4]byte

// 银行客户号类型
type TThostFtdcBankCustNoType [21]byte

// 递增的序列号类型
type TThostFtdcDBOPSeqNoType int32

// FBT表名类型
type TThostFtdcTableNameType [61]byte

// FBT表操作主键名类型
type TThostFtdcPKNameType [201]byte

// FBT表操作主键值类型
type TThostFtdcPKValueType [501]byte

// 记录操作类型类型
type TThostFtdcDBOperationType byte

const THOST_FTDC_DBOP_Insert TThostFtdcDBOperationType = '0' // 插入
const THOST_FTDC_DBOP_Update TThostFtdcDBOperationType = '1' // 更新
const THOST_FTDC_DBOP_Delete TThostFtdcDBOperationType = '2' // 删除

// 同步标记类型
type TThostFtdcSyncFlagType byte

const THOST_FTDC_SYNF_Yes TThostFtdcSyncFlagType = '0' // 已同步
const THOST_FTDC_SYNF_No TThostFtdcSyncFlagType = '1'  // 未同步

// 同步目标编号类型
type TThostFtdcTargetIDType [4]byte

// 同步类型类型
type TThostFtdcSyncTypeType byte

const THOST_FTDC_SYNT_OneOffSync TThostFtdcSyncTypeType = '0'    // 一次同步
const THOST_FTDC_SYNT_TimerSync TThostFtdcSyncTypeType = '1'     // 定时同步
const THOST_FTDC_SYNT_TimerFullSync TThostFtdcSyncTypeType = '2' // 定时完全同步

// 各种换汇时间类型
type TThostFtdcFBETimeType [7]byte

// 换汇银行行号类型
type TThostFtdcFBEBankNoType [13]byte

// 换汇凭证号类型
type TThostFtdcFBECertNoType [13]byte

// 换汇方向类型
type TThostFtdcExDirectionType byte

const THOST_FTDC_FBEDIR_Settlement TThostFtdcExDirectionType = '0' // 结汇
const THOST_FTDC_FBEDIR_Sale TThostFtdcExDirectionType = '1'       // 售汇

// 换汇银行账户类型
type TThostFtdcFBEBankAccountType [33]byte

// 换汇银行账户名类型
type TThostFtdcFBEBankAccountNameType [61]byte

// 各种换汇金额类型
type TThostFtdcFBEAmtType float64

// 换汇业务类型类型
type TThostFtdcFBEBusinessTypeType [3]byte

// 换汇附言类型
type TThostFtdcFBEPostScriptType [61]byte

// 换汇备注类型
type TThostFtdcFBERemarkType [71]byte

// 换汇汇率类型
type TThostFtdcExRateType float64

// 换汇成功标志类型
type TThostFtdcFBEResultFlagType byte

const THOST_FTDC_FBERES_Success TThostFtdcFBEResultFlagType = '0'             // 成功
const THOST_FTDC_FBERES_InsufficientBalance TThostFtdcFBEResultFlagType = '1' // 账户余额不足
const THOST_FTDC_FBERES_UnknownTrading TThostFtdcFBEResultFlagType = '8'      // 交易结果未知
const THOST_FTDC_FBERES_Fail TThostFtdcFBEResultFlagType = 'x'                // 失败

// 换汇返回信息类型
type TThostFtdcFBERtnMsgType [61]byte

// 换汇扩展信息类型
type TThostFtdcFBEExtendMsgType [61]byte

// 换汇记账流水号类型
type TThostFtdcFBEBusinessSerialType [31]byte

// 换汇流水号类型
type TThostFtdcFBESystemSerialType [21]byte

// 换汇交易总笔数类型
type TThostFtdcFBETotalExCntType int32

// 换汇交易状态类型
type TThostFtdcFBEExchStatusType byte

const THOST_FTDC_FBEES_Normal TThostFtdcFBEExchStatusType = '0'     // 正常
const THOST_FTDC_FBEES_ReExchange TThostFtdcFBEExchStatusType = '1' // 交易重发

// 换汇文件标志类型
type TThostFtdcFBEFileFlagType byte

const THOST_FTDC_FBEFG_DataPackage TThostFtdcFBEFileFlagType = '0' // 数据包
const THOST_FTDC_FBEFG_File TThostFtdcFBEFileFlagType = '1'        // 文件

// 换汇已交易标志类型
type TThostFtdcFBEAlreadyTradeType byte

const THOST_FTDC_FBEAT_NotTrade TThostFtdcFBEAlreadyTradeType = '0' // 未交易
const THOST_FTDC_FBEAT_Trade TThostFtdcFBEAlreadyTradeType = '1'    // 已交易

// 换汇账户开户行类型
type TThostFtdcFBEOpenBankType [61]byte

// 银期换汇用户事件类型类型
type TThostFtdcFBEUserEventTypeType byte

const THOST_FTDC_FBEUET_SignIn TThostFtdcFBEUserEventTypeType = '0'           // 签到
const THOST_FTDC_FBEUET_Exchange TThostFtdcFBEUserEventTypeType = '1'         // 换汇
const THOST_FTDC_FBEUET_ReExchange TThostFtdcFBEUserEventTypeType = '2'       // 换汇重发
const THOST_FTDC_FBEUET_QueryBankAccount TThostFtdcFBEUserEventTypeType = '3' // 银行账户查询
const THOST_FTDC_FBEUET_QueryExchDetial TThostFtdcFBEUserEventTypeType = '4'  // 换汇明细查询
const THOST_FTDC_FBEUET_QueryExchSummary TThostFtdcFBEUserEventTypeType = '5' // 换汇汇总查询
const THOST_FTDC_FBEUET_QueryExchRate TThostFtdcFBEUserEventTypeType = '6'    // 换汇汇率查询
const THOST_FTDC_FBEUET_CheckBankAccount TThostFtdcFBEUserEventTypeType = '7' // 对账文件通知
const THOST_FTDC_FBEUET_SignOut TThostFtdcFBEUserEventTypeType = '8'          // 签退
const THOST_FTDC_FBEUET_Other TThostFtdcFBEUserEventTypeType = 'Z'            // 其他

// 换汇相关文件名类型
type TThostFtdcFBEFileNameType [21]byte

// 换汇批次号类型
type TThostFtdcFBEBatchSerialType [21]byte

// 换汇发送标志类型
type TThostFtdcFBEReqFlagType byte

const THOST_FTDC_FBERF_UnProcessed TThostFtdcFBEReqFlagType = '0' // 未处理
const THOST_FTDC_FBERF_WaitSend TThostFtdcFBEReqFlagType = '1'    // 等待发送
const THOST_FTDC_FBERF_SendSuccess TThostFtdcFBEReqFlagType = '2' // 发送成功
const THOST_FTDC_FBERF_SendFailed TThostFtdcFBEReqFlagType = '3'  // 发送失败
const THOST_FTDC_FBERF_WaitReSend TThostFtdcFBEReqFlagType = '4'  // 等待重发

// 风险通知类型类型
type TThostFtdcNotifyClassType byte

const THOST_FTDC_NC_NOERROR TThostFtdcNotifyClassType = '0'   // 正常
const THOST_FTDC_NC_Warn TThostFtdcNotifyClassType = '1'      // 警示
const THOST_FTDC_NC_Call TThostFtdcNotifyClassType = '2'      // 追保
const THOST_FTDC_NC_Force TThostFtdcNotifyClassType = '3'     // 强平
const THOST_FTDC_NC_CHUANCANG TThostFtdcNotifyClassType = '4' // 穿仓
const THOST_FTDC_NC_Exception TThostFtdcNotifyClassType = '5' // 异常

// 客户风险通知消息类型
type TThostFtdcRiskNofityInfoType [257]byte

// 强平场景编号类型
type TThostFtdcForceCloseSceneIdType [24]byte

// 强平单类型类型
type TThostFtdcForceCloseTypeType byte

const THOST_FTDC_FCT_Manual TThostFtdcForceCloseTypeType = '0' // 手工强平
const THOST_FTDC_FCT_Single TThostFtdcForceCloseTypeType = '1' // 单一投资者辅助强平
const THOST_FTDC_FCT_Group TThostFtdcForceCloseTypeType = '2'  // 批量投资者辅助强平

// 多个产品代码,用+分隔,如cu+zn类型
type TThostFtdcInstrumentIDsType [101]byte

// 风险通知途径类型
type TThostFtdcRiskNotifyMethodType byte

const THOST_FTDC_RNM_System TThostFtdcRiskNotifyMethodType = '0' // 系统通知
const THOST_FTDC_RNM_SMS TThostFtdcRiskNotifyMethodType = '1'    // 短信通知
const THOST_FTDC_RNM_EMail TThostFtdcRiskNotifyMethodType = '2'  // 邮件通知
const THOST_FTDC_RNM_Manual TThostFtdcRiskNotifyMethodType = '3' // 人工通知

// 风险通知状态类型
type TThostFtdcRiskNotifyStatusType byte

const THOST_FTDC_RNS_NotGen TThostFtdcRiskNotifyStatusType = '0'    // 未生成
const THOST_FTDC_RNS_Generated TThostFtdcRiskNotifyStatusType = '1' // 已生成未发送
const THOST_FTDC_RNS_SendError TThostFtdcRiskNotifyStatusType = '2' // 发送失败
const THOST_FTDC_RNS_SendOk TThostFtdcRiskNotifyStatusType = '3'    // 已发送未接收
const THOST_FTDC_RNS_Received TThostFtdcRiskNotifyStatusType = '4'  // 已接收未确认
const THOST_FTDC_RNS_Confirmed TThostFtdcRiskNotifyStatusType = '5' // 已确认

// 风控用户操作事件类型
type TThostFtdcRiskUserEventType byte

const THOST_FTDC_RUE_ExportData TThostFtdcRiskUserEventType = '0' // 导出数据

// 参数代码类型
type TThostFtdcParamIDType int32

// 参数名类型
type TThostFtdcParamNameType [41]byte

// 参数值类型
type TThostFtdcParamValueType [41]byte

// 条件单索引条件类型
type TThostFtdcConditionalOrderSortTypeType byte

const THOST_FTDC_COST_LastPriceAsc TThostFtdcConditionalOrderSortTypeType = '0'  // 使用最新价升序
const THOST_FTDC_COST_LastPriceDesc TThostFtdcConditionalOrderSortTypeType = '1' // 使用最新价降序
const THOST_FTDC_COST_AskPriceAsc TThostFtdcConditionalOrderSortTypeType = '2'   // 使用卖价升序
const THOST_FTDC_COST_AskPriceDesc TThostFtdcConditionalOrderSortTypeType = '3'  // 使用卖价降序
const THOST_FTDC_COST_BidPriceAsc TThostFtdcConditionalOrderSortTypeType = '4'   // 使用买价升序
const THOST_FTDC_COST_BidPriceDesc TThostFtdcConditionalOrderSortTypeType = '5'  // 使用买价降序

// 报送状态类型
type TThostFtdcSendTypeType byte

const THOST_FTDC_UOAST_NoSend TThostFtdcSendTypeType = '0'    // 未发送
const THOST_FTDC_UOAST_Sended TThostFtdcSendTypeType = '1'    // 已发送
const THOST_FTDC_UOAST_Generated TThostFtdcSendTypeType = '2' // 已生成
const THOST_FTDC_UOAST_SendFail TThostFtdcSendTypeType = '3'  // 报送失败
const THOST_FTDC_UOAST_Success TThostFtdcSendTypeType = '4'   // 接收成功
const THOST_FTDC_UOAST_Fail TThostFtdcSendTypeType = '5'      // 接收失败
const THOST_FTDC_UOAST_Cancel TThostFtdcSendTypeType = '6'    // 取消报送

// 交易编码状态类型
type TThostFtdcClientIDStatusType byte

const THOST_FTDC_UOACS_NoApply TThostFtdcClientIDStatusType = '1'  // 未申请
const THOST_FTDC_UOACS_Submited TThostFtdcClientIDStatusType = '2' // 已提交申请
const THOST_FTDC_UOACS_Sended TThostFtdcClientIDStatusType = '3'   // 已发送申请
const THOST_FTDC_UOACS_Success TThostFtdcClientIDStatusType = '4'  // 完成
const THOST_FTDC_UOACS_Refuse TThostFtdcClientIDStatusType = '5'   // 拒绝
const THOST_FTDC_UOACS_Cancel TThostFtdcClientIDStatusType = '6'   // 已撤销编码

// 行业编码类型
type TThostFtdcIndustryIDType [17]byte

// 特有信息编号类型
type TThostFtdcQuestionIDType [5]byte

// 特有信息说明类型
type TThostFtdcQuestionContentType [41]byte

// 选项编号类型
type TThostFtdcOptionIDType [13]byte

// 选项说明类型
type TThostFtdcOptionContentType [61]byte

// 特有信息类型类型
type TThostFtdcQuestionTypeType byte

const THOST_FTDC_QT_Radio TThostFtdcQuestionTypeType = '1'  // 单选
const THOST_FTDC_QT_Option TThostFtdcQuestionTypeType = '2' // 多选
const THOST_FTDC_QT_Blank TThostFtdcQuestionTypeType = '3'  // 填空

// 业务流水号类型
type TThostFtdcProcessIDType [33]byte

// 流水号类型
type TThostFtdcSeqNoType int32

// 流程状态类型
type TThostFtdcUOAProcessStatusType [3]byte

// 流程功能类型类型
type TThostFtdcProcessTypeType [3]byte

// 业务类型类型
type TThostFtdcBusinessTypeType byte

const THOST_FTDC_BT_Request TThostFtdcBusinessTypeType = '1'  // 请求
const THOST_FTDC_BT_Response TThostFtdcBusinessTypeType = '2' // 应答
const THOST_FTDC_BT_Notice TThostFtdcBusinessTypeType = '3'   // 通知

// 监控中心返回码类型
type TThostFtdcCfmmcReturnCodeType byte

const THOST_FTDC_CRC_Success TThostFtdcCfmmcReturnCodeType = '0'    // 成功
const THOST_FTDC_CRC_Working TThostFtdcCfmmcReturnCodeType = '1'    // 该客户已经有流程在处理中
const THOST_FTDC_CRC_InfoFail TThostFtdcCfmmcReturnCodeType = '2'   // 监控中客户资料检查失败
const THOST_FTDC_CRC_IDCardFail TThostFtdcCfmmcReturnCodeType = '3' // 监控中实名制检查失败
const THOST_FTDC_CRC_OtherFail TThostFtdcCfmmcReturnCodeType = '4'  // 其他错误

// 交易所返回码类型
type TThostFtdcExReturnCodeType int32

// 客户类型类型
type TThostFtdcClientTypeType byte

const THOST_FTDC_CfMMCCT_All TThostFtdcClientTypeType = '0'          // 所有
const THOST_FTDC_CfMMCCT_Person TThostFtdcClientTypeType = '1'       // 个人
const THOST_FTDC_CfMMCCT_Company TThostFtdcClientTypeType = '2'      // 单位
const THOST_FTDC_CfMMCCT_Other TThostFtdcClientTypeType = '3'        // 其他
const THOST_FTDC_CfMMCCT_SpecialOrgan TThostFtdcClientTypeType = '4' // 特殊法人
const THOST_FTDC_CfMMCCT_Asset TThostFtdcClientTypeType = '5'        // 资管户

// 交易所编号类型
type TThostFtdcExchangeIDTypeType byte

const THOST_FTDC_EIDT_SHFE TThostFtdcExchangeIDTypeType = 'S'  // 上海期货交易所
const THOST_FTDC_EIDT_CZCE TThostFtdcExchangeIDTypeType = 'Z'  // 郑州商品交易所
const THOST_FTDC_EIDT_DCE TThostFtdcExchangeIDTypeType = 'D'   // 大连商品交易所
const THOST_FTDC_EIDT_CFFEX TThostFtdcExchangeIDTypeType = 'J' // 中国金融期货交易所
const THOST_FTDC_EIDT_INE TThostFtdcExchangeIDTypeType = 'N'   // 上海国际能源交易中心股份有限公司

// 交易编码类型类型
type TThostFtdcExClientIDTypeType byte

const THOST_FTDC_ECIDT_Hedge TThostFtdcExClientIDTypeType = '1'       // 套保
const THOST_FTDC_ECIDT_Arbitrage TThostFtdcExClientIDTypeType = '2'   // 套利
const THOST_FTDC_ECIDT_Speculation TThostFtdcExClientIDTypeType = '3' // 投机

// 客户分类码类型
type TThostFtdcClientClassifyType [11]byte

// 单位性质类型
type TThostFtdcUOAOrganTypeType [11]byte

// 国家代码类型
type TThostFtdcUOACountryCodeType [11]byte

// 区号类型
type TThostFtdcAreaCodeType [11]byte

// 监控中心为客户分配的代码类型
type TThostFtdcFuturesIDType [21]byte

// 日期类型
type TThostFtdcCffmcDateType [11]byte

// 时间类型
type TThostFtdcCffmcTimeType [11]byte

// 组织机构代码类型
type TThostFtdcNocIDType [21]byte

// 更新状态类型
type TThostFtdcUpdateFlagType byte

const THOST_FTDC_UF_NoUpdate TThostFtdcUpdateFlagType = '0'  // 未更新
const THOST_FTDC_UF_Success TThostFtdcUpdateFlagType = '1'   // 更新全部信息成功
const THOST_FTDC_UF_Fail TThostFtdcUpdateFlagType = '2'      // 更新全部信息失败
const THOST_FTDC_UF_TCSuccess TThostFtdcUpdateFlagType = '3' // 更新交易编码成功
const THOST_FTDC_UF_TCFail TThostFtdcUpdateFlagType = '4'    // 更新交易编码失败
const THOST_FTDC_UF_Cancel TThostFtdcUpdateFlagType = '5'    // 已丢弃

// 申请动作类型
type TThostFtdcApplyOperateIDType byte

const THOST_FTDC_AOID_OpenInvestor TThostFtdcApplyOperateIDType = '1'        // 开户
const THOST_FTDC_AOID_ModifyIDCard TThostFtdcApplyOperateIDType = '2'        // 修改身份信息
const THOST_FTDC_AOID_ModifyNoIDCard TThostFtdcApplyOperateIDType = '3'      // 修改一般信息
const THOST_FTDC_AOID_ApplyTradingCode TThostFtdcApplyOperateIDType = '4'    // 申请交易编码
const THOST_FTDC_AOID_CancelTradingCode TThostFtdcApplyOperateIDType = '5'   // 撤销交易编码
const THOST_FTDC_AOID_CancelInvestor TThostFtdcApplyOperateIDType = '6'      // 销户
const THOST_FTDC_AOID_FreezeAccount TThostFtdcApplyOperateIDType = '8'       // 账户休眠
const THOST_FTDC_AOID_ActiveFreezeAccount TThostFtdcApplyOperateIDType = '9' // 激活休眠账户

// 申请状态类型
type TThostFtdcApplyStatusIDType byte

const THOST_FTDC_ASID_NoComplete TThostFtdcApplyStatusIDType = '1' // 未补全
const THOST_FTDC_ASID_Submited TThostFtdcApplyStatusIDType = '2'   // 已提交
const THOST_FTDC_ASID_Checked TThostFtdcApplyStatusIDType = '3'    // 已审核
const THOST_FTDC_ASID_Refused TThostFtdcApplyStatusIDType = '4'    // 已拒绝
const THOST_FTDC_ASID_Deleted TThostFtdcApplyStatusIDType = '5'    // 已删除

// 发送方式类型
type TThostFtdcSendMethodType byte

const THOST_FTDC_UOASM_ByAPI TThostFtdcSendMethodType = '1'  // 文件发送
const THOST_FTDC_UOASM_ByFile TThostFtdcSendMethodType = '2' // 电子发送

// 业务操作类型类型
type TThostFtdcEventTypeType [33]byte

// 操作方法类型
type TThostFtdcEventModeType byte

const THOST_FTDC_EvM_ADD TThostFtdcEventModeType = '1'     // 增加
const THOST_FTDC_EvM_UPDATE TThostFtdcEventModeType = '2'  // 修改
const THOST_FTDC_EvM_DELETE TThostFtdcEventModeType = '3'  // 删除
const THOST_FTDC_EvM_CHECK TThostFtdcEventModeType = '4'   // 复核
const THOST_FTDC_EvM_COPY TThostFtdcEventModeType = '5'    // 复制
const THOST_FTDC_EvM_CANCEL TThostFtdcEventModeType = '6'  // 注销
const THOST_FTDC_EvM_Reverse TThostFtdcEventModeType = '7' // 冲销

// 统一开户申请自动发送类型
type TThostFtdcUOAAutoSendType byte

const THOST_FTDC_UOAA_ASR TThostFtdcUOAAutoSendType = '1'  // 自动发送并接收
const THOST_FTDC_UOAA_ASNR TThostFtdcUOAAutoSendType = '2' // 自动发送，不自动接收
const THOST_FTDC_UOAA_NSAR TThostFtdcUOAAutoSendType = '3' // 不自动发送，自动接收
const THOST_FTDC_UOAA_NSR TThostFtdcUOAAutoSendType = '4'  // 不自动发送，也不自动接收

// 查询深度类型
type TThostFtdcQueryDepthType int32

// 数据中心代码类型
type TThostFtdcDataCenterIDType int32

// 流程ID类型
type TThostFtdcFlowIDType byte

const THOST_FTDC_EvM_InvestorGroupFlow TThostFtdcFlowIDType = '1'     // 投资者对应投资者组设置
const THOST_FTDC_EvM_InvestorRate TThostFtdcFlowIDType = '2'          // 投资者手续费率设置
const THOST_FTDC_EvM_InvestorCommRateModel TThostFtdcFlowIDType = '3' // 投资者手续费率模板关系设置

// 复核级别类型
type TThostFtdcCheckLevelType byte

const THOST_FTDC_CL_Zero TThostFtdcCheckLevelType = '0' // 零级复核
const THOST_FTDC_CL_One TThostFtdcCheckLevelType = '1'  // 一级复核
const THOST_FTDC_CL_Two TThostFtdcCheckLevelType = '2'  // 二级复核

// 操作次数类型
type TThostFtdcCheckNoType int32

// 复核级别类型
type TThostFtdcCheckStatusType byte

const THOST_FTDC_CHS_Init TThostFtdcCheckStatusType = '0'     // 未复核
const THOST_FTDC_CHS_Checking TThostFtdcCheckStatusType = '1' // 复核中
const THOST_FTDC_CHS_Checked TThostFtdcCheckStatusType = '2'  // 已复核
const THOST_FTDC_CHS_Refuse TThostFtdcCheckStatusType = '3'   // 拒绝
const THOST_FTDC_CHS_Cancel TThostFtdcCheckStatusType = '4'   // 作废

// 生效状态类型
type TThostFtdcUsedStatusType byte

const THOST_FTDC_CHU_Unused TThostFtdcUsedStatusType = '0' // 未生效
const THOST_FTDC_CHU_Used TThostFtdcUsedStatusType = '1'   // 已生效
const THOST_FTDC_CHU_Fail TThostFtdcUsedStatusType = '2'   // 生效失败

// 模型名称类型
type TThostFtdcRateTemplateNameType [61]byte

// 用于查询的投资属性字段类型
type TThostFtdcPropertyStringType [2049]byte

// 账户来源类型
type TThostFtdcBankAcountOriginType byte

const THOST_FTDC_BAO_ByAccProperty TThostFtdcBankAcountOriginType = '0' // 手工录入
const THOST_FTDC_BAO_ByFBTransfer TThostFtdcBankAcountOriginType = '1'  // 银期转账

// 结算单月报成交汇总方式类型
type TThostFtdcMonthBillTradeSumType byte

const THOST_FTDC_MBTS_ByInstrument TThostFtdcMonthBillTradeSumType = '0' // 同日同合约
const THOST_FTDC_MBTS_ByDayInsPrc TThostFtdcMonthBillTradeSumType = '1'  // 同日同合约同价格
const THOST_FTDC_MBTS_ByDayIns TThostFtdcMonthBillTradeSumType = '2'     // 同合约

// 银期交易代码枚举类型
type TThostFtdcFBTTradeCodeEnumType string

const THOST_FTDC_FTC_BankLaunchBankToBroker TThostFtdcFBTTradeCodeEnumType = "102001"   // 银行发起银行转期货
const THOST_FTDC_FTC_BrokerLaunchBankToBroker TThostFtdcFBTTradeCodeEnumType = "202001" // 期货发起银行转期货
const THOST_FTDC_FTC_BankLaunchBrokerToBank TThostFtdcFBTTradeCodeEnumType = "102002"   // 银行发起期货转银行
const THOST_FTDC_FTC_BrokerLaunchBrokerToBank TThostFtdcFBTTradeCodeEnumType = "202002" // 期货发起期货转银行

// 模型代码类型
type TThostFtdcRateTemplateIDType [9]byte

// 风险度类型
type TThostFtdcRiskRateType [21]byte

// 时间戳类型
type TThostFtdcTimestampType int32

// 号段规则名称类型
type TThostFtdcInvestorIDRuleNameType [61]byte

// 号段规则表达式类型
type TThostFtdcInvestorIDRuleExprType [513]byte

// 上次OTP漂移值类型
type TThostFtdcLastDriftType int32

// 上次OTP成功值类型
type TThostFtdcLastSuccessType int32

// 令牌密钥类型
type TThostFtdcAuthKeyType [41]byte

// 序列号类型
type TThostFtdcSerialNumberType [17]byte

// 动态令牌类型类型
type TThostFtdcOTPTypeType byte

const THOST_FTDC_OTP_NONE TThostFtdcOTPTypeType = '0' // 无动态令牌
const THOST_FTDC_OTP_TOTP TThostFtdcOTPTypeType = '1' // 时间令牌

// 动态令牌提供商类型
type TThostFtdcOTPVendorsIDType [2]byte

// 动态令牌提供商名称类型
type TThostFtdcOTPVendorsNameType [61]byte

// 动态令牌状态类型
type TThostFtdcOTPStatusType byte

const THOST_FTDC_OTPS_Unused TThostFtdcOTPStatusType = '0' // 未使用
const THOST_FTDC_OTPS_Used TThostFtdcOTPStatusType = '1'   // 已使用
const THOST_FTDC_OTPS_Disuse TThostFtdcOTPStatusType = '2' // 注销

// 经济公司用户类型类型
type TThostFtdcBrokerUserTypeType byte

const THOST_FTDC_BUT_Investor TThostFtdcBrokerUserTypeType = '1'   // 投资者
const THOST_FTDC_BUT_BrokerUser TThostFtdcBrokerUserTypeType = '2' // 操作员

// 期货类型类型
type TThostFtdcFutureTypeType byte

const THOST_FTDC_FUTT_Commodity TThostFtdcFutureTypeType = '1' // 商品期货
const THOST_FTDC_FUTT_Financial TThostFtdcFutureTypeType = '2' // 金融期货

// 资金管理操作类型类型
type TThostFtdcFundEventTypeType byte

const THOST_FTDC_FET_Restriction TThostFtdcFundEventTypeType = '0'         // 转账限额
const THOST_FTDC_FET_TodayRestriction TThostFtdcFundEventTypeType = '1'    // 当日转账限额
const THOST_FTDC_FET_Transfer TThostFtdcFundEventTypeType = '2'            // 期商流水
const THOST_FTDC_FET_Credit TThostFtdcFundEventTypeType = '3'              // 资金冻结
const THOST_FTDC_FET_InvestorWithdrawAlm TThostFtdcFundEventTypeType = '4' // 投资者可提资金比例
const THOST_FTDC_FET_BankRestriction TThostFtdcFundEventTypeType = '5'     // 单个银行帐户转账限额
const THOST_FTDC_FET_Accountregister TThostFtdcFundEventTypeType = '6'     // 银期签约账户
const THOST_FTDC_FET_ExchangeFundIO TThostFtdcFundEventTypeType = '7'      // 交易所出入金
const THOST_FTDC_FET_InvestorFundIO TThostFtdcFundEventTypeType = '8'      // 投资者出入金

// 资金账户来源类型
type TThostFtdcAccountSourceTypeType byte

const THOST_FTDC_AST_FBTransfer TThostFtdcAccountSourceTypeType = '0'  // 银期同步
const THOST_FTDC_AST_ManualEntry TThostFtdcAccountSourceTypeType = '1' // 手工录入

// 交易编码来源类型
type TThostFtdcCodeSourceTypeType byte

const THOST_FTDC_CST_UnifyAccount TThostFtdcCodeSourceTypeType = '0' // 统一开户(已规范)
const THOST_FTDC_CST_ManualEntry TThostFtdcCodeSourceTypeType = '1'  // 手工录入(未规范)

// 操作员范围类型
type TThostFtdcUserRangeType byte

const THOST_FTDC_UR_All TThostFtdcUserRangeType = '0'    // 所有
const THOST_FTDC_UR_Single TThostFtdcUserRangeType = '1' // 单一操作员

// 时间跨度类型
type TThostFtdcTimeSpanType [9]byte

// 动态令牌导入批次编号类型
type TThostFtdcImportSequenceIDType [17]byte

// 交易统计表按客户统计方式类型
type TThostFtdcByGroupType byte

const THOST_FTDC_BG_Investor TThostFtdcByGroupType = '2' // 按投资者统计
const THOST_FTDC_BG_Group TThostFtdcByGroupType = '1'    // 按类统计

// 交易统计表按范围统计方式类型
type TThostFtdcTradeSumStatModeType byte

const THOST_FTDC_TSSM_Instrument TThostFtdcTradeSumStatModeType = '1' // 按合约统计
const THOST_FTDC_TSSM_Product TThostFtdcTradeSumStatModeType = '2'    // 按产品统计
const THOST_FTDC_TSSM_Exchange TThostFtdcTradeSumStatModeType = '3'   // 按交易所统计

// 组合成交类型类型
type TThostFtdcComTypeType int32

// 产品标识类型
type TThostFtdcUserProductIDType [33]byte

// 产品名称类型
type TThostFtdcUserProductNameType [65]byte

// 产品说明类型
type TThostFtdcUserProductMemoType [129]byte

// 新增或变更标志类型
type TThostFtdcCSRCCancelFlagType [2]byte

// 日期类型
type TThostFtdcCSRCDateType [11]byte

// 客户名称类型
type TThostFtdcCSRCInvestorNameType [201]byte

// 客户名称类型
type TThostFtdcCSRCOpenInvestorNameType [101]byte

// 客户代码类型
type TThostFtdcCSRCInvestorIDType [13]byte

// 证件号码类型
type TThostFtdcCSRCIdentifiedCardNoType [51]byte

// 交易编码类型
type TThostFtdcCSRCClientIDType [11]byte

// 银行标识类型
type TThostFtdcCSRCBankFlagType [3]byte

// 银行账户类型
type TThostFtdcCSRCBankAccountType [23]byte

// 开户人类型
type TThostFtdcCSRCOpenNameType [401]byte

// 说明类型
type TThostFtdcCSRCMemoType [101]byte

// 时间类型
type TThostFtdcCSRCTimeType [11]byte

// 成交流水号类型
type TThostFtdcCSRCTradeIDType [21]byte

// 合约代码类型
type TThostFtdcCSRCExchangeInstIDType [31]byte

// 质押品名称类型
type TThostFtdcCSRCMortgageNameType [7]byte

// 事由类型
type TThostFtdcCSRCReasonType [3]byte

// 是否为非结算会员类型
type TThostFtdcIsSettlementType [2]byte

// 资金类型
type TThostFtdcCSRCMoneyType float64

// 价格类型
type TThostFtdcCSRCPriceType float64

// 期权类型类型
type TThostFtdcCSRCOptionsTypeType [2]byte

// 执行价类型
type TThostFtdcCSRCStrikePriceType float64

// 标的品种类型
type TThostFtdcCSRCTargetProductIDType [3]byte

// 标的合约类型
type TThostFtdcCSRCTargetInstrIDType [31]byte

// 手续费率模板名称类型
type TThostFtdcCommModelNameType [161]byte

// 手续费率模板备注类型
type TThostFtdcCommModelMemoType [1025]byte

// 日期表达式设置类型类型
type TThostFtdcExprSetModeType byte

const THOST_FTDC_ESM_Relative TThostFtdcExprSetModeType = '1' // 相对已有规则设置
const THOST_FTDC_ESM_Typical TThostFtdcExprSetModeType = '2'  // 典型设置

// 投资者范围类型
type TThostFtdcRateInvestorRangeType byte

const THOST_FTDC_RIR_All TThostFtdcRateInvestorRangeType = '1'    // 公司标准
const THOST_FTDC_RIR_Model TThostFtdcRateInvestorRangeType = '2'  // 模板
const THOST_FTDC_RIR_Single TThostFtdcRateInvestorRangeType = '3' // 单一投资者

// 代理经纪公司代码类型
type TThostFtdcAgentBrokerIDType [13]byte

// 交易中心代码类型
type TThostFtdcDRIdentityIDType int32

// 交易中心名称类型
type TThostFtdcDRIdentityNameType [65]byte

// DBLink标识号类型
type TThostFtdcDBLinkIDType [31]byte

// 主次用系统数据同步状态类型
type TThostFtdcSyncDataStatusType byte

const THOST_FTDC_SDS_Initialize TThostFtdcSyncDataStatusType = '0'    // 未同步
const THOST_FTDC_SDS_Settlementing TThostFtdcSyncDataStatusType = '1' // 同步中
const THOST_FTDC_SDS_Settlemented TThostFtdcSyncDataStatusType = '2'  // 已同步

// 成交来源类型
type TThostFtdcTradeSourceType byte

const THOST_FTDC_TSRC_NORMAL TThostFtdcTradeSourceType = '0' // 来自交易所普通回报
const THOST_FTDC_TSRC_QUERY TThostFtdcTradeSourceType = '1'  // 来自查询

// 产品合约统计方式类型
type TThostFtdcFlexStatModeType byte

const THOST_FTDC_FSM_Product TThostFtdcFlexStatModeType = '1'  // 产品统计
const THOST_FTDC_FSM_Exchange TThostFtdcFlexStatModeType = '2' // 交易所统计
const THOST_FTDC_FSM_All TThostFtdcFlexStatModeType = '3'      // 统计所有

// 投资者范围统计方式类型
type TThostFtdcByInvestorRangeType byte

const THOST_FTDC_BIR_Property TThostFtdcByInvestorRangeType = '1' // 属性统计
const THOST_FTDC_BIR_All TThostFtdcByInvestorRangeType = '2'      // 统计所有

// 风险度类型
type TThostFtdcSRiskRateType [21]byte

// 序号类型
type TThostFtdcSequenceNo12Type int32

// 投资者范围类型
type TThostFtdcPropertyInvestorRangeType byte

const THOST_FTDC_PIR_All TThostFtdcPropertyInvestorRangeType = '1'      // 所有
const THOST_FTDC_PIR_Property TThostFtdcPropertyInvestorRangeType = '2' // 投资者属性
const THOST_FTDC_PIR_Single TThostFtdcPropertyInvestorRangeType = '3'   // 单一投资者

// 文件状态类型
type TThostFtdcFileStatusType byte

const THOST_FTDC_FIS_NoCreate TThostFtdcFileStatusType = '0' // 未生成
const THOST_FTDC_FIS_Created TThostFtdcFileStatusType = '1'  // 已生成
const THOST_FTDC_FIS_Failed TThostFtdcFileStatusType = '2'   // 生成失败

// 文件生成方式类型
type TThostFtdcFileGenStyleType byte

const THOST_FTDC_FGS_FileTransmit TThostFtdcFileGenStyleType = '0' // 下发
const THOST_FTDC_FGS_FileGen TThostFtdcFileGenStyleType = '1'      // 生成

// 系统日志操作方法类型
type TThostFtdcSysOperModeType byte

const THOST_FTDC_SoM_Add TThostFtdcSysOperModeType = '1'    // 增加
const THOST_FTDC_SoM_Update TThostFtdcSysOperModeType = '2' // 修改
const THOST_FTDC_SoM_Delete TThostFtdcSysOperModeType = '3' // 删除
const THOST_FTDC_SoM_Copy TThostFtdcSysOperModeType = '4'   // 复制
const THOST_FTDC_SoM_AcTive TThostFtdcSysOperModeType = '5' // 激活
const THOST_FTDC_SoM_CanCel TThostFtdcSysOperModeType = '6' // 注销
const THOST_FTDC_SoM_ReSet TThostFtdcSysOperModeType = '7'  // 重置

// 系统日志操作类型类型
type TThostFtdcSysOperTypeType byte

const THOST_FTDC_SoT_UpdatePassword TThostFtdcSysOperTypeType = '0'          // 修改操作员密码
const THOST_FTDC_SoT_UserDepartment TThostFtdcSysOperTypeType = '1'          // 操作员组织架构关系
const THOST_FTDC_SoT_RoleManager TThostFtdcSysOperTypeType = '2'             // 角色管理
const THOST_FTDC_SoT_RoleFunction TThostFtdcSysOperTypeType = '3'            // 角色功能设置
const THOST_FTDC_SoT_BaseParam TThostFtdcSysOperTypeType = '4'               // 基础参数设置
const THOST_FTDC_SoT_SetUserID TThostFtdcSysOperTypeType = '5'               // 设置操作员
const THOST_FTDC_SoT_SetUserRole TThostFtdcSysOperTypeType = '6'             // 用户角色设置
const THOST_FTDC_SoT_UserIpRestriction TThostFtdcSysOperTypeType = '7'       // 用户IP限制
const THOST_FTDC_SoT_DepartmentManager TThostFtdcSysOperTypeType = '8'       // 组织架构管理
const THOST_FTDC_SoT_DepartmentCopy TThostFtdcSysOperTypeType = '9'          // 组织架构向查询分类复制
const THOST_FTDC_SoT_Tradingcode TThostFtdcSysOperTypeType = 'A'             // 交易编码管理
const THOST_FTDC_SoT_InvestorStatus TThostFtdcSysOperTypeType = 'B'          // 投资者状态维护
const THOST_FTDC_SoT_InvestorAuthority TThostFtdcSysOperTypeType = 'C'       // 投资者权限管理
const THOST_FTDC_SoT_PropertySet TThostFtdcSysOperTypeType = 'D'             // 属性设置
const THOST_FTDC_SoT_ReSetInvestorPasswd TThostFtdcSysOperTypeType = 'E'     // 重置投资者密码
const THOST_FTDC_SoT_InvestorPersonalityInfo TThostFtdcSysOperTypeType = 'F' // 投资者个性信息维护

// 上报数据查询类型类型
type TThostFtdcCSRCDataQueyTypeType byte

const THOST_FTDC_CSRCQ_Current TThostFtdcCSRCDataQueyTypeType = '0' // 查询当前交易日报送的数据
const THOST_FTDC_CSRCQ_History TThostFtdcCSRCDataQueyTypeType = '1' // 查询历史报送的代理经纪公司的数据

// 休眠状态类型
type TThostFtdcFreezeStatusType byte

const THOST_FTDC_FRS_Normal TThostFtdcFreezeStatusType = '1' // 活跃
const THOST_FTDC_FRS_Freeze TThostFtdcFreezeStatusType = '0' // 休眠

// 规范状态类型
type TThostFtdcStandardStatusType byte

const THOST_FTDC_STST_Standard TThostFtdcStandardStatusType = '0'    // 已规范
const THOST_FTDC_STST_NonStandard TThostFtdcStandardStatusType = '1' // 未规范

// 休眠状态类型
type TThostFtdcCSRCFreezeStatusType [2]byte

// 配置类型类型
type TThostFtdcRightParamTypeType byte

const THOST_FTDC_RPT_Freeze TThostFtdcRightParamTypeType = '1'           // 休眠户
const THOST_FTDC_RPT_FreezeActive TThostFtdcRightParamTypeType = '2'     // 激活休眠户
const THOST_FTDC_RPT_OpenLimit TThostFtdcRightParamTypeType = '3'        // 开仓权限限制
const THOST_FTDC_RPT_RelieveOpenLimit TThostFtdcRightParamTypeType = '4' // 解除开仓权限限制

// 模板代码类型
type TThostFtdcRightTemplateIDType [9]byte

// 模板名称类型
type TThostFtdcRightTemplateNameType [61]byte

// 反洗钱审核表数据状态类型
type TThostFtdcDataStatusType byte

const THOST_FTDC_AMLDS_Normal TThostFtdcDataStatusType = '0'  // 正常
const THOST_FTDC_AMLDS_Deleted TThostFtdcDataStatusType = '1' // 已删除

// 审核状态类型
type TThostFtdcAMLCheckStatusType byte

const THOST_FTDC_AMLCHS_Init TThostFtdcAMLCheckStatusType = '0'         // 未复核
const THOST_FTDC_AMLCHS_Checking TThostFtdcAMLCheckStatusType = '1'     // 复核中
const THOST_FTDC_AMLCHS_Checked TThostFtdcAMLCheckStatusType = '2'      // 已复核
const THOST_FTDC_AMLCHS_RefuseReport TThostFtdcAMLCheckStatusType = '3' // 拒绝上报

// 日期类型类型
type TThostFtdcAmlDateTypeType byte

const THOST_FTDC_AMLDT_DrawDay TThostFtdcAmlDateTypeType = '0'  // 检查日期
const THOST_FTDC_AMLDT_TouchDay TThostFtdcAmlDateTypeType = '1' // 发生日期

// 审核级别类型
type TThostFtdcAmlCheckLevelType byte

const THOST_FTDC_AMLCL_CheckLevel0 TThostFtdcAmlCheckLevelType = '0' // 零级审核
const THOST_FTDC_AMLCL_CheckLevel1 TThostFtdcAmlCheckLevelType = '1' // 一级审核
const THOST_FTDC_AMLCL_CheckLevel2 TThostFtdcAmlCheckLevelType = '2' // 二级审核
const THOST_FTDC_AMLCL_CheckLevel3 TThostFtdcAmlCheckLevelType = '3' // 三级审核

// 反洗钱数据抽取审核流程类型
type TThostFtdcAmlCheckFlowType [2]byte

// 数据类型类型
type TThostFtdcDataTypeType [129]byte

// 导出文件类型类型
type TThostFtdcExportFileTypeType byte

const THOST_FTDC_EFT_CSV TThostFtdcExportFileTypeType = '0'   // CSV
const THOST_FTDC_EFT_EXCEL TThostFtdcExportFileTypeType = '1' // Excel
const THOST_FTDC_EFT_DBF TThostFtdcExportFileTypeType = '2'   // DBF

// 结算配置类型类型
type TThostFtdcSettleManagerTypeType byte

const THOST_FTDC_SMT_Before TThostFtdcSettleManagerTypeType = '1'       // 结算前准备
const THOST_FTDC_SMT_Settlement TThostFtdcSettleManagerTypeType = '2'   // 结算
const THOST_FTDC_SMT_After TThostFtdcSettleManagerTypeType = '3'        // 结算后核对
const THOST_FTDC_SMT_Settlemented TThostFtdcSettleManagerTypeType = '4' // 结算后处理

// 结算配置代码类型
type TThostFtdcSettleManagerIDType [33]byte

// 结算配置名称类型
type TThostFtdcSettleManagerNameType [129]byte

// 结算配置等级类型
type TThostFtdcSettleManagerLevelType byte

const THOST_FTDC_SML_Must TThostFtdcSettleManagerLevelType = '1'   // 必要
const THOST_FTDC_SML_Alarm TThostFtdcSettleManagerLevelType = '2'  // 警告
const THOST_FTDC_SML_Prompt TThostFtdcSettleManagerLevelType = '3' // 提示
const THOST_FTDC_SML_Ignore TThostFtdcSettleManagerLevelType = '4' // 不检查

// 模块分组类型
type TThostFtdcSettleManagerGroupType byte

const THOST_FTDC_SMG_Exhcange TThostFtdcSettleManagerGroupType = '1' // 交易所核对
const THOST_FTDC_SMG_ASP TThostFtdcSettleManagerGroupType = '2'      // 内部核对
const THOST_FTDC_SMG_CSRC TThostFtdcSettleManagerGroupType = '3'     // 上报数据核对

// 核对结果说明类型
type TThostFtdcCheckResultMemoType [1025]byte

// 功能链接类型
type TThostFtdcFunctionUrlType [1025]byte

// 客户端认证信息类型
type TThostFtdcAuthInfoType [129]byte

// 客户端认证码类型
type TThostFtdcAuthCodeType [17]byte

// 保值额度使用类型类型
type TThostFtdcLimitUseTypeType byte

const THOST_FTDC_LUT_Repeatable TThostFtdcLimitUseTypeType = '1'   // 可重复使用
const THOST_FTDC_LUT_Unrepeatable TThostFtdcLimitUseTypeType = '2' // 不可重复使用

// 数据来源类型
type TThostFtdcDataResourceType byte

const THOST_FTDC_DAR_Settle TThostFtdcDataResourceType = '1'   // 本系统
const THOST_FTDC_DAR_Exchange TThostFtdcDataResourceType = '2' // 交易所
const THOST_FTDC_DAR_CSRC TThostFtdcDataResourceType = '3'     // 报送数据

// 保证金类型类型
type TThostFtdcMarginTypeType byte

const THOST_FTDC_MGT_ExchMarginRate TThostFtdcMarginTypeType = '0'       // 交易所保证金率
const THOST_FTDC_MGT_InstrMarginRate TThostFtdcMarginTypeType = '1'      // 投资者保证金率
const THOST_FTDC_MGT_InstrMarginRateTrade TThostFtdcMarginTypeType = '2' // 投资者交易保证金率

// 生效类型类型
type TThostFtdcActiveTypeType byte

const THOST_FTDC_ACT_Intraday TThostFtdcActiveTypeType = '1' // 仅当日生效
const THOST_FTDC_ACT_Long TThostFtdcActiveTypeType = '2'     // 长期生效

// 冲突保证金率类型类型
type TThostFtdcMarginRateTypeType byte

const THOST_FTDC_MRT_Exchange TThostFtdcMarginRateTypeType = '1'      // 交易所保证金率
const THOST_FTDC_MRT_Investor TThostFtdcMarginRateTypeType = '2'      // 投资者保证金率
const THOST_FTDC_MRT_InvestorTrade TThostFtdcMarginRateTypeType = '3' // 投资者交易保证金率

// 备份数据状态类型
type TThostFtdcBackUpStatusType byte

const THOST_FTDC_BUS_UnBak TThostFtdcBackUpStatusType = '0'   // 未生成备份数据
const THOST_FTDC_BUS_BakUp TThostFtdcBackUpStatusType = '1'   // 备份数据生成中
const THOST_FTDC_BUS_BakUped TThostFtdcBackUpStatusType = '2' // 已生成备份数据
const THOST_FTDC_BUS_BakFail TThostFtdcBackUpStatusType = '3' // 备份数据失败

// 结算初始化状态类型
type TThostFtdcInitSettlementType byte

const THOST_FTDC_SIS_UnInitialize TThostFtdcInitSettlementType = '0' // 结算初始化未开始
const THOST_FTDC_SIS_Initialize TThostFtdcInitSettlementType = '1'   // 结算初始化中
const THOST_FTDC_SIS_Initialized TThostFtdcInitSettlementType = '2'  // 结算初始化完成

// 报表数据生成状态类型
type TThostFtdcReportStatusType byte

const THOST_FTDC_SRS_NoCreate TThostFtdcReportStatusType = '0'   // 未生成报表数据
const THOST_FTDC_SRS_Create TThostFtdcReportStatusType = '1'     // 报表数据生成中
const THOST_FTDC_SRS_Created TThostFtdcReportStatusType = '2'    // 已生成报表数据
const THOST_FTDC_SRS_CreateFail TThostFtdcReportStatusType = '3' // 生成报表数据失败

// 数据归档状态类型
type TThostFtdcSaveStatusType byte

const THOST_FTDC_SSS_UnSaveData TThostFtdcSaveStatusType = '0' // 归档未完成
const THOST_FTDC_SSS_SaveDatad TThostFtdcSaveStatusType = '1'  // 归档完成

// 结算确认数据归档状态类型
type TThostFtdcSettArchiveStatusType byte

const THOST_FTDC_SAS_UnArchived TThostFtdcSettArchiveStatusType = '0'  // 未归档数据
const THOST_FTDC_SAS_Archiving TThostFtdcSettArchiveStatusType = '1'   // 数据归档中
const THOST_FTDC_SAS_Archived TThostFtdcSettArchiveStatusType = '2'    // 已归档数据
const THOST_FTDC_SAS_ArchiveFail TThostFtdcSettArchiveStatusType = '3' // 归档数据失败

// CTP交易系统类型类型
type TThostFtdcCTPTypeType byte

const THOST_FTDC_CTPT_Unkown TThostFtdcCTPTypeType = '0'     // 未知类型
const THOST_FTDC_CTPT_MainCenter TThostFtdcCTPTypeType = '1' // 主中心
const THOST_FTDC_CTPT_BackUp TThostFtdcCTPTypeType = '2'     // 备中心

// 工具代码类型
type TThostFtdcToolIDType [9]byte

// 工具名称类型
type TThostFtdcToolNameType [81]byte

// 平仓处理类型类型
type TThostFtdcCloseDealTypeType byte

const THOST_FTDC_CDT_Normal TThostFtdcCloseDealTypeType = '0'    // 正常
const THOST_FTDC_CDT_SpecFirst TThostFtdcCloseDealTypeType = '1' // 投机平仓优先

// 货币质押资金可用范围类型
type TThostFtdcMortgageFundUseRangeType byte

const THOST_FTDC_MFUR_None TThostFtdcMortgageFundUseRangeType = '0'   // 不能使用
const THOST_FTDC_MFUR_Margin TThostFtdcMortgageFundUseRangeType = '1' // 用于保证金
const THOST_FTDC_MFUR_All TThostFtdcMortgageFundUseRangeType = '2'    // 用于手续费、盈亏、保证金
const THOST_FTDC_MFUR_CNY3 TThostFtdcMortgageFundUseRangeType = '3'   // 人民币方案3

// 币种单位数量类型
type TThostFtdcCurrencyUnitType float64

// 汇率类型
type TThostFtdcExchangeRateType float64

// 特殊产品类型类型
type TThostFtdcSpecProductTypeType byte

const THOST_FTDC_SPT_CzceHedge TThostFtdcSpecProductTypeType = '1'          // 郑商所套保产品
const THOST_FTDC_SPT_IneForeignCurrency TThostFtdcSpecProductTypeType = '2' // 货币质押产品
const THOST_FTDC_SPT_DceOpenClose TThostFtdcSpecProductTypeType = '3'       // 大连短线开平仓产品

// 货币质押类型类型
type TThostFtdcFundMortgageTypeType byte

const THOST_FTDC_FMT_Mortgage TThostFtdcFundMortgageTypeType = '1'   // 质押
const THOST_FTDC_FMT_Redemption TThostFtdcFundMortgageTypeType = '2' // 解质

// 投资者账户结算参数代码类型
type TThostFtdcAccountSettlementParamIDType byte

const THOST_FTDC_ASPI_BaseMargin TThostFtdcAccountSettlementParamIDType = '1'     // 基础保证金
const THOST_FTDC_ASPI_LowestInterest TThostFtdcAccountSettlementParamIDType = '2' // 最低权益标准

// 币种名称类型
type TThostFtdcCurrencyNameType [31]byte

// 币种符号类型
type TThostFtdcCurrencySignType [4]byte

// 货币质押方向类型
type TThostFtdcFundMortDirectionType byte

const THOST_FTDC_FMD_In TThostFtdcFundMortDirectionType = '1'  // 货币质入
const THOST_FTDC_FMD_Out TThostFtdcFundMortDirectionType = '2' // 货币质出

// 换汇类别类型
type TThostFtdcBusinessClassType byte

const THOST_FTDC_BT_Profit TThostFtdcBusinessClassType = '0' // 盈利
const THOST_FTDC_BT_Loss TThostFtdcBusinessClassType = '1'   // 亏损
const THOST_FTDC_BT_Other TThostFtdcBusinessClassType = 'Z'  // 其他

// 换汇数据来源类型
type TThostFtdcSwapSourceTypeType byte

const THOST_FTDC_SST_Manual TThostFtdcSwapSourceTypeType = '0'    // 手工
const THOST_FTDC_SST_Automatic TThostFtdcSwapSourceTypeType = '1' // 自动生成

// 换汇类型类型
type TThostFtdcCurrExDirectionType byte

const THOST_FTDC_CED_Settlement TThostFtdcCurrExDirectionType = '0' // 结汇
const THOST_FTDC_CED_Sale TThostFtdcCurrExDirectionType = '1'       // 售汇

// 申请状态类型
type TThostFtdcCurrencySwapStatusType byte

const THOST_FTDC_CSS_Entry TThostFtdcCurrencySwapStatusType = '1'   // 已录入
const THOST_FTDC_CSS_Approve TThostFtdcCurrencySwapStatusType = '2' // 已审核
const THOST_FTDC_CSS_Refuse TThostFtdcCurrencySwapStatusType = '3'  // 已拒绝
const THOST_FTDC_CSS_Revoke TThostFtdcCurrencySwapStatusType = '4'  // 已撤销
const THOST_FTDC_CSS_Send TThostFtdcCurrencySwapStatusType = '5'    // 已发送
const THOST_FTDC_CSS_Success TThostFtdcCurrencySwapStatusType = '6' // 换汇成功
const THOST_FTDC_CSS_Failure TThostFtdcCurrencySwapStatusType = '7' // 换汇失败

// 凭证号类型
type TThostFtdcCurrExchCertNoType [13]byte

// 批次号类型
type TThostFtdcBatchSerialNoType [21]byte

// 换汇发送标志类型
type TThostFtdcReqFlagType byte

const THOST_FTDC_REQF_NoSend TThostFtdcReqFlagType = '0'      // 未发送
const THOST_FTDC_REQF_SendSuccess TThostFtdcReqFlagType = '1' // 发送成功
const THOST_FTDC_REQF_SendFailed TThostFtdcReqFlagType = '2'  // 发送失败
const THOST_FTDC_REQF_WaitReSend TThostFtdcReqFlagType = '3'  // 等待重发

// 换汇返回成功标志类型
type TThostFtdcResFlagType byte

const THOST_FTDC_RESF_Success TThostFtdcResFlagType = '0'      // 成功
const THOST_FTDC_RESF_InsuffiCient TThostFtdcResFlagType = '1' // 账户余额不足
const THOST_FTDC_RESF_UnKnown TThostFtdcResFlagType = '8'      // 交易结果未知

// 换汇页面控制类型
type TThostFtdcPageControlType [2]byte

// 记录数类型
type TThostFtdcRecordCountType int32

// 换汇需确认信息类型
type TThostFtdcCurrencySwapMemoType [101]byte

// 修改状态类型
type TThostFtdcExStatusType byte

const THOST_FTDC_EXS_Before TThostFtdcExStatusType = '0' // 修改前
const THOST_FTDC_EXS_After TThostFtdcExStatusType = '1'  // 修改后

// 开户客户地域类型
type TThostFtdcClientRegionType byte

const THOST_FTDC_CR_Domestic TThostFtdcClientRegionType = '1' // 国内客户
const THOST_FTDC_CR_GMT TThostFtdcClientRegionType = '2'      // 港澳台客户
const THOST_FTDC_CR_Foreign TThostFtdcClientRegionType = '3'  // 国外客户

// 工作单位类型
type TThostFtdcWorkPlaceType [101]byte

// 经营期限类型
type TThostFtdcBusinessPeriodType [21]byte

// 网址类型
type TThostFtdcWebSiteType [101]byte

// 统一开户证件类型类型
type TThostFtdcUOAIdCardTypeType [3]byte

// 开户模式类型
type TThostFtdcClientModeType [3]byte

// 投资者全称类型
type TThostFtdcInvestorFullNameType [101]byte

// 境外中介机构ID类型
type TThostFtdcUOABrokerIDType [11]byte

// 邮政编码类型
type TThostFtdcUOAZipCodeType [11]byte

// 电子邮箱类型
type TThostFtdcUOAEMailType [101]byte

// 城市类型
type TThostFtdcOldCityType [41]byte

// 法人代表证件号码类型
type TThostFtdcCorporateIdentifiedCardNoType [101]byte

// 是否有董事会类型
type TThostFtdcHasBoardType byte

const THOST_FTDC_HB_No TThostFtdcHasBoardType = '0'  // 没有
const THOST_FTDC_HB_Yes TThostFtdcHasBoardType = '1' // 有

// 启动模式类型
type TThostFtdcStartModeType byte

const THOST_FTDC_SM_Normal TThostFtdcStartModeType = '1'  // 正常
const THOST_FTDC_SM_Emerge TThostFtdcStartModeType = '2'  // 应急
const THOST_FTDC_SM_Restore TThostFtdcStartModeType = '3' // 恢复

// 模型类型类型
type TThostFtdcTemplateTypeType byte

const THOST_FTDC_TPT_Full TThostFtdcTemplateTypeType = '1'      // 全量
const THOST_FTDC_TPT_Increment TThostFtdcTemplateTypeType = '2' // 增量
const THOST_FTDC_TPT_BackUp TThostFtdcTemplateTypeType = '3'    // 备份

// 登录模式类型
type TThostFtdcLoginModeType byte

const THOST_FTDC_LM_Trade TThostFtdcLoginModeType = '0'    // 交易
const THOST_FTDC_LM_Transfer TThostFtdcLoginModeType = '1' // 转账

// 日历提示类型类型
type TThostFtdcPromptTypeType byte

const THOST_FTDC_CPT_Instrument TThostFtdcPromptTypeType = '1' // 合约上下市
const THOST_FTDC_CPT_Margin TThostFtdcPromptTypeType = '2'     // 保证金分段生效

// 分户管理资产编码类型
type TThostFtdcLedgerManageIDType [51]byte

// 投资品种类型
type TThostFtdcInvestVarietyType [101]byte

// 账户类别类型
type TThostFtdcBankAccountTypeType [2]byte

// 开户银行类型
type TThostFtdcLedgerManageBankType [101]byte

// 开户营业部类型
type TThostFtdcCffexDepartmentNameType [101]byte

// 营业部代码类型
type TThostFtdcCffexDepartmentCodeType [9]byte

// 是否有托管人类型
type TThostFtdcHasTrusteeType byte

const THOST_FTDC_HT_Yes TThostFtdcHasTrusteeType = '1' // 有
const THOST_FTDC_HT_No TThostFtdcHasTrusteeType = '0'  // 没有

// 说明类型
type TThostFtdcCSRCMemo1Type [41]byte

// 代理资产管理业务的期货公司全称类型
type TThostFtdcAssetmgrCFullNameType [101]byte

// 资产管理业务批文号类型
type TThostFtdcAssetmgrApprovalNOType [51]byte

// 资产管理业务负责人姓名类型
type TThostFtdcAssetmgrMgrNameType [401]byte

// 机构类型类型
type TThostFtdcAmTypeType byte

const THOST_FTDC_AMT_Bank TThostFtdcAmTypeType = '1'       // 银行
const THOST_FTDC_AMT_Securities TThostFtdcAmTypeType = '2' // 证券公司
const THOST_FTDC_AMT_Fund TThostFtdcAmTypeType = '3'       // 基金公司
const THOST_FTDC_AMT_Insurance TThostFtdcAmTypeType = '4'  // 保险公司
const THOST_FTDC_AMT_Trust TThostFtdcAmTypeType = '5'      // 信托公司
const THOST_FTDC_AMT_Other TThostFtdcAmTypeType = '9'      // 其他

// 机构类型类型
type TThostFtdcCSRCAmTypeType [5]byte

// 出入金类型类型
type TThostFtdcCSRCFundIOTypeType byte

const THOST_FTDC_CFIOT_FundIO TThostFtdcCSRCFundIOTypeType = '0'       // 出入金
const THOST_FTDC_CFIOT_SwapCurrency TThostFtdcCSRCFundIOTypeType = '1' // 银期换汇

// 结算账户类型类型
type TThostFtdcCusAccountTypeType byte

const THOST_FTDC_CAT_Futures TThostFtdcCusAccountTypeType = '1'          // 期货结算账户
const THOST_FTDC_CAT_AssetmgrFuture TThostFtdcCusAccountTypeType = '2'   // 纯期货资管业务下的资管结算账户
const THOST_FTDC_CAT_AssetmgrTrustee TThostFtdcCusAccountTypeType = '3'  // 综合类资管业务下的期货资管托管账户
const THOST_FTDC_CAT_AssetmgrTransfer TThostFtdcCusAccountTypeType = '4' // 综合类资管业务下的资金中转账户

// 国籍类型
type TThostFtdcCSRCNationalType [4]byte

// 二级代理ID类型
type TThostFtdcCSRCSecAgentIDType [11]byte

// 通知语言类型类型
type TThostFtdcLanguageTypeType byte

const THOST_FTDC_LT_Chinese TThostFtdcLanguageTypeType = '1' // 中文
const THOST_FTDC_LT_English TThostFtdcLanguageTypeType = '2' // 英文

// 投资账户类型
type TThostFtdcAmAccountType [23]byte

// 资产管理客户类型类型
type TThostFtdcAssetmgrClientTypeType byte

const THOST_FTDC_AMCT_Person TThostFtdcAssetmgrClientTypeType = '1'       // 个人资管客户
const THOST_FTDC_AMCT_Organ TThostFtdcAssetmgrClientTypeType = '2'        // 单位资管客户
const THOST_FTDC_AMCT_SpecialOrgan TThostFtdcAssetmgrClientTypeType = '4' // 特殊单位资管客户

// 投资类型类型
type TThostFtdcAssetmgrTypeType byte

const THOST_FTDC_ASST_Futures TThostFtdcAssetmgrTypeType = '3'      // 期货类
const THOST_FTDC_ASST_SpecialOrgan TThostFtdcAssetmgrTypeType = '4' // 综合类

// 计量单位类型
type TThostFtdcUOMType [11]byte

// 上期所合约生命周期状态类型
type TThostFtdcSHFEInstLifePhaseType [3]byte

// 产品类型类型
type TThostFtdcSHFEProductClassType [11]byte

// 价格小数位类型
type TThostFtdcPriceDecimalType [2]byte

// 平值期权标志类型
type TThostFtdcInTheMoneyFlagType [2]byte

// 合约比较类型类型
type TThostFtdcCheckInstrTypeType byte

const THOST_FTDC_CIT_HasExch TThostFtdcCheckInstrTypeType = '0' // 合约交易所不存在
const THOST_FTDC_CIT_HasATP TThostFtdcCheckInstrTypeType = '1'  // 合约本系统不存在
const THOST_FTDC_CIT_HasDiff TThostFtdcCheckInstrTypeType = '2' // 合约比较不一致

// 交割类型类型
type TThostFtdcDeliveryTypeType byte

const THOST_FTDC_DT_HandDeliv TThostFtdcDeliveryTypeType = '1'   // 手工交割
const THOST_FTDC_DT_PersonDeliv TThostFtdcDeliveryTypeType = '2' // 到期交割

// 资金类型
type TThostFtdcBigMoneyType float64

// 大额单边保证金算法类型
type TThostFtdcMaxMarginSideAlgorithmType byte

const THOST_FTDC_MMSA_NO TThostFtdcMaxMarginSideAlgorithmType = '0'  // 不使用大额单边保证金算法
const THOST_FTDC_MMSA_YES TThostFtdcMaxMarginSideAlgorithmType = '1' // 使用大额单边保证金算法

// 资产管理客户类型类型
type TThostFtdcDAClientTypeType byte

const THOST_FTDC_CACT_Person TThostFtdcDAClientTypeType = '0'  // 自然人
const THOST_FTDC_CACT_Company TThostFtdcDAClientTypeType = '1' // 法人
const THOST_FTDC_CACT_Other TThostFtdcDAClientTypeType = '2'   // 其他

// 套利合约代码类型
type TThostFtdcCombinInstrIDType [61]byte

// 各腿结算价类型
type TThostFtdcCombinSettlePriceType [61]byte

// 优先级类型
type TThostFtdcDCEPriorityType int32

// 成交组号类型
type TThostFtdcTradeGroupIDType int32

// 是否校验开户可用资金类型
type TThostFtdcIsCheckPrepaType int32

// 投资类型类型
type TThostFtdcUOAAssetmgrTypeType byte

const THOST_FTDC_UOAAT_Futures TThostFtdcUOAAssetmgrTypeType = '1'      // 期货类
const THOST_FTDC_UOAAT_SpecialOrgan TThostFtdcUOAAssetmgrTypeType = '2' // 综合类

// 买卖方向类型
type TThostFtdcDirectionEnType byte

const THOST_FTDC_DEN_Buy TThostFtdcDirectionEnType = '0'  // Buy
const THOST_FTDC_DEN_Sell TThostFtdcDirectionEnType = '1' // Sell

// 开平标志类型
type TThostFtdcOffsetFlagEnType byte

const THOST_FTDC_OFEN_Open TThostFtdcOffsetFlagEnType = '0'            // Position Opening
const THOST_FTDC_OFEN_Close TThostFtdcOffsetFlagEnType = '1'           // Position Close
const THOST_FTDC_OFEN_ForceClose TThostFtdcOffsetFlagEnType = '2'      // Forced Liquidation
const THOST_FTDC_OFEN_CloseToday TThostFtdcOffsetFlagEnType = '3'      // Close Today
const THOST_FTDC_OFEN_CloseYesterday TThostFtdcOffsetFlagEnType = '4'  // Close Prev.
const THOST_FTDC_OFEN_ForceOff TThostFtdcOffsetFlagEnType = '5'        // Forced Reduction
const THOST_FTDC_OFEN_LocalForceClose TThostFtdcOffsetFlagEnType = '6' // Local Forced Liquidation

// 投机套保标志类型
type TThostFtdcHedgeFlagEnType byte

const THOST_FTDC_HFEN_Speculation TThostFtdcHedgeFlagEnType = '1' // Speculation
const THOST_FTDC_HFEN_Arbitrage TThostFtdcHedgeFlagEnType = '2'   // Arbitrage
const THOST_FTDC_HFEN_Hedge TThostFtdcHedgeFlagEnType = '3'       // Hedge

// 出入金类型类型
type TThostFtdcFundIOTypeEnType byte

const THOST_FTDC_FIOTEN_FundIO TThostFtdcFundIOTypeEnType = '1'       // Deposit/Withdrawal
const THOST_FTDC_FIOTEN_Transfer TThostFtdcFundIOTypeEnType = '2'     // Bank-Futures Transfer
const THOST_FTDC_FIOTEN_SwapCurrency TThostFtdcFundIOTypeEnType = '3' // Bank-Futures FX Exchange

// 资金类型类型
type TThostFtdcFundTypeEnType byte

const THOST_FTDC_FTEN_Deposite TThostFtdcFundTypeEnType = '1'      // Bank Deposit
const THOST_FTDC_FTEN_ItemFund TThostFtdcFundTypeEnType = '2'      // Payment/Fee
const THOST_FTDC_FTEN_Company TThostFtdcFundTypeEnType = '3'       // Brokerage Adj
const THOST_FTDC_FTEN_InnerTransfer TThostFtdcFundTypeEnType = '4' // Internal Transfer

// 出入金方向类型
type TThostFtdcFundDirectionEnType byte

const THOST_FTDC_FDEN_In TThostFtdcFundDirectionEnType = '1'  // Deposit
const THOST_FTDC_FDEN_Out TThostFtdcFundDirectionEnType = '2' // Withdrawal

// 货币质押方向类型
type TThostFtdcFundMortDirectionEnType byte

const THOST_FTDC_FMDEN_In TThostFtdcFundMortDirectionEnType = '1'  // Pledge
const THOST_FTDC_FMDEN_Out TThostFtdcFundMortDirectionEnType = '2' // Redemption

// 换汇业务种类类型
type TThostFtdcSwapBusinessTypeType [3]byte

// 期权类型类型
type TThostFtdcOptionsTypeType byte

const THOST_FTDC_CP_CallOptions TThostFtdcOptionsTypeType = '1' // 看涨
const THOST_FTDC_CP_PutOptions TThostFtdcOptionsTypeType = '2'  // 看跌

// 执行方式类型
type TThostFtdcStrikeModeType byte

const THOST_FTDC_STM_Continental TThostFtdcStrikeModeType = '0' // 欧式
const THOST_FTDC_STM_American TThostFtdcStrikeModeType = '1'    // 美式
const THOST_FTDC_STM_Bermuda TThostFtdcStrikeModeType = '2'     // 百慕大

// 执行类型类型
type TThostFtdcStrikeTypeType byte

const THOST_FTDC_STT_Hedge TThostFtdcStrikeTypeType = '0' // 自身对冲
const THOST_FTDC_STT_Match TThostFtdcStrikeTypeType = '1' // 匹配执行

// 中金所期权放弃执行申请类型类型
type TThostFtdcApplyTypeType byte

const THOST_FTDC_APPT_NotStrikeNum TThostFtdcApplyTypeType = '4' // 不执行数量

// 放弃执行申请数据来源类型
type TThostFtdcGiveUpDataSourceType byte

const THOST_FTDC_GUDS_Gen TThostFtdcGiveUpDataSourceType = '0'  // 系统生成
const THOST_FTDC_GUDS_Hand TThostFtdcGiveUpDataSourceType = '1' // 手工添加

// 执行宣告系统编号类型
type TThostFtdcExecOrderSysIDType [21]byte

// 执行结果类型
type TThostFtdcExecResultType byte

const THOST_FTDC_OER_NoExec TThostFtdcExecResultType = 'n'               // 没有执行
const THOST_FTDC_OER_Canceled TThostFtdcExecResultType = 'c'             // 已经取消
const THOST_FTDC_OER_OK TThostFtdcExecResultType = '0'                   // 执行成功
const THOST_FTDC_OER_NoPosition TThostFtdcExecResultType = '1'           // 期权持仓不够
const THOST_FTDC_OER_NoDeposit TThostFtdcExecResultType = '2'            // 资金不够
const THOST_FTDC_OER_NoParticipant TThostFtdcExecResultType = '3'        // 会员不存在
const THOST_FTDC_OER_NoClient TThostFtdcExecResultType = '4'             // 客户不存在
const THOST_FTDC_OER_NoInstrument TThostFtdcExecResultType = '6'         // 合约不存在
const THOST_FTDC_OER_NoRight TThostFtdcExecResultType = '7'              // 没有执行权限
const THOST_FTDC_OER_InvalidVolume TThostFtdcExecResultType = '8'        // 不合理的数量
const THOST_FTDC_OER_NoEnoughHistoryTrade TThostFtdcExecResultType = '9' // 没有足够的历史成交
const THOST_FTDC_OER_Unknown TThostFtdcExecResultType = 'a'              // 未知

// 执行序号类型
type TThostFtdcStrikeSequenceType int32

// 执行时间类型
type TThostFtdcStrikeTimeType [13]byte

// 组合类型类型
type TThostFtdcCombinationTypeType byte

const THOST_FTDC_COMBT_Future TThostFtdcCombinationTypeType = '0' // 期货组合
const THOST_FTDC_COMBT_BUL TThostFtdcCombinationTypeType = '1'    // 垂直价差BUL
const THOST_FTDC_COMBT_BER TThostFtdcCombinationTypeType = '2'    // 垂直价差BER
const THOST_FTDC_COMBT_STD TThostFtdcCombinationTypeType = '3'    // 跨式组合
const THOST_FTDC_COMBT_STG TThostFtdcCombinationTypeType = '4'    // 宽跨式组合
const THOST_FTDC_COMBT_PRT TThostFtdcCombinationTypeType = '5'    // 备兑组合
const THOST_FTDC_COMBT_CAS TThostFtdcCombinationTypeType = '6'    // 时间价差组合
const THOST_FTDC_COMBT_OPL TThostFtdcCombinationTypeType = '7'    // 期权对锁组合
const THOST_FTDC_COMBT_BFO TThostFtdcCombinationTypeType = '8'    // 买备兑组合
const THOST_FTDC_COMBT_BLS TThostFtdcCombinationTypeType = '9'    // 买入期权垂直价差组合
const THOST_FTDC_COMBT_BES TThostFtdcCombinationTypeType = 'a'    // 卖出期权垂直价差组合

// 组合类型类型
type TThostFtdcDceCombinationTypeType byte

const THOST_FTDC_DCECOMBT_SPL TThostFtdcDceCombinationTypeType = '0' // 期货对锁组合
const THOST_FTDC_DCECOMBT_OPL TThostFtdcDceCombinationTypeType = '1' // 期权对锁组合
const THOST_FTDC_DCECOMBT_SP TThostFtdcDceCombinationTypeType = '2'  // 期货跨期组合
const THOST_FTDC_DCECOMBT_SPC TThostFtdcDceCombinationTypeType = '3' // 期货跨品种组合
const THOST_FTDC_DCECOMBT_BLS TThostFtdcDceCombinationTypeType = '4' // 买入期权垂直价差组合
const THOST_FTDC_DCECOMBT_BES TThostFtdcDceCombinationTypeType = '5' // 卖出期权垂直价差组合
const THOST_FTDC_DCECOMBT_CAS TThostFtdcDceCombinationTypeType = '6' // 期权日历价差组合
const THOST_FTDC_DCECOMBT_STD TThostFtdcDceCombinationTypeType = '7' // 期权跨式组合
const THOST_FTDC_DCECOMBT_STG TThostFtdcDceCombinationTypeType = '8' // 期权宽跨式组合
const THOST_FTDC_DCECOMBT_BFO TThostFtdcDceCombinationTypeType = '9' // 买入期货期权组合
const THOST_FTDC_DCECOMBT_SFO TThostFtdcDceCombinationTypeType = 'a' // 卖出期货期权组合

// 期权权利金价格类型类型
type TThostFtdcOptionRoyaltyPriceTypeType byte

const THOST_FTDC_ORPT_PreSettlementPrice TThostFtdcOptionRoyaltyPriceTypeType = '1'    // 昨结算价
const THOST_FTDC_ORPT_OpenPrice TThostFtdcOptionRoyaltyPriceTypeType = '4'             // 开仓价
const THOST_FTDC_ORPT_MaxPreSettlementPrice TThostFtdcOptionRoyaltyPriceTypeType = '5' // 最新价与昨结算价较大值

// 权益算法类型
type TThostFtdcBalanceAlgorithmType byte

const THOST_FTDC_BLAG_Default TThostFtdcBalanceAlgorithmType = '1'           // 不计算期权市值盈亏
const THOST_FTDC_BLAG_IncludeOptValLost TThostFtdcBalanceAlgorithmType = '2' // 计算期权市值亏损

// 执行类型类型
type TThostFtdcActionTypeType byte

const THOST_FTDC_ACTP_Exec TThostFtdcActionTypeType = '1'    // 执行
const THOST_FTDC_ACTP_Abandon TThostFtdcActionTypeType = '2' // 放弃

// 询价状态类型
type TThostFtdcForQuoteStatusType byte

const THOST_FTDC_FQST_Submitted TThostFtdcForQuoteStatusType = 'a' // 已经提交
const THOST_FTDC_FQST_Accepted TThostFtdcForQuoteStatusType = 'b'  // 已经接受
const THOST_FTDC_FQST_Rejected TThostFtdcForQuoteStatusType = 'c'  // 已经被拒绝

// 取值方式类型
type TThostFtdcValueMethodType byte

const THOST_FTDC_VM_Absolute TThostFtdcValueMethodType = '0' // 按绝对值
const THOST_FTDC_VM_Ratio TThostFtdcValueMethodType = '1'    // 按比率

// 期权行权后是否保留期货头寸的标记类型
type TThostFtdcExecOrderPositionFlagType byte

const THOST_FTDC_EOPF_Reserve TThostFtdcExecOrderPositionFlagType = '0'   // 保留
const THOST_FTDC_EOPF_UnReserve TThostFtdcExecOrderPositionFlagType = '1' // 不保留

// 期权行权后生成的头寸是否自动平仓类型
type TThostFtdcExecOrderCloseFlagType byte

const THOST_FTDC_EOCF_AutoClose TThostFtdcExecOrderCloseFlagType = '0'  // 自动平仓
const THOST_FTDC_EOCF_NotToClose TThostFtdcExecOrderCloseFlagType = '1' // 免于自动平仓

// 产品类型类型
type TThostFtdcProductTypeType byte

const THOST_FTDC_PTE_Futures TThostFtdcProductTypeType = '1' // 期货
const THOST_FTDC_PTE_Options TThostFtdcProductTypeType = '2' // 期权

// 郑商所结算文件名类型
type TThostFtdcCZCEUploadFileNameType byte

const THOST_FTDC_CUFN_CUFN_O TThostFtdcCZCEUploadFileNameType = 'O' // ^\d{8}_zz_\d{4}
const THOST_FTDC_CUFN_CUFN_T TThostFtdcCZCEUploadFileNameType = 'T' // ^\d{8}成交表
const THOST_FTDC_CUFN_CUFN_P TThostFtdcCZCEUploadFileNameType = 'P' // ^\d{8}单腿持仓表new
const THOST_FTDC_CUFN_CUFN_N TThostFtdcCZCEUploadFileNameType = 'N' // ^\d{8}非平仓了结表
const THOST_FTDC_CUFN_CUFN_L TThostFtdcCZCEUploadFileNameType = 'L' // ^\d{8}平仓表
const THOST_FTDC_CUFN_CUFN_F TThostFtdcCZCEUploadFileNameType = 'F' // ^\d{8}资金表
const THOST_FTDC_CUFN_CUFN_C TThostFtdcCZCEUploadFileNameType = 'C' // ^\d{8}组合持仓表
const THOST_FTDC_CUFN_CUFN_M TThostFtdcCZCEUploadFileNameType = 'M' // ^\d{8}保证金参数表

// 大商所结算文件名类型
type TThostFtdcDCEUploadFileNameType byte

const THOST_FTDC_DUFN_DUFN_O TThostFtdcDCEUploadFileNameType = 'O' // ^\d{8}_dl_\d{3}
const THOST_FTDC_DUFN_DUFN_T TThostFtdcDCEUploadFileNameType = 'T' // ^\d{8}_成交表
const THOST_FTDC_DUFN_DUFN_P TThostFtdcDCEUploadFileNameType = 'P' // ^\d{8}_持仓表
const THOST_FTDC_DUFN_DUFN_F TThostFtdcDCEUploadFileNameType = 'F' // ^\d{8}_资金结算表
const THOST_FTDC_DUFN_DUFN_C TThostFtdcDCEUploadFileNameType = 'C' // ^\d{8}_优惠组合持仓明细表
const THOST_FTDC_DUFN_DUFN_D TThostFtdcDCEUploadFileNameType = 'D' // ^\d{8}_持仓明细表
const THOST_FTDC_DUFN_DUFN_M TThostFtdcDCEUploadFileNameType = 'M' // ^\d{8}_保证金参数表
const THOST_FTDC_DUFN_DUFN_S TThostFtdcDCEUploadFileNameType = 'S' // ^\d{8}_期权执行表

// 上期所结算文件名类型
type TThostFtdcSHFEUploadFileNameType byte

const THOST_FTDC_SUFN_SUFN_O TThostFtdcSHFEUploadFileNameType = 'O' // ^\d{4}_\d{8}_\d{8}_DailyFundChg
const THOST_FTDC_SUFN_SUFN_T TThostFtdcSHFEUploadFileNameType = 'T' // ^\d{4}_\d{8}_\d{8}_Trade
const THOST_FTDC_SUFN_SUFN_P TThostFtdcSHFEUploadFileNameType = 'P' // ^\d{4}_\d{8}_\d{8}_SettlementDetail
const THOST_FTDC_SUFN_SUFN_F TThostFtdcSHFEUploadFileNameType = 'F' // ^\d{4}_\d{8}_\d{8}_Capital

// 中金所结算文件名类型
type TThostFtdcCFFEXUploadFileNameType byte

const THOST_FTDC_CFUFN_SUFN_T TThostFtdcCFFEXUploadFileNameType = 'T' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_Trade
const THOST_FTDC_CFUFN_SUFN_P TThostFtdcCFFEXUploadFileNameType = 'P' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_SettlementDetail
const THOST_FTDC_CFUFN_SUFN_F TThostFtdcCFFEXUploadFileNameType = 'F' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_Capital
const THOST_FTDC_CFUFN_SUFN_S TThostFtdcCFFEXUploadFileNameType = 'S' // ^\d{4}_SG\d{1}_\d{8}_\d{1}_OptionExec

// 组合指令方向类型
type TThostFtdcCombDirectionType byte

const THOST_FTDC_CMDR_Comb TThostFtdcCombDirectionType = '0'    // 申请组合
const THOST_FTDC_CMDR_UnComb TThostFtdcCombDirectionType = '1'  // 申请拆分
const THOST_FTDC_CMDR_DelComb TThostFtdcCombDirectionType = '2' // 操作员删组合单

// 行权偏移类型类型
type TThostFtdcStrikeOffsetTypeType byte

const THOST_FTDC_STOV_RealValue TThostFtdcStrikeOffsetTypeType = '1'   // 实值额
const THOST_FTDC_STOV_ProfitValue TThostFtdcStrikeOffsetTypeType = '2' // 盈利额
const THOST_FTDC_STOV_RealRatio TThostFtdcStrikeOffsetTypeType = '3'   // 实值比例
const THOST_FTDC_STOV_ProfitRatio TThostFtdcStrikeOffsetTypeType = '4' // 盈利比例

// 预约开户状态类型
type TThostFtdcReserveOpenAccStasType byte

const THOST_FTDC_ROAST_Processing TThostFtdcReserveOpenAccStasType = '0' // 等待处理中
const THOST_FTDC_ROAST_Cancelled TThostFtdcReserveOpenAccStasType = '1'  // 已撤销
const THOST_FTDC_ROAST_Opened TThostFtdcReserveOpenAccStasType = '2'     // 已开户
const THOST_FTDC_ROAST_Invalid TThostFtdcReserveOpenAccStasType = '3'    // 无效请求

// 登录备注类型
type TThostFtdcLoginRemarkType [36]byte

// 投资单元代码类型
type TThostFtdcInvestUnitIDType [17]byte

// 公告编号类型
type TThostFtdcBulletinIDType int32

// 公告类型类型
type TThostFtdcNewsTypeType [3]byte

// 紧急程度类型
type TThostFtdcNewsUrgencyType byte

// 消息摘要类型
type TThostFtdcAbstractType [81]byte

// 消息来源类型
type TThostFtdcComeFromType [21]byte

// WEB地址类型
type TThostFtdcURLLinkType [201]byte

// 长个人姓名类型
type TThostFtdcLongIndividualNameType [161]byte

// 长换汇银行账户名类型
type TThostFtdcLongFBEBankAccountNameType [161]byte

// 日期时间类型
type TThostFtdcDateTimeType [17]byte

// 弱密码来源类型
type TThostFtdcWeakPasswordSourceType byte

const THOST_FTDC_WPSR_Lib TThostFtdcWeakPasswordSourceType = '1'    // 弱密码库
const THOST_FTDC_WPSR_Manual TThostFtdcWeakPasswordSourceType = '2' // 手工录入

// 随机串类型
type TThostFtdcRandomStringType [17]byte

// 期权行权的头寸是否自对冲类型
type TThostFtdcOptSelfCloseFlagType byte

const THOST_FTDC_OSCF_CloseSelfOptionPosition TThostFtdcOptSelfCloseFlagType = '1'     // 自对冲期权仓位
const THOST_FTDC_OSCF_ReserveOptionPosition TThostFtdcOptSelfCloseFlagType = '2'       // 保留期权仓位
const THOST_FTDC_OSCF_SellCloseSelfFuturePosition TThostFtdcOptSelfCloseFlagType = '3' // 自对冲卖方履约后的期货仓位
const THOST_FTDC_OSCF_ReserveFuturePosition TThostFtdcOptSelfCloseFlagType = '4'       // 保留卖方履约后的期货仓位

// 业务类型类型
type TThostFtdcBizTypeType byte

const THOST_FTDC_BZTP_Future TThostFtdcBizTypeType = '1' // 期货
const THOST_FTDC_BZTP_Stock TThostFtdcBizTypeType = '2'  // 证券

// 用户App类型类型
type TThostFtdcAppTypeType byte

const THOST_FTDC_APP_TYPE_Investor TThostFtdcAppTypeType = '1'      // 直连的投资者
const THOST_FTDC_APP_TYPE_InvestorRelay TThostFtdcAppTypeType = '2' // 为每个投资者都创建连接的中继
const THOST_FTDC_APP_TYPE_OperatorRelay TThostFtdcAppTypeType = '3' // 所有投资者共享一个操作员连接的中继
const THOST_FTDC_APP_TYPE_UnKnown TThostFtdcAppTypeType = '4'       // 未知

// App代码类型
type TThostFtdcAppIDType [33]byte

// 系统信息长度类型
type TThostFtdcSystemInfoLenType int32

// 补充信息长度类型
type TThostFtdcAdditionalInfoLenType int32

// 交易终端系统信息类型
type TThostFtdcClientSystemInfoType [273]byte

// 系统外部信息类型
type TThostFtdcAdditionalInfoType [261]byte

// base64交易终端系统信息类型
type TThostFtdcBase64ClientSystemInfoType [365]byte

// base64系统外部信息类型
type TThostFtdcBase64AdditionalInfoType [349]byte

// 当前可用的认证模式，0代表无需认证模式 A从低位开始最后一位代表图片验证码，倒数第二位代表动态口令，倒数第三位代表短信验证码类型
type TThostFtdcCurrentAuthMethodType int32

// 图片验证信息长度类型
type TThostFtdcCaptchaInfoLenType int32

// 图片验证信息类型
type TThostFtdcCaptchaInfoType [2561]byte

// 用户短信验证码的编号类型
type TThostFtdcUserTextSeqType int32

// 握手数据内容类型
type TThostFtdcHandshakeDataType [301]byte

// 握手数据内容长度类型
type TThostFtdcHandshakeDataLenType int32

// api与front通信密钥版本号类型
type TThostFtdcCryptoKeyVersionType [31]byte

// 公钥版本号类型
type TThostFtdcRsaKeyVersionType int32

// 交易软件商ID类型
type TThostFtdcSoftwareProviderIDType [22]byte

// 信息采集时间类型
type TThostFtdcCollectTimeType [21]byte

// 查询频率类型
type TThostFtdcQueryFreqType int32

// 应答类型类型
type TThostFtdcResponseValueType byte

const THOST_FTDC_RV_Right TThostFtdcResponseValueType = '0'  // 检查成功
const THOST_FTDC_RV_Refuse TThostFtdcResponseValueType = '1' // 检查失败

// OTC成交类型类型
type TThostFtdcOTCTradeTypeType byte

const THOST_FTDC_OTC_TRDT_Block TThostFtdcOTCTradeTypeType = '0' // 大宗交易
const THOST_FTDC_OTC_TRDT_EFP TThostFtdcOTCTradeTypeType = '1'   // 期转现

// 期现风险匹配方式类型
type TThostFtdcMatchTypeType byte

const THOST_FTDC_OTC_MT_DV01 TThostFtdcMatchTypeType = '1'     // 基点价值
const THOST_FTDC_OTC_MT_ParValue TThostFtdcMatchTypeType = '2' // 面值

// OTC交易员代码类型
type TThostFtdcOTCTraderIDType [31]byte

// 期货风险值类型
type TThostFtdcRiskValueType float64

// 握手数据内容类型
type TThostFtdcIDBNameType [100]byte

// 折扣率类型
type TThostFtdcDiscountRatioType float64

// 用户终端认证方式类型
type TThostFtdcAuthTypeType byte

const THOST_FTDC_AU_WHITE TThostFtdcAuthTypeType = '0' // 白名单校验
const THOST_FTDC_AU_BLACK TThostFtdcAuthTypeType = '1' // 黑名单校验

// 合约分类方式类型
type TThostFtdcClassTypeType byte

const THOST_FTDC_INS_ALL TThostFtdcClassTypeType = '0'    // 所有合约
const THOST_FTDC_INS_FUTURE TThostFtdcClassTypeType = '1' // 期货、即期、期转现、Tas、金属指数合约
const THOST_FTDC_INS_OPTION TThostFtdcClassTypeType = '2' // 期货、现货期权合约
const THOST_FTDC_INS_COMB TThostFtdcClassTypeType = '3'   // 组合合约

// 合约交易状态分类方式类型
type TThostFtdcTradingTypeType byte

const THOST_FTDC_TD_ALL TThostFtdcTradingTypeType = '0'     // 所有状态
const THOST_FTDC_TD_TRADE TThostFtdcTradingTypeType = '1'   // 交易
const THOST_FTDC_TD_UNTRADE TThostFtdcTradingTypeType = '2' // 非交易

// 产品状态类型
type TThostFtdcProductStatusType byte

const THOST_FTDC_PS_tradeable TThostFtdcProductStatusType = '1'   // 可交易
const THOST_FTDC_PS_untradeable TThostFtdcProductStatusType = '2' // 不可交易

// 追平状态类型
type TThostFtdcSyncDeltaStatusType byte

const THOST_FTDC_SDS_Readable TThostFtdcSyncDeltaStatusType = '1' // 交易可读
const THOST_FTDC_SDS_Reading TThostFtdcSyncDeltaStatusType = '2'  // 交易在读
const THOST_FTDC_SDS_Readend TThostFtdcSyncDeltaStatusType = '3'  // 交易读取完成
const THOST_FTDC_SDS_OptErr TThostFtdcSyncDeltaStatusType = 'e'   // 追平失败 交易本地状态结算不存在

// 操作标志类型
type TThostFtdcActionDirectionType byte

const THOST_FTDC_ACD_Add TThostFtdcActionDirectionType = '1' // 增加
const THOST_FTDC_ACD_Del TThostFtdcActionDirectionType = '2' // 删除
const THOST_FTDC_ACD_Upd TThostFtdcActionDirectionType = '3' // 更新

// 撤单时选择席位算法类型
type TThostFtdcOrderCancelAlgType byte

const THOST_FTDC_OAC_Balance TThostFtdcOrderCancelAlgType = '1'   // 轮询席位撤单
const THOST_FTDC_OAC_OrigFirst TThostFtdcOrderCancelAlgType = '2' // 优先原报单席位撤单

// 追平描述类型
type TThostFtdcSyncDescriptionType [257]byte

// 通用int类型类型
type TThostFtdcCommonIntType int32

// 系统版本类型
type TThostFtdcSysVersionType [41]byte

// 开仓量限制粒度类型
type TThostFtdcOpenLimitControlLevelType byte

const THOST_FTDC_PLCL_None TThostFtdcOpenLimitControlLevelType = '0'    // 不控制
const THOST_FTDC_PLCL_Product TThostFtdcOpenLimitControlLevelType = '1' // 产品级别
const THOST_FTDC_PLCL_Inst TThostFtdcOpenLimitControlLevelType = '2'    // 合约级别

// 报单频率控制粒度类型
type TThostFtdcOrderFreqControlLevelType byte

const THOST_FTDC_OFCL_None TThostFtdcOrderFreqControlLevelType = '0'    // 不控制
const THOST_FTDC_OFCL_Product TThostFtdcOrderFreqControlLevelType = '1' // 产品级别
const THOST_FTDC_OFCL_Inst TThostFtdcOrderFreqControlLevelType = '2'    // 合约级别

// 枚举bool类型类型
type TThostFtdcEnumBoolType byte

const THOST_FTDC_EBL_False TThostFtdcEnumBoolType = '0' // false
const THOST_FTDC_EBL_True TThostFtdcEnumBoolType = '1'  // true