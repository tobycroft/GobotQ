package cron

import (
	"main.go/app/bot/event"
	"main.go/app/bot/model/GroupMsgModel"
)

func GroupMsgRecv() {
	for gm := range event.GroupMsgChan {
		GroupMsgModel.Api_insert(gm.SelfID, gm.UserID, gm.GroupId, gm.Message, gm.RawMessage, gm.MessageID, gm.SubType)
	}
}
