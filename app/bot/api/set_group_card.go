package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

func SetGroupCard(self_id, group_id, user_id, card interface{}) bool {
	post := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"card":     card,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_group_card", nil, post, nil, nil)
	return true
}
