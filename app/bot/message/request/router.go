package request

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotSettingModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"net/netip"
	"time"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageRequest) {
		var es EventStruct[RequestMessage]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			em := es.Json
			bot := BotModel.Api_find(em.SelfId)
			if len(bot) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				return
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if bot["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), em.SelfId)
				return
			}
			self_id := em.SelfId
			request_type := em.RequestType

			switch request_type {
			case "friend":
				var esf EventStruct[requestFriend]
				rr := esf.Json
				err := sonic.UnmarshalString(c.Payload, &esf)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c)
					return
				}
				user_id := rr.UserId
				flag := rr.Flag
				botinfo := BotModel.Api_find_byOwnerandBot(user_id, self_id)
				if len(botinfo) > 0 {
					iapi.Api.SetFriendAddRequest(self_id, flag, true, nil)
					go func() {
						time.Sleep(5 * time.Second)
						Private.App_refresh_friend_list(self_id)
					}()
				} else {
					iapi.Api.SetFriendAddRequest(self_id, flag, false, "你不在机器人的允许列表中")
				}
				break

			case "group":
				var esf EventStruct[requestGroup]
				rr := esf.Json
				err := sonic.UnmarshalString(c.Payload, &esf)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c)
					return
				}
				user_id := rr.UserId
				flag := rr.Flag
				group_id := rr.GroupId
				sub_type := rr.SubType
				groupfunction := GroupFunctionModel.Api_find(group_id)
				if len(groupfunction) < 1 {
					GroupFunctionModel.Api_insert(group_id)
					groupfunction = GroupFunctionModel.Api_find(group_id)
				}
				switch sub_type {
				case "add":
					if groupfunction["auto_join"].(int64) == 1 {
						if len(GroupBlackListModel.Api_find(group_id, user_id)) > 0 {
							go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, false, "您在黑名单中请联系管理")
						} else {
							Redis.String_set("__request_comment__"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), "", 86400*time.Second)
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
					bot_setting := BotSettingModel.Api_find(self_id)
					if len(BotGroupAllowModel.Api_find(self_id, group_id)) > 0 || bot_setting["add_group"] == 1 {
						go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, true, "")
					} else {
						go iapi.Api.SetGroupAddRequestRet(self_id, flag, sub_type, false, "不在群列表中")
					}
					break

				default:
					fmt.Println("request no route sub_type", em)
					LogRecvModel.Api_insert(c.Payload)
					break

				}
				break

			default:
				fmt.Println("request no route", em)
				LogRecvModel.Api_insert(c.Payload)
				break

			}
		}

	}

}
