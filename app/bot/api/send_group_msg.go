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

var Group_send_chan = make(chan GroupSendStruct, 20)

func Sendgroupmsg(Self_id, Group_id interface{}, Message string, AutoRetract bool) {
	var gss GroupSendStruct
	gss.Self_id = Self_id
	gss.Group_id = Group_id
	gss.Message = Message
	gss.AutoRetract = AutoRetract

	select {
	case Group_send_chan <- gss:

	case <-time.After(5 * time.Second):
		return
	}
}

type GroupSendStruct struct {
	Self_id     interface{}
	Group_id    interface{}
	Message     string
	AutoRetract bool
}

func Send_group() {
	for gss := range Group_send_chan {
		gmr, err := sendgroupmsg(gss)
		if err != nil {

		} else {
			if gss.AutoRetract {
				var r Struct_Retract
				r.Self_id = gss.Self_id
				r.MessageId = gmr.MessageId
				Retract_chan <- r
			}
		}
	}
}

func sendgroupmsg(gss GroupSendStruct) (MessageRet, error) {
	post := map[string]interface{}{
		"group_id":    gss.Group_id,
		"message":     gss.Message,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(gss.Self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(gss.Self_id))
		return MessageRet{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/sendgroupmsg", nil, post, nil, nil)
	if err != nil {
		return MessageRet{}, err
	}
	var gm MessageRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return MessageRet{}, err
	}

	return gm, nil
}
