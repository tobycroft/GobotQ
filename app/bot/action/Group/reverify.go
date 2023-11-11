package Group

import (
	"errors"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/apipost"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"

	"main.go/tuuz/Redis"
	"strings"
	"time"
)

func App_reverify(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]any, groupfunction map[string]any) {
	_, err := reverify(self_id, group_id, user_id, message, false, false)
	if err != nil {
		AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
	}
}

func App_reverify_death(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]any, groupfunction map[string]any) {
	_, err := reverify(self_id, group_id, user_id, message, true, false)
	if err != nil {
		AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
	}
}

func App_reverify_force(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]any, groupfunction map[string]any) {
	_, err := reverify(self_id, group_id, user_id, message, false, true)
	if err != nil {
		AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
	}
}

func reverify(self_id, group_id, user_id any, send_to_message string, kick, force bool) (string, error) {
	qq := service.Serv_get_qq(send_to_message)
	cq_mess, to_user_id := service.Serv_at_who(send_to_message)
	qq_num := ""
	if to_user_id != "" {
		qq_num = to_user_id
		send_to_message = strings.ReplaceAll(send_to_message, cq_mess, "")
	} else if qq != "" {
		qq_num = qq
		send_to_message = strings.ReplaceAll(send_to_message, qq, "")
	} else {
		return "", errors.New("没有找到验证人,请重新at")
	}
	member := GroupMemberModel.Api_find(group_id, qq_num)
	if len(member) < 1 {
		return "", errors.New("群成员不在群内")
	}
	user := GroupBanPermenentModel.Api_find(group_id, member["user_id"])
	go apipost.ApiPost{}.SetGroupBan(self_id, group_id, member["user_id"], 0)
	if len(user) > 0 || force {
		if len(user) < 1 {
			GroupBanPermenentModel.Api_insert(group_id, member["user_id"], time.Now().Unix()+app_conf.Auto_ban_time-86400)
		}
		num := Calc.Rand(1000, 9999)
		Redis.String_set("verify_"+Calc.Any2String(group_id)+"_"+Calc.Any2String(member["user_id"]), num, app_conf.Retract_time_second+10)
		err := Redis.String_set("ban_"+Calc.Any2String(group_id)+"_"+Calc.Any2String(member["user_id"]), num, 3600)
		if err != nil {
			return "", err
		}
		at := service.Serv_at(member["user_id"])
		go apipost.ApiPost{}.Sendgroupmsg(self_id, group_id, at+"你已被临时解禁，请在120秒内在群内输入验证码数字：\n"+Calc.Any2String(num), true)
		go func(self_id, group_id, user_id any, kick bool) {
			time.Sleep(120 * time.Second)
			ok, err := Redis.String_getBool("ban_" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
			if err != nil {
			} else {
				if ok {
					go apipost.ApiPost{}.Sendgroupmsg(self_id, group_id, at+"看起来你没有完成活人验证，我现在将你加入永久小黑屋，但是你依然可以让其他管理员帮你解除", true)
					if kick {
						apipost.ApiPost{}.SetGroupKick(self_id, group_id, user_id, false)
					} else {
						apipost.ApiPost{}.SetGroupBan(self_id, group_id, user_id, app_conf.Auto_ban_time)
					}
				}
			}
		}(self_id, group_id, member["user_id"], kick)
		return "", nil
	} else {
		return "", errors.New("群成员没有在小黑屋内")
	}
}
