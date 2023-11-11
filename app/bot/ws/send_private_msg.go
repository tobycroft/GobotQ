package api

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"time"
)

var Private_send_chan = make(chan PrivateSendStruct, 20)

type MessageRet struct {
	Data    Message `json:"data"`
	Retcode int64   `json:"retcode"`
	Status  string  `json:"status"`
}

type Message struct {
	MessageId int64 `json:"message_id"`
}

func (ws Ws) Sendprivatemsg(Self_id, UserId, GroupId any, Message string, AutoRetract bool) {
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
	Self_id     any
	UserId      any
	GroupId     any
	Message     string
	AutoRetract bool
}

func Send_private() {
	for pss := range Private_send_chan {
		if Redis.CheckExists("SendCheck:" + pss.Message) {
			continue
		}
		Redis.String_set("SendCheck:"+pss.Message, true, 110)
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
	post := map[string]any{
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
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/send_private_msg", nil, post, nil, nil).RetString()

	if err != nil {
		return Message{}, err
	}
	var pmr MessageRet

	err = sonic.UnmarshalString(data, &pmr)
	if err != nil {
		return Message{}, err
	}
	return pmr.Data, nil
}
