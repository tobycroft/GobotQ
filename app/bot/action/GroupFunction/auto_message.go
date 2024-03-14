package GroupFunction

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
)

func AutoMessage(self_id, group_id, user_id int64, message *MessageBuilder.IMessageBuilder, groupfunction map[string]any) {
	AutoRetract := false
	if Calc.Any2Int64(groupfunction["auto_retract"]) == 1 {
		AutoRetract = true
	}
	if Calc.Any2Int64(groupfunction["all_send_private"]) == 1 {
		go iapi.Api.SendPrivateMsg(self_id, user_id, group_id, message, AutoRetract)
	} else {
		go iapi.Api.SendGroupMsg(self_id, group_id, message, AutoRetract)
	}
}
