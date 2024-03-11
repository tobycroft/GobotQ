package operation

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"log"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotSettingModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/config/app_conf"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/netip"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageOperation) {
		var es EventStruct[OperationEvent]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, c.Payload)
		} else {
			oe := es.Json
			bot := BotModel.Api_find(oe.Echo.SelfId)
			if len(bot) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				continue
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if bot["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), oe.Echo.SelfId)
				continue
			}
			self_id := oe.Echo.SelfId
			switch oe.Echo.Action {
			case "get_login_info":
				var es EventStruct[iapi.LoginInfoRet]
				err := sonic.UnmarshalString(c.Payload, &es)
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					logininfo := es.Json
					user_id := logininfo.Data.UserID
					nickname := logininfo.Data.Nickname
					if !BotModel.Api_update_cname(user_id, nickname) {
						Log.Crrs(errors.New("机器人用户名无法更新"), tuuz.FUNCTION_ALL())
					} else {
						fmt.Println("机器人更新完毕：", logininfo.Data)
					}
					if len(BotSettingModel.Api_find(self_id)) < 1 {
						BotSettingModel.Api_insert(self_id, 0, 0)
					}
					iapi.Api.GetFriendList(self_id)
				}
				break

			case "get_friend_list":
				var es EventStruct[iapi.FriendListRet]
				err := sonic.UnmarshalString(c.Payload, &es)
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					data := es.Json
					Private.App_refresh_friend_list_action(self_id, data.Data)
					fmt.Println("好友列表更新完毕：", oe.Echo.SelfId)
					iapi.Api.GetGroupList(self_id)
				}
				break

			case "get_group_list":
				var es EventStruct[iapi.GroupListRet]
				err := sonic.UnmarshalString(c.Payload, &es)
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					data := es.Json
					GroupFunction.App_refresh_group_list_action(self_id, data.Data)
					fmt.Println("群列表更新完毕：", oe.Echo.SelfId)
					for _, datum := range data.Data {
						num := GroupMemberModel.Api_count_byGroupIdAndRole(datum.GroupId, nil)
						if num-datum.MemberNum != 0 {
							log.Println("需要更新的群：", self_id, datum.GroupId, num, datum.MemberNum, datum.MemberCount)
							iapi.Api.GetGroupMemberList(self_id, datum.GroupId)
							Redis.PubSub{}.Publish_struct(types.RefreshGroupMembers, GroupFunction.App_group_member{
								SelfId:  self_id,
								GroupId: datum.GroupId,
							})
						}
					}
				}
				break

			case "get_group_member_list":
				var es EventStruct[iapi.GroupMemberListRet]
				err := sonic.UnmarshalString(c.Payload, &es)
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					data := es.Json
					GroupFunction.App_refresh_group_member_one_action(self_id, data.Data)
					fmt.Println("群成员更新完毕：", oe.Echo.SelfId, Calc.Any2Int64(oe.Echo.Extra))
				}
				break

			case "send_private_msg", "send_group_msg":
				if oe.Echo.Extra.(bool) {
					var es EventStruct[iapi.MessageRet]
					err := sonic.UnmarshalString(c.Payload, &es)
					if err != nil {
						Log.Errs(err, tuuz.FUNCTION_ALL())
					} else {
						if es.Json.Retcode != 0 {
							Log.Crrs(errors.New(es.Json.Message), tuuz.FUNCTION_ALL())
							fmt.Println("发送消息，有错误", es.Json)
							break
						}
						data := es.Json
						var rm iapi.RetractMessage
						rm.MessageId = data.Data.MessageId
						rm.SelfId = oe.Echo.SelfId
						rm.Time = app_conf.Retract_time_duration
						ps.Publish_struct(types.RetractChannel, rm)
						fmt.Println("发送消息，有撤回", oe.Echo.SelfId, data.Data.MessageId)
					}
				} else {
					fmt.Println("发送消息，无撤回", oe.Echo.SelfId)
				}
				break

			case "delete_msg":
				var es EventStruct[iapi.RetractWsRetStruct]
				err := sonic.UnmarshalString(c.Payload, &es)
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					if es.Json.Retcode == 0 {
						fmt.Println("消息撤回完成：", oe.Echo.SelfId, Calc.Any2Int64(oe.Echo.Extra))
					} else {
						fmt.Println("消息撤回失败：", oe.Echo.SelfId, Calc.Any2Int64(es.Json.Message))
					}
				}
				break

			default:
				fmt.Println("event-notfound:", c.Payload)
				LogRecvModel.Api_insert(c.Payload)
				break

			}

		}

	}

}
