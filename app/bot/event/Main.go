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

func Speed_Limit() {
	private_speed_limit = map[int]interface{}{}
	group_speed_limit = map[int]interface{}{}
}

func EventRouter(json string) {
	//save the data in the first place
	LogsModel.Api_insert(json, "main")
	data, err := Jsong.JObject(json)
	if err != nil {
		LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
	} else {
		Type := data["Type"]
		if Type == nil {
			fmt.Println("typenot on", data)
			return
		}
		fmt.Println(data)
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary

		switch Type {
		case "PrivateMsg":
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

		default:
			LogRecvModel.Api_insert(json)
			break
		}
	}
}
