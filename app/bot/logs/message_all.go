package logs

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/model/LogsModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net"
)

func LogsInit() {
	go log_message_all()
	go log_message_group()
	go log_message_private()
}

type eventStruct struct {
	SelfId      int64    `json:"self_id"`
	MessageType string   `json:"message_type"`
	PostType    string   `json:"post_type"`
	Json        string   `json:"json"`
	RemoteAddr  net.Addr `json:"remote_addr"`
}

func log_message_all() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageEvent) {
		var es eventStruct
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, c.Payload)
		} else {
			if es.PostType != "meta_event" {
				LogsModel.Api_insert(es.Json, types.MessageEvent, es.RemoteAddr)
			}
		}
	}
}
