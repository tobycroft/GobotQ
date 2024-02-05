package index

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageEvent) {
		var es EventStruct[string]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			LogErrorModel.Api_insert(err.Error(), c.Payload)
			continue
		}

		switch es.PostType {
		case "message":
			message_type := es.MessageType
			switch message_type {
			case "private":
				ps.Publish(types.MessagePrivate, c)
				break

			case "group":
				ps.Publish(types.MessageGroup, c)
				break

			default:
				Log.Crrs(errors.New("undefine route"), es.Json)
				break
			}
			break

		case "notice":
			//fmt.Println(es.PostType, message)
			ps.Publish(types.MessageNotice, c)
			break

		case "request":
			//fmt.Println(es.PostType, message)
			ps.Publish(types.MessageRequest, c)
			break

		case "meta_event":
			//trigger the
			ps.Publish(types.MessageMetaEvent, c)
			break

		default:
			ps.Publish(types.MessageOperation, c)

			fmt.Println("event-notfound:", es.Json)
			LogRecvModel.Api_insert(es.Json)
		}

	}
}
