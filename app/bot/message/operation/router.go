package operation

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"log"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotSettingModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/model/LogErrorModel"
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
			fmt.Println(err)
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
				logininfo := iapi.LoginInfoRet{}
				err := sonic.UnmarshalString(oe.json, &logininfo)
				if err != nil {
					fmt.Println(oe.json)
					break
				}
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
				iapi.Api.Getfriendlist(self_id)
				break

			case "get_friend_list":
				data := iapi.FriendListRet{}
				err := sonic.UnmarshalString(oe.json, &data)
				if err != nil {
					fmt.Println(err, oe.json)
					break
				}
				Private.App_refresh_friend_list_action(self_id, data.Data)
				fmt.Println("好友列表更新完毕：", oe.Echo.SelfId)
				iapi.Api.Getgrouplist(self_id)
				break

			case "get_group_list":
				data := iapi.GroupListRet{}
				err := sonic.UnmarshalString(oe.json, &data)
				if err != nil {
					fmt.Println(err, oe.json)
					break
				}
				Group.App_refresh_group_list_action(self_id, data.Data)
				fmt.Println("群列表更新完毕：", oe.Echo.SelfId)
				for _, datum := range data.Data {
					num := GroupMemberModel.Api_count_byGroupIdAndRole(datum.GroupId, nil)
					if num-datum.MemberNum != 0 {
						log.Println("需要更新的群：", self_id, datum.GroupId, num, datum.MemberNum, datum.MemberCount)
						Group.Chan_refresh_group_member <- Group.App_group_member{
							SelfId: self_id, GroupId: datum.GroupId,
						}
					}
				}
				break

			case "get_group_member_list":
				data := iapi.GroupMemberListRet{}
				err := sonic.UnmarshalString(oe.json, &data)
				if err != nil {
					fmt.Println(err, oe.json)
					break
				}
				Group.App_refresh_group_member_one_action(self_id, data.Data)
				fmt.Println("群成员更新完毕：", oe.Echo.SelfId, Calc.Any2Int64(oe.Echo.Extra))
				break

			case "send_private_msg", "send_group_msg":
				if oe.Echo.Extra.(bool) {
					data := iapi.MessageRet{}
					err := sonic.UnmarshalString(oe.json, &data)
					if err != nil {
						fmt.Println(err, oe.json)
						break
					}
					iapi.Retract_chan <- iapi.Struct_Retract{
						SelfId:    oe.Echo.SelfId,
						MessageId: data.Data.MessageId,
					}
					fmt.Println("发送消息，有撤回", oe.Echo.SelfId, data.Data.MessageId)
				} else {
					fmt.Println("发送消息，无撤回", oe.Echo.SelfId)

				}
				break

			default:
				fmt.Println(oe)
				break

			}

		}

	}

}
