package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

func SetGroupAddRequestRet(self_id, flag, sub_type interface{}, approve bool, reason string) bool {
	post := map[string]interface{}{
		"flag":     flag,
		"sub_type": sub_type,
		"type":     sub_type,
		"approve":  approve,
		"reason":   reason,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_group_add_request", nil, post, nil, nil)
	return true
}
