package Group

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupSignModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func App_group_sign(bot, gid, uid interface{}) {
	sign := GroupSignModel.Api_find(gid, uid)
	if len(sign) > 0 {
		at := service.Serv_at(uid)
		api.Sendgroupmsg(bot, gid, at+"你今天已经签到过了")
	} else {
		group_model := GroupBalanceModel.Api_find(gid, uid)
		db := tuuz.Db()
		db.Begin()
		var gbp GroupBalanceModel.Interface
		gbp.Db = db
		if len(group_model) < 1 {
			if !gbp.Api_insert(gid, uid) {
				Log.Errs(errors.New("GroupBalanceModel,写入失败"), tuuz.FUNCTION_ALL())
				return
			}
		}
		if !gbp.Api_incr(gid, uid, app_conf.Group_Sign_incr) {
			Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
			return
		}

		var gsp GroupSignModel.Interface
		gsp.Db = db
		if !gsp.Api_insert(gid, uid) {
			Log.Errs(errors.New("GroupSignModel,插入失败"), tuuz.FUNCTION_ALL())
			return
		} else {
			at := service.Serv_at(uid)
			api.Sendgroupmsg(bot, gid, at+"签到成功")
		}
	}
}
