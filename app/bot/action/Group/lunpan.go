package Group

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupLunpanModel"
	"main.go/app/bot/service"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"regexp"
)

func App_group_lunpan(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	sign := GroupLunpanModel.Api_find(group_id, user_id)
	if groupfunction["sign_send_retract"].(int64) == 1 {
		var ret api.Struct_Retract
		ret.MessageId = message_id
		ret.Self_id = self_id
		api.Retract_chan <- ret
	}
	if len(sign) > 0 {
		at := service.Serv_at(user_id)
		AutoMessage(self_id, group_id, user_id, "你今天已经挑战过了，请明天再来"+at, groupfunction)
	} else {
		group_model := GroupBalanceModel.Api_find(group_id, user_id)
		if group_model["balance"] == nil {
			group_model["balance"] = 0
		}
		db := tuuz.Db()
		db.Begin()
		var gbp GroupBalanceModel.Interface
		gbp.Db = db
		if len(group_model) < 1 {
			if !gbp.Api_insert(group_id, user_id) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,写入失败"), tuuz.FUNCTION_ALL())
				return
			}
		}

		reg := regexp.MustCompile("[0-9]")
		active := reg.MatchString(message)
		if active {
			reg := regexp.MustCompile("[^0-9]")
			active := reg.MatchString(message)
			new_text := reg.ReplaceAllString(message, "")
			//左轮模式
		} else {
			//普通模式
			rand := Calc.Rand(0, 100)
			amount := 0
			str := ""
			if rand <= 1 {
				amount = amount - rand
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "完胜,当前余额:" + Calc.Any2String(group_model["balance"]) + "十倍奖励！"
			} else if rand > 1 && rand <= 20 {
				amount = amount - rand
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "小败,当前余额:" + Calc.Any2String(group_model["balance"]) + ",扣除:" + Calc.Any2String(amount) + ";"
			} else if rand > 20 && rand <= 50 {
				amount = amount + 2
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "小胜,当前余额:" + Calc.Any2String(group_model["balance"]) + ",赢得2;"
			} else if rand > 50 && rand <= 85 {
				amount = amount + 5
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "胜利,当前余额:" + Calc.Any2String(group_model["balance"]) + ",赢得5;"
			} else if rand > 85 && rand <= 95 {
				amount = amount - rand
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "大胜,当前余额:" + Calc.Any2String(group_model["balance"]) + ",赢得10！"
			} else if rand > 95 && rand <= 99 {
				amount = amount - rand
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "轮盘大败,当前余额:" + Calc.Any2String(group_model["balance"]) + ",现在扣除一半。"
			} else {
				amount = amount - rand
				if !gbp.Api_incr(group_id, user_id, amount) {
					db.Rollback()
					Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
					return
				}
				at := service.Serv_at(user_id)
				str += at + "轮盘完败,你的余额已不复存在。"
			}
			count_lunpan := GroupLunpanModel.Api_count_userId(group_id, user_id)
			if count_lunpan < 1 {
				str += "\n这是你第一次参与轮盘，下次你可以用“轮盘[模式字母][数字]，" +
					"\n例如轮盘A10，来挑战1/6的获胜几率，挑战成功，奖励1/6押注威望，" +
					"\n同时你可以使用轮盘B10来挑战2/6的胜率，获得2/6的奖励，" +
					"\n可选模式有ABCDE，挑战威望无上限，你可以使用威望查询来查看自己的可用情况" +
					"\n觉得自己运气还不错的话可以试试哦~"
			}
			AutoMessage(self_id, group_id, user_id, str, groupfunction)
		}
		var gsp GroupLunpanModel.Interface
		gsp.Db = db
		if !gsp.Api_insert(group_id, user_id) {
			db.Rollback()
			Log.Errs(errors.New("GroupLunpanModel,插入失败"), tuuz.FUNCTION_ALL())
			return
		} else {
			db.Commit()

		}
	}
}
