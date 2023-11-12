package cron

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/event"
	"main.go/app/bot/model/GroupMsgModel"
	"main.go/tuuz/Log"
)

func GroupMsgRecv() {
	for gm := range event.GroupMsgChan {
		message, err := sonic.MarshalString(gm.Message)
		if err != nil {
			Log.Errs(err, gm.RawMessage)
		} else {
			GroupMsgModel.Api_insert(gm.SelfId, gm.UserId, gm.GroupId, message, gm.RawMessage, gm.MessageId, gm.SubType)
		}

	}
}
