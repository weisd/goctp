#include "../CTPv6.6.9_20220922/ThostFtdcTraderApi.h"

class Trade: CThostFtdcTraderSpi{
public:
	[[ range .On ]]// [[ .Comment ]]    
    typedef void [[ .Name ]]Type([[ range $idx, $param := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Type ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ end ]]);
    void *_[[ .Name ]];
    virtual void [[ .Name ]]([[ range $idx, $param := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Type ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ end ]]){
        if (_[[ .Name ]]) {
			(([[ .Name ]]Type*)_[[ .Name ]])([[ range $idx, $param := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]][[ end ]]);
		}
    }
	[[ end ]]
};