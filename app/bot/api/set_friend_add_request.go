package api

import (
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

func SetFriendAddRequest(self_id, flag, approve, remark interface{}) bool {
	post := map[string]interface{}{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false
	}
	Net.Post(botinfo["url"].(string)+"/set_friend_add_request", nil, post, nil, nil)
	return true
}
