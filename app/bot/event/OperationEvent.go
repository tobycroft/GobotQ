package event

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
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
			fmt.Println("机器人更新：", logininfo.Data)
		}
		break

	default:
		fmt.Println(oe)
		break

	}

}
