package iapi

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/url"
	"time"
)

var Group_send_chan = make(chan GroupSendStruct, 100)

func (api Post) Sendgroupmsg(Self_id, Group_id any, Message string, AutoRetract bool) {
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
func (api Ws) Sendgroupmsg(Self_id, Group_id any, Message string, AutoRetract bool) {
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
	Self_id     any
	Group_id    any
	Message     string
	AutoRetract bool
}

func (api Post) Send_group() {
	for gss := range Group_send_chan {
		if Redis.CheckExists("SendCheck:" + gss.Message) {
			continue
		}
		Redis.String_set("SendCheck:"+gss.Message, true, 110)
		gmr, err := api.sendgroupmsg(gss)
		if err != nil {

		} else {
			if gss.AutoRetract {
				if gmr.MessageId < 2 {
					fmt.Println("gmr.MessageId:", gmr.MessageId)
				}
				var r Struct_Retract
				r.Self_id = gss.Self_id
				r.MessageId = gmr.MessageId
				Retract_chan <- r
			}
		}
	}
}
func (api Ws) Send_group() {
	for gss := range Group_send_chan {
		if Redis.CheckExists("SendCheck:" + gss.Message) {
			continue
		}
		Redis.String_set("SendCheck:"+gss.Message, true, 110)
		gmr, err := api.sendgroupmsg(gss)
		if err != nil {

		} else {
			if gss.AutoRetract {
				if gmr.MessageId < 2 {
					fmt.Println("gmr.MessageId:", gmr.MessageId)
				}
				var r Struct_Retract
				r.Self_id = gss.Self_id
				r.MessageId = gmr.MessageId
				Retract_chan <- r
			}
		}
	}
}

func (api Post) sendgroupmsg(gss GroupSendStruct) (Message, error) {
	msg := url.QueryEscape(gss.Message)
	post := map[string]any{
		"group_id":    gss.Group_id,
		"message":     msg,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(gss.Self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(gss.Self_id))
		return Message{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/send_group_msg", nil, post, nil, nil).RetString()
	if err != nil {
		return Message{}, err
	}
	var gm MessageRet

	err = sonic.UnmarshalString(data, &gm)
	if err != nil {
		return Message{}, err
	}

	return gm.Data, nil
}
func (api Ws) sendgroupmsg(gss GroupSendStruct) (Message, error) {
	msg := url.QueryEscape(gss.Message)
	post := map[string]any{
		"group_id":    gss.Group_id,
		"message":     msg,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(gss.Self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(gss.Self_id))
		return Message{}, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "send_group_msg",
		Params: post,
		Echo: echo{
			Action: "send_group_msg",
			SelfId: Calc.Any2Int64(gss.Self_id),
		},
	})
	if err != nil {
		return Message{}, err
	}
	conn, ok := ClientToConn.Load(gss.Self_id)
	if !ok {
		return Message{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return Message{}, err
}
