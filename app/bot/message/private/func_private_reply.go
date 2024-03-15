package private

import (
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/extend/TTS"
	"strings"
)

func private_auto_reply(selfId, user_id, group_id int64, message string) bool {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(selfId)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			rec, err := TTS.Audio{}.New().Huihui(auto_reply["value"].(string))
			if err == nil {
				iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), false)
				return true
			}
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text(auto_reply["value"].(string))
			iapi.Api.SendPrivateMsg(selfId, user_id, group_id, msg, true)
			return true
		}
	}
	return false
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
			rec, err := TTS.Audio{}.New().Huihui(auto_reply["value"].(string))
			if err == nil {
				iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), false)
				continue
			}
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text(auto_reply["value"].(string))
			iapi.Api.SendPrivateMsg(selfId, user_id, group_id, msg, false)
			return true
		}
	}
	return false
}
