package logs

import (
	"github.com/bytedance/sonic"
	event "main.go/app/bot/message"
	"main.go/app/bot/model/LogsModel"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net"
)

func LogsInit() {
	go log_message_all()
}

type eventStruct struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        string
	Remoteaddr  net.Addr
}

func log_message_all() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(event.MessageEvent) {
		var es eventStruct
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, c.Payload)
		} else {
			if es.PostType != "meta_event" {
				LogsModel.Api_insert(es.Json, event.MessageEvent, es.Remoteaddr)
			}
		}
	}
}
