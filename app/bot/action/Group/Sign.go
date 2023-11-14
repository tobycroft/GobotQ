package Group

import (
	"errors"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupSignModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/tuuz"

	"main.go/tuuz/Date"
	"main.go/tuuz/Log"
)

func App_group_sign(self_id, group_id, user_id, message_id int64, groupmember map[string]any, groupfunction map[string]any) {
	var gsm GroupSignModel.Interface
	db := tuuz.Db()
	db.Begin()
	gsm.Db = db
	sign := gsm.Api_find(group_id, user_id)
	//private_mode := false
	//if groupfunction["sign_send_private"].(int64) == 1 {
	//	private_mode = true
	//}
	go func(self_id, group_id, user_id, message_id int64, groupmember map[string]any, groupfunction map[string]any) {
		if groupfunction["sign_send_retract"].(int64) == 1 {
			var ret iapi.Struct_Retract
			ret.MessageId = message_id
			ret.SelfId = self_id
			iapi.Retract_chan <- ret
		}
	}(self_id, group_id, user_id, message_id, groupmember, groupfunction)
	if len(sign) > 0 {
		db.Rollback()
		at := service.Serv_at(user_id)
		AutoMessage(self_id, group_id, user_id, "你今天已经签到过了"+at, groupfunction)
	} else {
		rank := gsm.Api_count(group_id)
		order := rank + 1
		amount := float64(app_conf.Group_Sign_incr - rank)
		if amount <= 0 {
			amount = 1
		}
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
			yesterday := Date.Yesterday()
			yesterday_sign := gsm.Api_count_userId(group_id, user_id, yesterday)
			week := Date.WeekBefore()
			week_sign := gsm.Api_count_userId(group_id, user_id, week)
			var groupbalance GroupBalanceModel.Interface
			groupbalance.Db = db
			group_model := groupbalance.Api_find(group_id, user_id)
			rest_bal := float64(0)
			if group_model["balance"] == nil {
				rest_bal = 0
			} else {
				rest_bal = group_model["balance"].(float64)
			}
			rank := groupbalance.Api_count_gt_balance(group_id, rest_bal)
			if len(group_model) < 1 {
				if !groupbalance.Api_insert(group_id, user_id) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,写入失败"), tuuz.FUNCTION_ALL())
					return
				}
			}

			if yesterday_sign > 1 {
				if week_sign > 6 {
					if !groupbalance.Api_incr(group_id, user_id, amount+7) {
						db.Rollback()
						Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
						return
					}
					AutoMessage(self_id, group_id, user_id, at+",您是今日第"+Calc.Int642String(order)+"个签到,"+
						"威望奖励"+Calc.Float642String(amount)+",连续签到"+Calc.Any2String(week_sign)+"天,"+"额外奖励＋7"+
						"现有威望："+Calc.Any2String(rest_bal+amount)+",排名第："+Calc.Int642String(rank+1), groupfunction)
				} else {
					if !groupbalance.Api_incr(group_id, user_id, amount+float64(week_sign)) {
						db.Rollback()
						Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
						return
					}
					AutoMessage(self_id, group_id, user_id, at+",您是今日第"+Calc.Int642String(order)+"个签到"+
						"威望奖励"+Calc.Float642String(amount)+",连续签到"+Calc.Any2String(week_sign)+"天,"+"额外奖励＋"+Calc.Any2String(week_sign)+
						"现有威望："+Calc.Any2String(rest_bal+amount)+",排名第："+Calc.Int642String(rank+1), groupfunction)
				}
			} else {
				if !groupbalance.Api_incr(group_id, user_id, amount+float64(week_sign)) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				AutoMessage(self_id, group_id, user_id, at+",您是今日第"+Calc.Int642String(order)+"个签到,威望奖励"+Calc.Float642String(amount)+","+
					"现有威望："+Calc.Any2String(rest_bal+amount)+",排名第："+Calc.Int642String(rank+1)+",明日继续签到可堆叠奖励呢！", groupfunction)
			}
			db.Commit()
		}
	}
}
