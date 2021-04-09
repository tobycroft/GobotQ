package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

type SetGroupWholeMuteRet struct {
	Ret string `json:"ret"`
}

func SetGroupWholeBan(self_id, group_id interface{}, enable bool) bool {
	post := map[string]interface{}{
		"self_id":  self_id,
		"group_id": group_id,
		"enable":   enable,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_group_whole_ban", nil, post, nil, nil)
	return true
}
