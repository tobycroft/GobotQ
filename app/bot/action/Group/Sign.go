package Group

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupSignModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

func App_group_sign(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	sign := GroupSignModel.Api_find(gid, uid)
	private_mode := false
	if groupfunction["sign_send_private"].(int64) == 1 {
		private_mode = true
	}
	auto_retract := false
	if groupfunction["sign_send_private"].(int64) == 1 {
		auto_retract = true
		var ret api.Retract_group
		ret.Group = gid
		ret.Fromqq = bot
		ret.Random = random
		ret.Req = req
		api.Retract_chan_group <- ret
	}
	if len(sign) > 0 {
		if private_mode {
			api.Sendgrouptempmsg(bot, gid, uid, "你今天已经签到过了")
		} else {
			at := service.Serv_at(uid)
			api.Sendgroupmsg(bot, gid, "你今天已经签到过了"+at, auto_retract)
		}
	} else {
		rank := GroupSignModel.Api_count(gid)
		order := rank + 1
		amount := app_conf.Group_Sign_incr - rank
		if amount < 0 {
			amount = 1
		}
		group_model := GroupBalanceModel.Api_find(gid, uid)
		db := tuuz.Db()
		db.Begin()
		var gbp GroupBalanceModel.Interface
		gbp.Db = db
		if len(group_model) < 1 {
			if !gbp.Api_insert(gid, uid) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,写入失败"), tuuz.FUNCTION_ALL())
				return
			}
		}

		//加分模式

		if !gbp.Api_incr(gid, uid, amount) {
			db.Rollback()
			Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
			return
		}

		var gsp GroupSignModel.Interface
		gsp.Db = db
		if !gsp.Api_insert(gid, uid) {
			db.Rollback()
			Log.Errs(errors.New("GroupSignModel,插入失败"), tuuz.FUNCTION_ALL())
			return
		} else {
			db.Commit()
			if private_mode {
				api.Sendgrouptempmsg(bot, gid, uid, "签到成功,您是第"+Calc.Int642String(order)+"个签到,积分奖励"+Calc.Int642String(amount))
			} else {
				at := service.Serv_at(uid)
				api.Sendgroupmsg(bot, gid, "签到成功"+at+",您是第"+Calc.Int642String(order)+"个签到,积分奖励"+Calc.Int642String(amount), auto_retract)
			}
		}
	}
}
