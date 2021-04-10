package cron

import (
	"main.go/app/bot/api"
	"main.go/config/app_conf"
	"time"
)

func Retract() {
	go retract_private()
	retract_instant()
}

func retract_private() {
	for r := range api.Retract_chan {
		go func(retract api.Struct_Retract) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case api.Retract_chan_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_instant() {
	for r := range api.Retract_chan_instant {
		api.DeleteMsg(r.Self_id, r.MessageId)
	}
}
