package Group

import (
	"fmt"
	"main.go/app/bot/api"
	"main.go/app/bot/model/DaojuModel"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupDaojuModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"math"
	"time"
)

func App_ban_user(self_id, group_id, user_id interface{}, auto_retract bool, groupfunction map[string]interface{}, reason string) {
	at := service.Serv_at(user_id)
	time := GroupBanModel.Api_count(group_id, user_id)
	GroupBanModel.Api_insert(group_id, user_id)
	left_time := groupfunction["ban_limit"].(int64) - 1 - time
	db := tuuz.Db()
	db.Begin()
	var balance GroupBalanceModel.Interface
	balance.Db = db
	var daoju GroupDaojuModel.Interface
	daoju.Db = db
	dj_data := DaojuModel.Api_find_byName("anti_ban")
	user_dj := daoju.Api_find(group_id, user_id, dj_data["id"])
	if len(user_dj) > 0 && user_dj["num"].(int64) > 0 {
		if daoju.Api_decr(group_id, user_id, dj_data["id"]) {
			dj_left := daoju.Api_value(group_id, user_id, dj_data["id"])
			db.Commit()
			str := "\r\n[" + Calc.Any2String(dj_data["cname"]) + "]还剩下" + Calc.Any2String(dj_left)
			AutoMessage(self_id, group_id, user_id, app_default.Daoju_use_for_ban+str, groupfunction)
			return
		}
	}
	if left_time > 0 {
		groupbal := GroupBalanceModel.Api_value_balance(group_id, user_id)
		if groupbal != nil {
			bal, _ := groupbal.(float64)
			balance_decr := float64(time+1) * 10
			balance_left := bal - balance_decr
			fmt.Println("当前积分", bal, balance_decr, balance_left)
			if balance_left >= 0 {
				balance.Api_decr(group_id, user_id, math.Abs(balance_decr))
				api.Sendgroupmsg(self_id, group_id, at+"这是你第:"+Calc.Any2String(time+1)+"次扣分，扣除"+Calc.Any2String(balance_decr)+"分\n"+"本次扣分原因："+reason+"\n你还剩下："+
					""+Calc.Any2String(balance_left)+"分", auto_retract)
				return
			}
		}
		api.Sendgroupmsg(self_id, group_id, at+"这是你第:"+Calc.Any2String(time+1)+"次，接受惩罚\n"+"本次惩罚原因："+reason+"\n你还剩下："+Calc.Any2String(left_time)+"点生命值", auto_retract)
		api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
	} else {
		App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, reason+"\n且他已经没有生命值了")
	}
}

func App_kick_user(self_id, group_id, user_id interface{}, auto_retract bool, groupfunction map[string]interface{}, reason string) {
	var daoju GroupDaojuModel.Interface
	daoju.Db = tuuz.Db()
	dj_data := DaojuModel.Api_find_byName("anti_kick")
	user_dj := daoju.Api_find(group_id, user_id, dj_data["id"])
	if len(user_dj) > 0 && user_dj["num"].(int64) > 0 {
		if daoju.Api_decr(group_id, user_id, dj_data["id"]) {
			dj_left := daoju.Api_value(group_id, user_id, dj_data["id"])
			str := "\r\n[" + Calc.Any2String(dj_data["cname"]) + "]还剩下" + Calc.Any2String(dj_left)
			AutoMessage(self_id, group_id, user_id, app_default.Daoju_use_for_kick+str, groupfunction)
			return
		}
	}

	auto_kick_out := groupfunction["auto_kick_out"].(int64)
	str := ""
	if auto_kick_out == 1 {
		if groupfunction["kick_to_black"].(int64) == 1 {
			str = "并被拉黑"
			GroupBlackListModel.Api_insert(group_id, user_id, self_id)
		}
		gm := GroupMemberModel.Api_find(group_id, user_id)
		api.SetGroupKick(self_id, group_id, user_id, false)
		if len(gm) > 0 {
			nickname := Calc.Any2String(gm["nickname"])
			api.Sendgroupmsg(self_id, group_id, nickname+"被T出"+str+"，原因为："+reason, auto_retract)
		} else {
			api.Sendgroupmsg(self_id, group_id, Calc.Any2String(user_id)+"被T出"+str+"，原因为："+reason, auto_retract)
		}
	} else {
		if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {

		} else {
			if GroupBanPermenentModel.Api_insert(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time-86400) {
				at := service.Serv_at(user_id)
				api.SetGroupBan(self_id, group_id, user_id, app_conf.Auto_ban_time)
				api.Sendgroupmsg(self_id, group_id, at+"你已经低于生命值，现在将你加入永久禁言列表，仅允许管理员解禁", auto_retract)
			}
		}
	}
}

func Api_retract_send(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {

}
