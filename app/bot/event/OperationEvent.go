package event

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
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"net"
)

type OperationEvent struct {
	remoteaddr net.Addr
	json       string
	Echo       struct {
		Action string `json:"action"`
		SelfId int64  `json:"self_id"`
		Extra  any    `json:"extra"`
	} `json:"echo"`
}

func (oe OperationEvent) OperationRouter() {
	self_id := oe.Echo.SelfId
	switch oe.Echo.Action {
	case "get_login_info":
		logininfo := iapi.LoginInfoRet{}
		err := sonic.UnmarshalString(oe.json, &logininfo)
		if err != nil {
			fmt.Println(oe.json)
			return
		}
		user_id := logininfo.Data.UserID
		nickname := logininfo.Data.Nickname
		if !BotModel.Api_update_cname(user_id, nickname) {
			Log.Crrs(errors.New("机器人用户名无法更新"), tuuz.FUNCTION_ALL())
		} else {
			fmt.Println("机器人更新完毕：", logininfo.Data)
		}
		iapi.Api.Getfriendlist(self_id)
		break

	case "get_friend_list":
		data := iapi.FriendListRet{}
		err := sonic.UnmarshalString(oe.json, &data)
		if err != nil {
			fmt.Println(err, oe.json)
			return
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
			return
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
			return
		}
		Group.App_refresh_group_member_one_action(self_id, data.Data)
		fmt.Println("群成员更新完毕：", oe.Echo.SelfId, Calc.Any2Int64(oe.Echo.Extra))
		break

	default:
		fmt.Println(oe)
		break

	}

}