package event

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/app/bot/model/LogsModel"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
)

func EventRouter(json string) {
	LogsModel.Api_insert(json, "main")
	data, err := Jsong.JObject(json)
	if err != nil {
		LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
	} else {
		post_type := data["post_type"]
		if post_type == nil {
			fmt.Println("typenot on", data)
			return
		}
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary

		switch post_type {

		case "message":
			message_type := data["message_type"].(string)
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
