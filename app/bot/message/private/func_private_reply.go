package private

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"strings"
)

func private_auto_reply(selfId int64, message string) (string, bool) {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(selfId)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			return Calc.Any2String(auto_reply["value"]), true
		}
	}
	return "", false
}

func private_default_reply(message string) (string, bool) {
	default_reply := BotDefaultReplyModel.Api_select()
	for _, auto_reply := range default_reply {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			return Calc.Any2String(auto_reply["value"]), true
		}
	}
	return "", false
}
