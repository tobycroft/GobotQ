package event

import (
	"fmt"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/tuuz/Redis"
	"net"
	"time"
)

type RequestMessage struct {
	remoteaddr  net.Addr
	json        string
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
	GroupId     int64  `json:"group_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	SelfID      int64  `json:"self_id"`
	SubType     string `json:"sub_type"`
	Time        int64  `json:"time"`
	UserID      int64  `json:"user_id"`
}

func (em RequestMessage) RequestMsg() {

	self_id := em.SelfID
	user_id := em.UserID
	group_id := em.GroupId
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
			iapi.Api.SetFriendAddRequest(self_id, flag, true, nil)
			go func() {
				time.Sleep(5 * time.Second)
				Private.App_refresh_friend_list(self_id)
			}()
		} else {
			go iapi.Api.SetFriendAddRequest(self_id, flag, false, "你不在机器人的允许列表中")
		}
		break

	case "group":
		switch sub_type {
		case "add":
			if groupfunction["auto_join"].(int64) == 1 {
				if len(GroupBlackListModel.Api_find(group_id, user_id)) > 0 {
					go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, false, "您在黑名单中请联系管理")
				} else {
					Redis.String_set("__request_comment__"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), comment, 86400*time.Second)
					go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, true, "")
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
				go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, true, "")
			} else {
				go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, false, "不在群列表中")
			}
			break

		default:
			fmt.Println("request no route sub_type", em)
			LogRecvModel.Api_insert(em.json)
			break

		}
		break

	default:
		fmt.Println("request no route", em)
		LogRecvModel.Api_insert(em.json)

		break

	}
}
