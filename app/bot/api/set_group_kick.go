package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

type KickGroupMemberRet struct {
	Ret string `json:"ret"`
}

func SetGroupKick(self_id, group_id, user_id interface{}, reject_add_request bool) bool {
	post := map[string]interface{}{
		"group_id":           group_id,
		"user_id":            user_id,
		"reject_add_request": reject_add_request,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_group_kick", nil, post, nil, nil)
	return true
}
