package event

import (
	"fmt"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupKickModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/model/GroupMsgModel"
	"main.go/app/bot/service"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
)

type Notice struct {
	Duration   int    `json:"duration"`
	GroupID    int64  `json:"group_id"`
	NoticeType string `json:"notice_type"`
	OperatorID int64  `json:"operator_id"`
	PostType   string `json:"post_type"`
	SelfID     int64  `json:"self_id"`
	SubType    string `json:"sub_type"`
	Time       int64  `json:"time"`
	UserID     int64  `json:"user_id"`
}

func NoticeMsg(em Notice) {
	self_id := em.SelfID
	user_id := em.UserID
	group_id := em.GroupID
	notice_type := em.NoticeType
	sub_type := em.SubType
	operator_id := em.OperatorID

	var group RefreshGroupStruct
	group.GroupId = group_id
	group.SelfId = self_id
	group.UserId = user_id
	RefreshGroupChan <- group

	groupfunction := GroupFunctionModel.Api_find(group_id)
	if len(groupfunction) < 1 {
		GroupFunctionModel.Api_insert(group_id)
		groupfunction = GroupFunctionModel.Api_find(group_id)
	}

	auto_retract := true
	if groupfunction["auto_retract"].(int64) == 0 {
		auto_retract = false
	}

	fmt.Println(em)
	switch notice_type {
	//取消管理
	case "group_admin":
		switch sub_type {
		case "set":
			if user_id == self_id {
				if GroupMemberModel.Api_update_type(group_id, user_id, "admin") {
					api.Sendgroupmsg(self_id, group_id, "Acfur-On，已获取权限，可使用acfurhelp查看功能", auto_retract)
				} else {
					api.Sendgroupmsg(self_id, group_id, "Acfur-On，已获取权限，数据故障，请使用acfur刷新人数来更新信息", auto_retract)
				}
			} else {
				if GroupMemberModel.Api_update_type(group_id, user_id, "admin") {
					api.Sendgroupmsg(self_id, group_id, "恭喜上位"+service.Serv_at(user_id), auto_retract)
				} else {
					api.Sendgroupmsg(self_id, group_id, "恭喜上位,但是权限变动失败", auto_retract)
				}
			}

			break

		case "unset":
			if user_id == self_id {
				if GroupMemberModel.Api_update_type(group_id, user_id, "member") {
					api.Sendgroupmsg(self_id, group_id, "Acfur-Off，权限已回收，将在2小时内退群", auto_retract)
				} else {
					api.Sendgroupmsg(self_id, group_id, "Acfur-Off，权限已回收，数据故障", auto_retract)
				}
			} else {
				if GroupMemberModel.Api_update_type(group_id, user_id, "member") {
					api.Sendgroupmsg(self_id, group_id, "管理员列表更新", auto_retract)
				} else {
					api.Sendgroupmsg(self_id, group_id, "管理员权限变动失败", auto_retract)
				}
			}
			break
		}
		break

	case "group_increase":
		if user_id == self_id {

		} else {
			if groupfunction["join_alert"].(int64) == 1 {
				api.Sendgroupmsg(self_id, group_id, "成员+1", auto_retract)
			}
			if groupfunction["auto_welcome"].(int64) == 1 {
				api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+Calc.Any2String(groupfunction["welcome_word"]), auto_retract)
			}
		}
		break

	case "group_decrease":
		switch sub_type {
		case "leave":
			if groupfunction["exit_alert"].(int64) == 1 {
				api.Sendgroupmsg(self_id, group_id, "成员-1", auto_retract)
			}
			if groupfunction["exit_to_black"].(int64) == 1 {
				GroupBlackListModel.Api_insert(group_id, user_id, operator_id)
			}
			break

		case "kick":
			groupmsg := GroupMsgModel.Api_select(group_id, user_id, 10)
			last_msg := []string{}
			for _, data := range groupmsg {
				last_msg = append(last_msg, Calc.Any2String(data["text"]))
			}
			jsonmsg, _ := Jsong.Encode(last_msg)
			if GroupKickModel.Api_insert(self_id, group_id, user_id, jsonmsg) {
				api.Sendgroupmsg(self_id, group_id, "群成员T出报告已经生成，请在APP中查看", auto_retract)
			} else {
				api.Sendgroupmsg(self_id, group_id, "群成员T出报告生成失败", auto_retract)
			}
			if groupfunction["kick_to_black"].(int64) == 1 {
				GroupBlackListModel.Api_insert(group_id, user_id, operator_id)
			}
			break

		case "kick_me":
			break

		default:
			fmt.Println("notice no route sub_type", em)
			break
		}
		break

	case "group_ban":
		switch sub_type {
		case "ban":

			break

		case "lift_ban":
			if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {
				GroupBanPermenentModel.Api_delete(group_id, user_id)
				api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+"你已经脱离永久小黑屋了", auto_retract)
			}
			break
		}
		break

	case "friend_add":

		break

	case "friend_recall":
		break

	case "notice":
		break

	default:
		fmt.Println("notice no route", em)
		break
	}

}
