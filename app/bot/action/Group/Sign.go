package Group

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupSignModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

func App_group_sign(self_id, group_id, user_id, message_id int64, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	sign := GroupSignModel.Api_find(group_id, user_id)
	//private_mode := false
	//if groupfunction["sign_send_private"].(int64) == 1 {
	//	private_mode = true
	//}
	if groupfunction["sign_send_retract"].(int64) == 1 {
		var ret api.Struct_Retract
		ret.MessageId = message_id
		ret.Self_id = self_id
		api.Retract_chan <- ret
	}
	if len(sign) > 0 {
		at := service.Serv_at(user_id)
		AutoMessage(self_id, group_id, user_id, "你今天已经签到过了"+at, groupfunction)
	} else {
		rank := GroupSignModel.Api_count(group_id)
		order := rank + 1
		amount := app_conf.Group_Sign_incr - rank
		if amount <= 0 {
			amount = 1
		}
		db := tuuz.Db()
		db.Begin()
		var gsp GroupSignModel.Interface
		gsp.Db = db
		if !gsp.Api_insert(group_id, user_id) {
			db.Rollback()
			Log.Errs(errors.New("GroupSignModel,插入失败"), tuuz.FUNCTION_ALL())
			return
		}
		at := service.Serv_at(user_id)
		if len(GroupBanModel.Api_find(group_id, user_id)) > 1 {
			//奖励生命模式
			if GroupBanModel.Api_delete_userId(group_id, user_id) {
				AutoMessage(self_id, group_id, user_id, at+",您是今日第"+Calc.Int642String(order)+"个签到,生命值已经补满", groupfunction)
			}
			db.Commit()
		} else {
			//加分模式
			group_model := GroupBalanceModel.Api_find(group_id, user_id)
			var gbp GroupBalanceModel.Interface
			gbp.Db = db
			if len(group_model) < 1 {
				if !gbp.Api_insert(group_id, user_id) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,写入失败"), tuuz.FUNCTION_ALL())
					return
				}
			}
			if !gbp.Api_incr(group_id, user_id, amount) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
				return
			}
			db.Commit()
			AutoMessage(self_id, group_id, user_id, at+",您是今日第"+Calc.Int642String(order)+"个签到,威望奖励"+Calc.Int642String(amount)+",现有威望："+Calc.Any2String(group_model["balance"]), groupfunction)
		}
	}
}
