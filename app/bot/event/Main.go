package event

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/app/bot/model/LogsModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func EventListener() {
	for c := range Net.WsServer_ReadChannel {
		EventRouter(string(c.Message), c.Conn.RemoteAddr().String())
	}
}

type EventStruct struct {
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
}

func EventRouter(json string, remoteip string) {
	go LogsModel.Api_insert(json, "main", remoteip)
	var data EventStruct
	err := sonic.UnmarshalString(json, &data)
	if err != nil {
		LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
	} else {

		switch data.PostType {
		case "message":
			message_type := data.MessageType
			switch message_type {
			case "private":
				var pm PM
				err = sonic.UnmarshalString(json, &pm)
				if err != nil {
					fmt.Println(err)
				} else {
					PrivateMsg(pm, remoteip)
				}
				break

			case "group":
				var gm GM
				err = sonic.UnmarshalString(json, &gm)
				if err != nil {
					fmt.Println(err)
				} else {
					GroupMsg(gm, remoteip)
				}
				break

			default:
				Log.Crrs(errors.New("undefine route"), json)
				break
			}
			break

		case "notice":
			//fmt.Println(json)
			var notice Notice
			err = sonic.UnmarshalString(json, &notice)
			if err != nil {
				fmt.Println(err)
			} else {
				NoticeMsg(notice, remoteip)
			}
			break

		case "request":
			var req Request
			err = sonic.UnmarshalString(json, &req)
			if err != nil {
				fmt.Println(err)
			} else {
				RequestMsg(req, remoteip)
			}
			break

		case "meta_event":
			var aaa MetaEventStruct
			err = sonic.UnmarshalString(json, &aaa)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("HeartBeat:", aaa.SelfId, aaa.Status.QqStatus)
			}
			break

		default:
			LogRecvModel.Api_insert(json)
			break
		}
	}
}

type MetaEventStruct struct {
	Time          int    `json:"time"`
	SelfId        int    `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
	Status        struct {
		Self struct {
			Platform string `json:"platform"`
			UserId   int    `json:"user_id"`
		} `json:"self"`
		Online   bool   `json:"online"`
		Good     bool   `json:"good"`
		QqStatus string `json:"qq.status"`
	} `json:"status"`
	Interval int `json:"interval"`
}
