package api

import (
	"errors"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
)

type Retract struct {
	Self_id   interface{}
	MessageId interface{}
}

var Retract_chan = make(chan Retract, 20)
var Retract_chan_instant = make(chan Retract, 20)

func DeleteMsg(self_id, message_id interface{}) (bool, error) {
	post := map[string]interface{}{
		"message_id": message_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		return false, errors.New("botinfo_notfound")
	}
	Net.Post(botinfo["url"].(string)+"/delete_msg", nil, post, nil, nil)
	return true, nil
}
