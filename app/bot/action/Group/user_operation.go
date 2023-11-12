package Group

import (
	"fmt"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/DaojuModel"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupDaojuModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz"

	"main.go/tuuz/Date"
	"math"
	"time"
)

func App_ban_user(self_id, group_id, user_id any, auto_retract bool, groupfunction map[string]any, reason string) {
	at := service.Serv_at(user_id)
	time := GroupBanModel.Api_count(group_id, user_id)
	GroupBanModel.Api_insert(group_id, user_id)
	left_time := groupfunction["ban_limit"].(int64) - 1 - time
	var daoju GroupDaojuModel.Interface
	daoju.Db = tuuz.Db()
	daoju.Db.Begin()
	dj_data := DaojuModel.Api_find_byName("anti_ban")
	user_dj := daoju.Api_find(group_id, user_id, dj_data["id"])
	if len(user_dj) > 0 && user_dj["num"].(int64) > 0 {
		if daoju.Api_decr(group_id, user_id, dj_data["id"]) {
			dj_left := daoju.Api_value_num(group_id, user_id, dj_data["id"])
			daoju.Db.Commit()
			str := "\r\n[" + Calc.Any2String(dj_data["cname"]) + "]还剩下" + Calc.Any2String(dj_left)
			AutoMessage(self_id, group_id, user_id, app_default.Daoju_use_for_ban+str, groupfunction)
			return
		}
	}
	daoju.Db.Rollback()
	if left_time > 0 {
		var balance GroupBalanceModel.Interface
		balance.Db = tuuz.Db()
		groupbal := balance.Api_value_balance(group_id, user_id)
		if groupbal != nil {
			bal, _ := groupbal.(float64)
			balance_decr := float64(time+1) * 10
			balance_left := bal - balance_decr
			fmt.Println("当前积分", bal, balance_decr, balance_left)
			if balance_left >= 0 {
				balance.Api_decr(group_id, user_id, math.Abs(balance_decr))
				go iapi.Api.Sendgroupmsg(self_id, group_id, at+Calc.Any2String(time+1)+"这是你第:"+"次扣分，扣除"+Calc.Any2String(balance_decr)+"分\n"+"本次扣分原因："+reason+"\n你还剩下："+
					""+Calc.Any2String(balance_left)+"分", auto_retract)
				return
			}
		}
		go iapi.Api.Sendgroupmsg(self_id, group_id, at+"这是你第:"+Calc.Any2String(time+1)+"次，接受惩罚\n"+"本次惩罚原因："+reason+"\n你还剩下："+Calc.Any2String(left_time)+"点生命值", auto_retract)
		iapi.Api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
	} else {
		App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, reason+"\n且他已经没有生命值了")
	}
}

func App_kick_user(self_id, group_id, user_id any, auto_retract bool, groupfunction map[string]any, reason string) {
	var daoju GroupDaojuModel.Interface
	daoju.Db = tuuz.Db()
	daoju.Db.Begin()
	dj_data := DaojuModel.Api_find_byName("anti_kick")
	user_dj := daoju.Api_find(group_id, user_id, dj_data["id"])
	if len(user_dj) > 0 && user_dj["num"].(int64) > 0 {
		if daoju.Api_decr(group_id, user_id, dj_data["id"]) {
			dj_left := daoju.Api_value_num(group_id, user_id, dj_data["id"])
			daoju.Db.Commit()
			str := "\r\n[" + Calc.Any2String(dj_data["cname"]) + "]还剩下" + Calc.Any2String(dj_left)
			AutoMessage(self_id, group_id, user_id, app_default.Daoju_use_for_kick+str, groupfunction)
			return
		}
	}
	daoju.Db.Rollback()
	auto_kick_out := groupfunction["auto_kick_out"].(int64)
	str := ""
	if auto_kick_out == 1 {
		if groupfunction["kick_to_black"].(int64) == 1 {
			str = "并被拉黑"
			GroupBlackListModel.Api_insert(group_id, user_id, self_id)
		}
		gm := GroupMemberModel.Api_find(group_id, user_id)
		iapi.Api.SetGroupKick(self_id, group_id, user_id, false)
		if len(gm) > 0 {
			nickname := Calc.Any2String(gm["nickname"])
			go iapi.Api.Sendgroupmsg(self_id, group_id, nickname+"被T出"+str+"，原因为："+reason, auto_retract)
		} else {
			go iapi.Api.Sendgroupmsg(self_id, group_id, Calc.Any2String(user_id)+"被T出"+str+"，原因为："+reason, auto_retract)
		}
	} else {
		if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {

		} else {
			if GroupBanPermenentModel.Api_insert(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time-86400) {
				at := service.Serv_at(user_id)
				iapi.Api.SetGroupBan(self_id, group_id, user_id, app_conf.Auto_ban_time)
				go iapi.Api.Sendgroupmsg(self_id, group_id, at+"你已经低于生命值，现在将你加入永久禁言列表，仅允许管理员解禁", auto_retract)
			}
		}
	}
}

func Api_retract_send(bot, gid, uid int, req int, random int, groupmember map[string]any, groupfunction map[string]any) {

}

func App_drcrease_member(self_id, group_id, user_id any, groupfunction map[string]any, reason string) {
	group_list_data := GroupListModel.Api_find(group_id)
	if len(group_list_data) > 0 {
		group_member_count := GroupMemberModel.Api_count_byGroupIdAndRole(group_id, nil)
		if group_list_data["max_member_count"].(int64) < group_member_count {
			group_member_datas := GroupMemberModel.Api_select_byGroupId(group_id, "last_sent_time desc", int(group_list_data["max_member_count"].(int64)-20), 2)
			if len(group_member_datas) > 0 {
				go iapi.Api.Sendgroupmsg(self_id, group_id, "本群将被清除"+Calc.Any2String(len(group_member_datas))+
					"人，\n第一个被T出的人为:"+Calc.Any2String(group_member_datas[0]["nickname"])+"，他最后一次说话是在："+
					Date.Date_format_second(group_member_datas[0]["last_date"].(time.Time))+
					"\n最后一个被清除的为:"+Calc.Any2String(group_member_datas[len(group_member_datas)-1]["nickname"])+
					"，他最后一次说话是在："+Date.Date_format_second(group_member_datas[len(group_member_datas)-1]["last_date"].(time.Time)), false)
				for _, data := range group_member_datas {
					ok, err := iapi.Api.SetGroupKick(self_id, group_id, data["user_id"], false)
					if err != nil {
						fmt.Println(err)
					} else {
						if ok {
							GroupMemberModel.Api_delete_byUid(self_id, group_id, data["user_id"])
						}
					}
				}
			} else {
				go iapi.Api.Sendgroupmsg(self_id, group_id, "没有需要清理的人", true)
			}
		} else {
			go iapi.Api.Sendgroupmsg(self_id, group_id, "未达到清理下限无需调用", true)
		}
	} else {
		go iapi.Api.Sendgroupmsg(self_id, group_id, "未找到当前群信息", true)
	}
}
