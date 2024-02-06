package index

import (
	"errors"
	"github.com/bytedance/sonic"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageEvent) {
		var es EventStruct
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
				ps.Publish(types.MessagePrivate, c.Payload)
				break

			case "group":
				ps.Publish(types.MessageGroup, c.Payload)
				break

			default:
				Log.Crrs(errors.New("undefine route"), c.Payload)
				break
			}
			break

		case "notice":
			//fmt.Println(es.PostType, message)
			ps.Publish(types.MessageNotice, c.Payload)
			break

		case "request":
			//fmt.Println(es.PostType, message)
			ps.Publish(types.MessageRequest, c.Payload)
			break

		case "meta_event":
			//trigger the
			ps.Publish(types.MessageMetaEvent, c.Payload)
			break

		default:
			ps.Publish(types.MessageOperation, c.Payload)

		}

	}
}
