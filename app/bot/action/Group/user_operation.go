package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/app/v1/user/action/BalanceAction"
	"main.go/tuuz/Calc"
	"math"
)

func App_ban_user(self_id, group_id, user_id interface{}, auto_retract bool, groupfunction map[string]interface{}, reason string) {
	at := service.Serv_at(user_id)
	time := GroupBanModel.Api_count(group_id, user_id)
	GroupBanModel.Api_insert(group_id, user_id)
	left_time := groupfunction["ban_limit"].(int64) - 1
	var balance BalanceAction.Interface
	if left_time > 0 {
		bal, _ := balance.App_check_balance(user_id)
		balance_decr := float64(time) * math.Pow10(int(time))
		balance_left := bal - balance_decr
		if balance_left < 0 {
			api.Sendgroupmsg(self_id, group_id, at+"这是你第:"+Calc.Any2String(time+1)+"次，接受惩罚\n"+"本次惩罚原因："+reason+"\n你还剩下："+Calc.Any2String(left_time)+"点生命值", auto_retract)
			api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		} else {
			BalanceAction.App_single_balance(user_id, "", -math.Abs(balance_decr), "扣分")
			api.Sendgroupmsg(self_id, group_id, at+"这是你第:"+Calc.Any2String(time+1)+"次扣分\n"+"本次扣分原因："+reason+"\n你还剩下："+Calc.Any2String(balance_left)+"分", auto_retract)
		}
	} else {
		App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, reason+"\n且他已经没有生命值了")
	}
}

func App_kick_user(self_id, group_id, user_id interface{}, auto_retract bool, groupfunction map[string]interface{}, reason string) {
	auto_kick_out := groupfunction["auto_kick_out"].(int64)
	if auto_kick_out == 1 {
		gm := GroupMemberModel.Api_find(group_id, user_id)
		if len(gm) > 0 {
			nickname := Calc.Any2String(gm["nickname"])
			api.Sendgroupmsg(self_id, group_id, nickname+"被T出，原因为："+reason, auto_retract)
		} else {
			api.Sendgroupmsg(self_id, group_id, Calc.Any2String(user_id)+"被T出，原因为："+reason, auto_retract)
		}
	} else {
		if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {

		} else {
			if GroupBanPermenentModel.Api_insert(group_id, user_id) {
				at := service.Serv_at(user_id)
				api.Sendgroupmsg(self_id, group_id, at+"你已经低于生命值，现在将你加入永久禁言列表，仅允许管理员解禁", auto_retract)
			}
		}
	}
}

func Api_retract_send(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {

}
