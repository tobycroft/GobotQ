package event

import (
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
		Type := data["type"]
		if Type == nil {
			return
		}
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary

		switch Type {
		case "PrivateMsg":
			var pm PM
			jsr.UnmarshalFromString(json, &pm)
			PrivateMsg(pm)
			break

		case "GroupMsg":
			var gm GM
			jsr.UnmarshalFromString(json, &gm)
			GroupMsg(gm)
			break

		case "EventMsg":
			var em EM
			jsr.UnmarshalFromString(json, &em)
			EventMsg(em)
			break

		default:
			break
		}
		LogRecvModel.Api_insert(json)
	}
}
