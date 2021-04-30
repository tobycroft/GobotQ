package event

import (
	"errors"
	"fmt"
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

func EventRouter(json string) {
	LogsModel.Api_insert(json, "main")
	var data EventStruct
	err := jsoniter.UnmarshalFromString(json, &data)
	if err != nil {
		LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
	} else {
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary

		switch data.PostType {
		case "message":
			message_type := data.MessageType
			switch message_type {
			case "private":
				var pm PM
				err = jsr.UnmarshalFromString(json, &pm)
				if err != nil {
					fmt.Println(err)
				} else {
					PrivateMsg(pm)
				}
				break

			case "group":
				var gm GM
				err = jsr.UnmarshalFromString(json, &gm)
				if err != nil {
					fmt.Println(err)
				} else {
					GroupMsg(gm)
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
			err = jsr.UnmarshalFromString(json, &notice)
			if err != nil {
				fmt.Println(err)
			} else {
				NoticeMsg(notice)
			}
			break

		case "request":
			var req Request
			err = jsr.UnmarshalFromString(json, &req)
			if err != nil {
				fmt.Println(err)
			} else {
				RequestMsg(req)
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
