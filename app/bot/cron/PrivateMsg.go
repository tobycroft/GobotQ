package cron

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/event"
	"main.go/app/bot/model/PrivateMsgModel"
	"main.go/tuuz/Log"
)

func PrivateMsgRecv() {
	for pm := range event.PrivateMsgChan {
		message, err := sonic.MarshalString(pm.Message)
		if err != nil {
			Log.Errs(err, pm.RawMessage)
		} else {
			PrivateMsgModel.Api_insert(pm.SelfId, pm.UserId, pm.MessageId, message, pm.RawMessage, pm.SubType, pm.Time)
		}
	}
}
