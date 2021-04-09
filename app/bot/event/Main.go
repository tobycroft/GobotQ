package event

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/app/bot/model/LogsModel"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
)

func EventRouter(json string) {
	//save the data in the first place
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

			case "GroupMsg":
				var gm GM
				err = jsr.UnmarshalFromString(json, &gm)
				if err != nil {
					fmt.Println(err)
				} else {
					GroupMsg(gm)
				}
				break

			case "EventMsg":
				var em EM
				err = jsr.UnmarshalFromString(json, &em)
				if err != nil {
					fmt.Println(err)
				} else {
					EventMsg(em)
				}
				break
			}
			break

		case "meta_event":
			break

		default:
			LogRecvModel.Api_insert(json)
			break
		}
	}
}
