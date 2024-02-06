package logs

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/message/index"
	"main.go/app/bot/model/LogsModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

func LogsInit() {
	go log_message_all()
	go log_message_group()
	go log_message_private()
}

func log_message_all() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageEvent) {
		var es index.EventStruct
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, c.Payload)
		} else {
			if es.PostType != "meta_event" {
				json, err := sonic.MarshalString(es.Json)
				if err == nil {
					LogsModel.Api_insert(json, types.MessageEvent, es.RemoteAddr)
				}
			}
		}
	}
}
