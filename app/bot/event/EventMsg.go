package event

import (
	"main.go/app/bot/action/Private"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupKickModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/model/GroupMsgModel"
	"main.go/app/bot/service"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"time"
)

type EM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN      int    `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"FromQQ"`
	OperateQQ struct {
		UIN      int    `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"OperateQQ"`
	LogonQQ   int `json:"LogonQQ"`
	FromGroup struct {
		GIN  int    `json:"GIN"`
		Name string `json:"Name"`
	} `json:"FromGroup"`
	Msg struct {
		Seq       int    `json:"Seq"`
		TimeStamp int    `json:"TimeStamp"`
		Type      int    `json:"Type"`
		SubType   int    `json:"SubType"`
		Text      string `json:"Text"`
	} `json:"Msg"`
}

func EventMsg(em EM) {
	//operator := em.OperateQQ.UIN
	//text := em.Msg.Text
	seq := em.Msg.Seq
	bot := em.LogonQQ
	uid := em.FromQQ.UIN
	gid := em.FromGroup.GIN
	Type := em.Msg.Type
	operator := em.OperateQQ.UIN
	var group RefreshGroupStruct
	group.Gid = gid
	group.Bot = bot
	group.Uid = uid
	RefreshGroupChan <- group
	groupfunction := GroupFunctionModel.Api_find(gid)
	if len(groupfunction) < 1 {
		GroupFunctionModel.Api_insert(gid)
		groupfunction = GroupFunctionModel.Api_find(gid)
	}

	auto_retract := true
	if groupfunction["auto_retract"].(int64) == 0 {
		auto_retract = false
	}
	switch Type {
	//取消管理
	case 9:
		if uid == bot {
			if GroupMemberModel.Api_update_type(gid, uid, "member") {
				api.Sendgroupmsg(bot, gid, "Acfur-Off，权限已回收，将在2小时内退群", auto_retract)
			} else {
				api.Sendgroupmsg(bot, gid, "Acfur-Off，权限已回收，数据故障", auto_retract)
			}
		} else {
			if GroupMemberModel.Api_update_type(gid, uid, "member") {
				api.Sendgroupmsg(bot, gid, "管理员列表更新", auto_retract)
			} else {
				api.Sendgroupmsg(bot, gid, "管理员权限变动失败", auto_retract)
			}
		}
		break

	//设定管理
	case 10:
		if uid == bot {
			if GroupMemberModel.Api_update_type(gid, uid, "admin") {
				api.Sendgroupmsg(bot, gid, "Acfur-On，已获取权限，可使用acfurhelp查看功能", auto_retract)
			} else {
				api.Sendgroupmsg(bot, gid, "Acfur-On，已获取权限，数据故障，请使用acfur刷新人数来更新信息", auto_retract)
			}
		} else {
			if GroupMemberModel.Api_update_type(gid, uid, "admin") {
				api.Sendgroupmsg(bot, gid, "恭喜上位"+service.Serv_at(uid), auto_retract)
			} else {
				api.Sendgroupmsg(bot, gid, "恭喜上位,但是权限变动失败", auto_retract)
			}
		}
		break

	//T出某个人
	case 6:
		groupmsg := GroupMsgModel.Api_select(gid, uid, 10)
		last_msg := []string{}
		for _, data := range groupmsg {
			last_msg = append(last_msg, Calc.Any2String(data["text"]))
		}
		jsonmsg, _ := Jsong.Encode(last_msg)
		if GroupKickModel.Api_insert(bot, gid, uid, jsonmsg) {
			api.Sendgroupmsg(bot, gid, "群成员T出报告已经生成，请在APP中查看", auto_retract)
		} else {
			api.Sendgroupmsg(bot, gid, "群成员T出报告生成失败", auto_retract)
		}
		if groupfunction["kick_to_black"].(int64) == 1 {
			GroupBlackListModel.Api_insert(gid, uid, operator)
		}
		break

	//申请加群信息
	case 3:
		if groupfunction["auto_join"].(int64) == 1 {
			api.Setgroupaddrequest(bot, gid, uid, seq, api.Request_group_approve, api.Request_group_type_join, "")
		}
		//auto_verify := true
		//if groupfunction["auto_verify"].(int64) == 0 {
		//	auto_verify = false
		//}
		//auto_hold := true
		//if groupfunction["auto_hold"].(int64) == 0 {
		//	auto_hold = false
		//}
		break

	//退群信息
	case 5:
		if groupfunction["exit_alert"].(int64) == 1 {
			api.Sendgroupmsg(bot, gid, "成员-1", auto_retract)
		}
		if groupfunction["exit_to_black"].(int64) == 1 {
			GroupBlackListModel.Api_insert(gid, uid, operator)
		}
		break

	//加群消息
	case 2:
		if uid == bot {

		} else {
			if groupfunction["join_alert"].(int64) == 1 {
				api.Sendgroupmsg(bot, gid, "成员+1", auto_retract)
			}
			if groupfunction["auto_welcome"].(int64) == 1 {
				api.Sendgroupmsg(bot, gid, service.Serv_at(uid)+Calc.Any2String(groupfunction["welcome_word"]), auto_retract)
			}
		}
		break

	//私人加好友
	case 105:
		botinfo := BotModel.Api_find_byOwnerandBot(uid, bot)
		if len(botinfo) > 0 {
			api.Setfriendaddrequest(bot, uid, seq, api.Request_friend_approve)
			go func() {
				time.Sleep(5 * time.Second)
				Private.App_refresh_friend_list(bot)
			}()
		} else {
			api.Setfriendaddrequest(bot, uid, seq, api.Request_friend_denide)
		}
		break

	//机器人被T出
	case 30:
		break

	//机器人被邀请进群
	case 1:
		if len(BotGroupAllowModel.Api_find(bot, gid)) > 0 {
			api.Setgroupaddrequest(bot, gid, uid, seq, api.Request_group_approve, api.Request_group_type_invite, "")
		} else {
			api.Setgroupaddrequest(bot, gid, uid, seq, api.Request_group_denide, api.Request_group_type_invite, "不在群列表中")
		}
		break

	default:
		break
	}

}
