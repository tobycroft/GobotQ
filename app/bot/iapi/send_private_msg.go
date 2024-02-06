package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"log"
	"main.go/app/bot/action/FriendListAction"
	"main.go/app/bot/action/GroupMemberAction"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogSendModel"
	"main.go/config/app_conf"
	"main.go/config/types"
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

func (api Post) Sendprivatemsg(Self_id, UserId, GroupId any, Message string, AutoRetract bool) {
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
func (api Ws) Sendprivatemsg(Self_id, UserId, GroupId any, Message string, AutoRetract bool) {
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

func (api Post) Send_private() {
	ps := Redis.PubSub{}
	for pss := range Private_send_chan {
		if Redis.CheckExists("SendCheck:" + pss.Message) {
			continue
		}
		Redis.String_set("SendCheck:"+pss.Message, true, 110*time.Second)
		pmr, err := api.sendprivatemsg(pss)
		if err != nil {

		} else {
			if pss.AutoRetract {
				rm := RetractMessage{
					SelfId:    pss.Self_id,
					MessageId: pmr.MessageId,
					Time:      app_conf.Retract_time_duration,
				}
				ps.Publish_struct(types.RetractChannel, rm)
			}
		}
	}
}
func (api Ws) Send_private() {
	ps := Redis.PubSub{}
	for pss := range Private_send_chan {
		if Redis.CheckExists("SendCheck:" + pss.Message) {
			log.Println("SendCheck:" + pss.Message)
			continue
		}
		Redis.String_set("SendCheck:"+pss.Message, true, 110*time.Second)
		pmr, err := api.sendprivatemsg(pss)
		if err != nil {

		} else {
			if pss.AutoRetract {
				rm := RetractMessage{
					SelfId:    pss.Self_id,
					MessageId: pmr.MessageId,
					Time:      app_conf.Retract_time_duration,
				}
				ps.Publish_struct(types.RetractChannel, rm)
			}
		}
	}
}

func (api Post) sendprivatemsg(pss PrivateSendStruct) (Message, error) {
	post := map[string]any{
		"user_id":     pss.UserId,
		"message":     pss.Message,
		"group_id":    pss.GroupId, //如果没加群不要使用群ID
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
func (api Ws) sendprivatemsg(pss PrivateSendStruct) (Message, error) {
	post := map[string]any{
		"user_id":     pss.UserId,
		"message":     pss.Message,
		"auto_escape": false,
	}

	botinfo := BotModel.Api_find(pss.Self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(pss.Self_id))
		return Message{}, errors.New("botinfo_notfound")
	}
	LogSendModel.Api_insert(pss.Self_id, "private", 0, pss.Message)
	_, err := FriendListAction.App_find_friendList(pss.Self_id, pss.UserId)
	if err != nil {
		data, err := GroupMemberAction.App_find_groupMember(pss.Self_id, pss.UserId, nil)
		if err != nil {
			return Message{}, err
		} else {
			post["group_id"] = data.GroupId
		}
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "send_private_msg",
		Params: post,
		Echo: echo{
			Action: "send_private_msg",
			SelfId: Calc.Any2Int64(pss.Self_id),
			Extra:  pss.AutoRetract,
		},
	})
	if err != nil {
		return Message{}, err
	}
	conn, ok := ClientToConn.Load(pss.Self_id)
	if !ok {
		return Message{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return Message{}, err
}
