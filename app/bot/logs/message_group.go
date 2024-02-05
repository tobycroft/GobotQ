package logs

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/message/group"
	"main.go/app/bot/model/GroupMsgModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

func log_message_group() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroup) {
		var es group.EventStruct[group.GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, c.Payload)
		} else {
			gm := es.Json
			message, err := sonic.MarshalString(gm.Message)
			if err != nil {
				Log.Errs(err, gm.RawMessage)
			} else {
				GroupMsgModel.Api_insert(gm.SelfId, gm.UserId, gm.GroupId, message, gm.RawMessage, gm.MessageId, gm.SubType)
			}
		}
	}
}
