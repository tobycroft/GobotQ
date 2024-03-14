package private

import (
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"strings"
)

func private_auto_reply(selfId, user_id, group_id int64, message string) {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(selfId)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text(auto_reply["value"].(string))
			iapi.Api.SendPrivateMsg(selfId, user_id, group_id, msg, true)
			break
		}
	}
}

func private_default_reply(selfId, user_id, group_id int64, message string) bool {
	default_reply := BotDefaultReplyModel.Api_select()
	for _, auto_reply := range default_reply {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text(auto_reply["value"].(string))
			iapi.Api.SendPrivateMsg(selfId, user_id, group_id, msg, false)
			return true
		}
	}
	return false
}
