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
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"time"
)

//var Private_send_chan = make(chan PrivateSendStruct, 20)

type MessageRet struct {
	Data    Message `json:"data"`
	Retcode int64   `json:"retcode"`
	Status  string  `json:"status"`
}

type Message struct {
	MessageId int64 `json:"message_id"`
}

func (api Post) SendPrivateMsg(Self_id, UserId, GroupId int64, Message string, AutoRetract bool) {
	var pss PrivateSendStruct
	pss.SelfId = Self_id
	pss.UserId = UserId
	pss.Message = Message
	pss.GroupId = GroupId
	pss.AutoRetract = AutoRetract
	pss.RetractTime = app_conf.Retract_time_duration

	Redis.PubSub{}.Publish_struct(types.SendPrivateChannel, pss)
}
func (api Ws) SendPrivateMsg(Self_id, UserId, GroupId int64, Message string, AutoRetract bool) {
	var pss PrivateSendStruct
	pss.SelfId = Self_id
	pss.UserId = UserId
	pss.Message = Message
	pss.GroupId = GroupId
	pss.AutoRetract = AutoRetract
	pss.RetractTime = app_conf.Retract_time_duration

	Redis.PubSub{}.Publish_struct(types.SendPrivateChannel, pss)
}

func (api Post) SendPrivateMsgWithTime(Self_id, UserId, GroupId int64, Message string, AutoRetract bool, Time time.Duration) {
	var pss PrivateSendStruct
	pss.SelfId = Self_id
	pss.UserId = UserId
	pss.Message = Message
	pss.GroupId = GroupId
	pss.AutoRetract = AutoRetract
	pss.RetractTime = Time

	Redis.PubSub{}.Publish_struct(types.SendPrivateChannel, pss)
}

func (api Ws) SendPrivateMsgWithTime(Self_id, UserId, GroupId int64, Message string, AutoRetract bool, Time time.Duration) {
	var pss PrivateSendStruct
	pss.SelfId = Self_id
	pss.UserId = UserId
	pss.Message = Message
	pss.GroupId = GroupId
	pss.AutoRetract = AutoRetract
	pss.RetractTime = Time

	Redis.PubSub{}.Publish_struct(types.SendPrivateChannel, pss)
}

type PrivateSendStruct struct {
	SelfId      int64         `json:"self_id"`
	GroupId     int64         `json:"group_id"`
	UserId      int64         `json:"user_id"`
	Message     string        `json:"message"`
	AutoRetract bool          `json:"auto_retract"`
	RetractTime time.Duration `json:"retract_time"`
}

func (api Post) Send_private() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.SendPrivateChannel) {
		var pss PrivateSendStruct
		err := sonic.UnmarshalString(c.Payload, &pss)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
		if Redis.CheckExists("SendCheck:" + pss.Message) {
			continue
		}
		Redis.String_set("SendCheck:"+pss.Message, true, app_conf.Retract_time_duration)
		pmr, err := api.sendprivatemsg(pss)
		if err != nil {

		} else {
			if pss.AutoRetract {
				rm := RetractMessage{
					SelfId:    pss.SelfId,
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
	for c := range ps.Subscribe(types.SendPrivateChannel) {
		var pss PrivateSendStruct
		err := sonic.UnmarshalString(c.Payload, &pss)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
		if Redis.CheckExists("SendCheck:" + pss.Message) {
			log.Println("SendCheck:" + pss.Message)
			continue
		}
		Redis.String_set("SendCheck:"+pss.Message, true, app_conf.Retract_time_duration)
		api.sendprivatemsg(pss)
	}
}

func (api Post) sendprivatemsg(pss PrivateSendStruct) (Message, error) {
	post := map[string]any{
		"user_id":     pss.UserId,
		"message":     pss.Message,
		"group_id":    pss.GroupId, //如果没加群不要使用群ID
		"auto_escape": false,
	}
	botinfo := BotModel.Api_find(pss.SelfId)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(pss.SelfId))
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

	botinfo := BotModel.Api_find(pss.SelfId)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(pss.SelfId))
		return Message{}, errors.New("botinfo_notfound")
	}
	LogSendModel.Api_insert(pss.SelfId, "private", 0, pss.Message)
	_, err := FriendListAction.App_find_friendList(pss.SelfId, pss.UserId)
	if err != nil {
		data, err := GroupMemberAction.App_find_groupMember(pss.SelfId, pss.UserId, nil)
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
			SelfId: Calc.Any2Int64(pss.SelfId),
			Extra:  pss.AutoRetract,
		},
	})
	if err != nil {
		return Message{}, err
	}
	conn, ok := ClientToConn.Load(pss.SelfId)
	if !ok {
		return Message{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return Message{}, err
}
