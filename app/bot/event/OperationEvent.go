package event

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
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
		friend_list := []iapi.FriendList{}
		err := sonic.UnmarshalString(oe.json, &friend_list)
		if err != nil {
			fmt.Println(oe.json)
			return
		}
		Private.App_refresh_friend_list_action(self_id, friend_list)
		fmt.Println("好友列表更新完毕：", oe.Echo.SelfId)
		break

	default:
		fmt.Println(oe)
		break

	}

}
