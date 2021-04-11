package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
	"main.go/tuuz/Redis"
	"time"
)

var Private_send_chan = make(chan PrivateSendStruct, 20)

type MessageRet struct {
	Data    Message `json:"data"`
	Retcode int     `json:"retcode"`
	Status  string  `json:"status"`
}

type Message struct {
	MessageId int `json:"message_id"`
}

func Sendprivatemsg(Self_id, UserId, GroupId interface{}, Message string, AutoRetract bool) {
	var pss PrivateSendStruct
	pss.Self_id = Self_id
	pss.UserId = UserId
	pss.Message = Message
	pss.GroupId = GroupId
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
	GroupId     interface{}
	Message     string
	AutoRetract bool
}

func Send_private() {
	for pss := range Private_send_chan {
		if Redis.CheckExists("SendCheck:" + pss.Message) {
			continue
		}
		Redis.SetRaw("SendCheck:"+pss.Message, true, 110)
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

func sendprivatemsg(pss PrivateSendStruct) (Message, error) {
	post := map[string]interface{}{
		"user_id":     pss.UserId,
		"message":     pss.Message,
		"group_id":    pss.GroupId,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(pss.Self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(pss.Self_id))
		return Message{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/send_private_msg", nil, post, nil, nil)

	if err != nil {
		return Message{}, err
	}
	var pmr MessageRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &pmr)
	if err != nil {
		return Message{}, err
	}
	return pmr.Data, nil
}
