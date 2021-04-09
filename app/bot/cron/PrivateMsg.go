package cron

import (
	"main.go/app/bot/event"
	"main.go/app/bot/model/PrivateMsgModel"
)

func PrivateMsgRecv() {
	for pm := range event.PrivateMsgChan {
		PrivateMsgModel.Api_insert(pm.SelfID, pm.UserID, pm.MessageID, pm.Message, pm.RawMessage, pm.SubType, pm.Time)
	}
}
