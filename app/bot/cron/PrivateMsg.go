package cron

import (
	"main.go/app/bot/event"
	"main.go/app/bot/model/PrivateMsgModel"
)

func PrivateMsgRecv() {
	for pm := range event.PrivateMsgChan {
		PrivateMsgModel.Api_insert(pm.LogonQQ, pm.FromQQ.UIN, pm.Msg.Text, pm.Msg.Req, pm.Msg.Seq, pm.Msg.Type, pm.Msg.SubType, pm.File.ID,
			pm.File.MD5, pm.File.Name, pm.File.Size)
	}
}
