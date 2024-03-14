package GroupFunction

import (
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
)

func App_PalWorld(self_id, group_id, user_id int64, message MessageBuilder.IMessageBuilder, groupfunction map[string]any) {
	AutoRetract := false
	if groupfunction["auto_retract"].(int64) == 1 {
		AutoRetract = true
	}
	if groupfunction["all_send_private"].(int64) == 1 {
		go iapi.Api.SendPrivateMsg(self_id, user_id, group_id, message, AutoRetract)
	} else {
		go iapi.Api.SendGroupMsg(self_id, group_id, message, AutoRetract)
	}

}
