package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
	"time"
)

var Private_send_chan = make(chan PrivateSendStruct, 20)

type MessageRet struct {
	MessageId int `json:"message_id"`
}

func Sendprivatemsg(Self_id, UserId interface{}, Message string, AutoRetract bool) {
	var pss PrivateSendStruct
	pss.Self_id = Self_id
	pss.UserId = UserId
	pss.Message = Message
	pss.AutoRetract = AutoRetract

	select {
	case Private_send_chan <- pss:

	case <-time.After(5 * time.Second):
		return
	}
}

type PrivateSendStruct struct {
	Self_id     interface{}
	UserId      interface{}
	Message     interface{}
	AutoRetract bool
}

func Send_private() {
	for pss := range Private_send_chan {
		pmr, err := sendprivatemsg(pss)
		if err != nil {

		} else {
			if pss.AutoRetract {
				var r Struct_Retract
				r.Self_id = pss.Self_id
				r.MessageId = pmr.MessageId
				Retract_chan <- r
			}
		}
	}
}

func sendprivatemsg(pss PrivateSendStruct) (MessageRet, error) {
	post := map[string]interface{}{
		"user_id":     pss.UserId,
		"message":     pss.Message,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(pss.Self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(pss.Self_id))
		return MessageRet{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/send_private_msg", nil, post, nil, nil)
	if err != nil {
		return MessageRet{}, err
	}
	var pmr MessageRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &pmr)
	if err != nil {
		return MessageRet{}, err
	}
	return pmr, nil
}
