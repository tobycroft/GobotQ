package iapi

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"log"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogSendModel"
	"main.go/config/app_conf"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Redis"
	"reflect"

	"main.go/tuuz/Log"
	"net/url"
	"time"
)

func (api Post) SenndGroupMsg(Self_id, Group_id int64, Message string, AutoRetract bool) {
	var gss GroupSendStruct
	gss.SelfId = Self_id
	gss.GroupId = Group_id
	gss.Message = Message
	gss.AutoRetract = AutoRetract
	gss.RetractTime = app_conf.Retract_time_duration

	Redis.PubSub{}.Publish_struct(types.SendGroupChannel, gss)
}
func (api Post) SenndGroupMsgWithTime(Self_id, Group_id int64, Message string, AutoRetract bool, Time time.Duration) {
	var gss GroupSendStruct
	gss.SelfId = Self_id
	gss.GroupId = Group_id
	gss.Message = Message
	gss.AutoRetract = AutoRetract
	gss.RetractTime = Time

	Redis.PubSub{}.Publish_struct(types.SendGroupChannel, gss)
}
func (api Ws) SenndGroupMsg(Self_id, Group_id int64, Message string, AutoRetract bool) {
	var gss GroupSendStruct
	gss.SelfId = Self_id
	gss.GroupId = Group_id
	gss.Message = Message
	gss.AutoRetract = AutoRetract
	gss.RetractTime = app_conf.Retract_time_duration

	Redis.PubSub{}.Publish_struct(types.SendGroupChannel, gss)
}
func (api Ws) SenndGroupMsgWithTime(Self_id, Group_id int64, Message string, AutoRetract bool, Time time.Duration) {
	var gss GroupSendStruct
	gss.SelfId = Self_id
	gss.GroupId = Group_id
	gss.Message = Message
	gss.AutoRetract = AutoRetract
	gss.RetractTime = Time

	Redis.PubSub{}.Publish_struct(types.SendGroupChannel, gss)
}

type GroupSendStruct struct {
	SelfId      int64         `json:"self_id"`
	GroupId     int64         `json:"group_id"`
	Message     string        `json:"message"`
	AutoRetract bool          `json:"auto_retract"`
	RetractTime time.Duration `json:"retract_time"`
}

func (api Post) Send_group() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.SendGroupChannel) {
		var gss GroupSendStruct
		err := sonic.UnmarshalString(c.Payload, &gss)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
		if Redis.CheckExists(types.SendCheck + gss.Message) {
			continue
		}
		Redis.String_set(types.SendCheck+gss.Message, true, app_conf.Retract_time_duration)
		gmr, err := api.sendgroupmsg(gss)
		if err != nil {

		} else {
			if gss.AutoRetract {
				if gmr.MessageId < 2 {
					fmt.Println("gmr.MessageId:", gmr.MessageId)
				}

				rm := RetractMessage{
					SelfId:    gss.SelfId,
					MessageId: gmr.MessageId,
					Time:      app_conf.Retract_time_duration,
				}
				ps.Publish_struct(types.RetractChannel, rm)

			}
		}
	}
}
func (api Ws) Send_group() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.SendGroupChannel) {
		var gss GroupSendStruct
		err := sonic.UnmarshalString(c.Payload, &gss)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
		if Redis.CheckExists(types.SendCheck + gss.Message) {
			continue
		}
		Redis.String_set(types.SendCheck+gss.Message, true, app_conf.Retract_time_duration)
		api.sendgroupmsg(gss)
	}
}

func (api Post) sendgroupmsg(gss GroupSendStruct) (Message, error) {
	msg := url.QueryEscape(gss.Message)
	post := map[string]any{
		"group_id":    gss.GroupId,
		"message":     msg,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(gss.SelfId)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(gss.SelfId))
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
	//msg := url.QueryEscape(gss.Message)
	post := map[string]any{
		"group_id":    gss.GroupId,
		"message":     gss.Message,
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(gss.SelfId)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(gss.SelfId))
		return Message{}, errors.New("botinfo_notfound")
	}
	LogSendModel.Api_insert(gss.SelfId, "private", 0, gss.Message)
	data, err := sonic.Marshal(sendStruct{
		Action: "send_group_msg",
		Params: post,
		Echo: echo{
			Action: "send_group_msg",
			SelfId: Calc.Any2Int64(gss.SelfId),
			Extra:  gss.AutoRetract,
		},
	})
	if err != nil {
		return Message{}, err
	}

	conn, ok := ClientToConn.Load(gss.SelfId)
	if !ok {
		log.Println(tuuz.FUNCTION_ALL(), "ClientNotFound", gss.SelfId, reflect.TypeOf(gss.SelfId))
		return Message{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return Message{}, err
}
