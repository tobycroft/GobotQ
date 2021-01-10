package cron

import (
	"main.go/app/bot/event"
	"main.go/app/bot/model/GroupMsgModel"
)

func GroupMsgRecv() {
	for gm := range event.GroupMsgChan {
		GroupMsgModel.Api_insert(gm.LogonQQ, gm.FromQQ.UIN, gm.FromGroup.GIN, gm.Msg.Text, gm.Msg.Req, gm.Msg.Random, gm.File.ID, gm.File.MD5,
			gm.File.Name, gm.File.Size)
	}
}
