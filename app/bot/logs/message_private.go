package logs

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/message/group"
	"main.go/app/bot/model/PrivateMsgModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

func log_message_private() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroup) {
		var es group.EventStruct[group.GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, c.Payload)
		} else {
			pm := es.Json
			message, err := sonic.MarshalString(pm.Message)
			if err != nil {
				Log.Errs(err, pm.RawMessage)
			} else {
				PrivateMsgModel.Api_insert(pm.SelfId, pm.UserId, pm.MessageId, message, pm.RawMessage, pm.SubType, pm.Time)
			}
		}
	}
}
