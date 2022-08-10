package event

import (
	"fmt"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/tuuz/Redis"
	"time"
)

type Request struct {
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
	GroupID     int    `json:"group_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	SelfID      int    `json:"self_id"`
	SubType     string `json:"sub_type"`
	Time        int    `json:"time"`
	UserID      int    `json:"user_id"`
}

func RequestMsg(em Request, remoteip string) {

	self_id := em.SelfID
	user_id := em.UserID
	group_id := em.GroupID
	request_type := em.RequestType
	sub_type := em.SubType
	flag := em.Flag
	comment := em.Comment

	groupfunction := GroupFunctionModel.Api_find(group_id)
	if len(groupfunction) < 1 {
		GroupFunctionModel.Api_insert(group_id)
		groupfunction = GroupFunctionModel.Api_find(group_id)
	}

	switch request_type {
	case "friend":
		botinfo := BotModel.Api_find_byOwnerandBot(user_id, self_id)
		if len(botinfo) > 0 {
			api.SetFriendAddRequest(self_id, flag, true, nil)
			go func() {
				time.Sleep(5 * time.Second)
				Private.App_refresh_friend_list(self_id)
			}()
		} else {
			go api.SetFriendAddRequest(self_id, flag, false, "你不在机器人的允许列表中")
		}
		break

	case "group":
		switch sub_type {
		case "add":
			if groupfunction["auto_join"].(int64) == 1 {
				if len(GroupBlackListModel.Api_find(group_id, user_id)) > 0 {
					go api.SetGroupAddRequestRet(self_id, flag, sub_type, false, "您在黑名单中请联系管理")
				} else {
					Redis.String_set("__request_comment__"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), comment, 86400)
					go api.SetGroupAddRequestRet(self_id, flag, sub_type, true, "")
				}
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

		case "invite":
			if len(BotGroupAllowModel.Api_find(self_id, group_id)) > 0 {
				go api.SetGroupAddRequestRet(self_id, flag, sub_type, true, "")
			} else {
				go api.SetGroupAddRequestRet(self_id, flag, sub_type, false, "不在群列表中")
			}
			break

		default:
			fmt.Println("request no route sub_type", em)
			break

		}
		break

	default:
		fmt.Println("request no route", em)
		break

	}
}
