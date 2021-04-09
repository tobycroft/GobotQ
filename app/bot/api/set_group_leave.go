package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

func SetGroupLeave(self_id, group_id interface{}) bool {
	post := map[string]interface{}{
		"group_id":   group_id,
		"is_dismiss": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_group_leave", nil, post, nil, nil)
	return true
}
