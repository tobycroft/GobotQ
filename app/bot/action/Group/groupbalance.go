package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/tuuz/Calc"
)

func App_check_balance(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	auto_retract := false
	if groupfunction["sign_send_retract"].(int64) == 1 {
		auto_retract = true
		var ret api.Retract_group
		ret.Group = gid
		ret.Fromqq = bot
		ret.Random = random
		ret.Req = req
		api.Retract_chan_group <- ret
	}
	gbl := GroupBalanceModel.Api_find(gid, uid)
	str := "您当前拥有" + Calc.Any2String(gbl["balance"]) + "分"
	api.Sendgroupmsg(bot, gid, str, auto_retract)
}

func App_check_rank(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	auto_retract := false
	if groupfunction["sign_send_retract"].(int64) == 1 {
		auto_retract = true
		var ret api.Retract_group
		ret.Group = gid
		ret.Fromqq = bot
		ret.Random = random
		ret.Req = req
		api.Retract_chan_group <- ret
	}
	gbl := GroupBalanceModel.Api_select(gid, 10)
	str := ""
	for i1, i2 := range gbl {
		user := GroupMemberModel.Api_find(gid, i2["uid"].(int64))
		if len(user) > 0 {
			if len(Calc.Any2String(user["card"])) > 2 && Calc.Any2String(user["card"]) != "null" {
				str += "第" + Calc.Int2String(i1+1) + "名：" + user["card"].(string) + "，" + Calc.Any2String(i2["balance"]) + "分" + "\r\n"
			} else {
				str += "第" + Calc.Int2String(i1+1) + "名：" + user["nickname"].(string) + "，" + Calc.Any2String(i2["balance"]) + "分" + "\r\n"
			}
		}
	}
	api.Sendgroupmsg(bot, gid, str, auto_retract)
}
