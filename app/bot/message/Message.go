package message

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/iapi"
	"main.go/app/bot/message/group"
	"main.go/app/bot/message/index"
	"main.go/app/bot/message/meta_event"
	"main.go/app/bot/message/notice"
	"main.go/app/bot/message/operation"
	"main.go/app/bot/message/private"
	"main.go/app/bot/message/request"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Redis"
)

func MainRouter() {
	go index.Router()
	go meta_event.Router()
	go notice.Router()

	go request.Router()

	go private.Router()
	go group.Router()

	go operation.Router()

	for c := range Net.WsServer_ReadChannel {
		if c.Status {
			var es index.EventStruct
			err := sonic.Unmarshal(c.Message, &es)
			if err != nil {
				LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
				continue
			}
			mp := map[string]any{}
			err = sonic.Unmarshal(c.Message, &mp)
			if err != nil {
				LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
				continue
			}
			es.Json = mp
			es.RemoteAddr = c.Conn.RemoteAddr().String()
			iapi.ClientToConn.Store(es.SelfId, c.Conn)
			iapi.ConnToClient.Store(c.Conn, es.SelfId)

			Redis.PubSub{}.Publish_struct(types.MessageEvent, es)
		} else {
			fmt.Println(c.Conn.RemoteAddr(), "断开链接")
			client, ok := iapi.ConnToClient.Load(c.Conn)
			if ok {
				iapi.ClientToConn.Delete(client)
			}
		}
	}
}
