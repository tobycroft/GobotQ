package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/tuuz/Calc"
)

func App_check_balance(bot, gid, uid interface{}) {
	gbl := GroupBalanceModel.Api_find(gid, uid)
	str := "您当前拥有" + Calc.Any2String(gbl["balance"]) + "分"
	api.Sendgroupmsg(bot, gid, str, true)
}
