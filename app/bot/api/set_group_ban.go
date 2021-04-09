package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

type MuteGroupmeMberRet struct {
	Ret string `json:"ret"`
}

func SetGroupBan(self_id, group_id, user_id interface{}, duration float64) bool {
	post := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"duration": duration,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_group_ban", nil, post, nil, nil)
	return true
}
