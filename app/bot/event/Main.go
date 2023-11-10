package event

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/app/bot/model/LogsModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

type EventStruct struct {
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
}

func EventRouter(json string, remoteip string) {
	go LogsModel.Api_insert(json, "main", remoteip)
	var data EventStruct
	err := jsoniter.UnmarshalFromString(json, &data)
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
			fmt.Println(json)
			break

		default:
			LogRecvModel.Api_insert(json)
			break
		}
	}
}
