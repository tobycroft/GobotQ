package event

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/iapi"
	"main.go/app/bot/message/index"
	"main.go/app/bot/message/meta_event"
	"main.go/app/bot/message/notice"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/tuuz"
	"main.go/tuuz/Redis"
)

func MainRouter() {
	go index.Router()
	go meta_event.Router()
	go notice.Router()

	for c := range Net.WsServer_ReadChannel {
		if c.Status {
			var es index.EventStruct
			err := sonic.Unmarshal(c.Message, &es)
			if err != nil {
				go LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
				continue
			}
			if c.Status {
				iapi.ClientToConn.Store(es.SelfId, c.Conn)
				iapi.ConnToClient.Store(c.Conn, es.SelfId)
			}
			Redis.PubSub{}.Publish_struct(MessageEvent, es)
		} else {
			fmt.Println(c.Conn.RemoteAddr(), "断开链接")
			client, ok := iapi.ConnToClient.Load(c.Conn)
			if ok {
				iapi.ClientToConn.Delete(client)
			}
		}
	}
}
