package cron

import (
	"main.go/app/bot/iapi"
	"main.go/config/app_conf"
	"time"
)

func Retract() {
	go retract_private()
	retract_instant()
}

func retract_private() {
	for r := range iapi.Retract_chan {
		go func(retract iapi.Struct_Retract) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case iapi.Retract_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_instant() {
	for r := range iapi.Retract_instant {
		iapi.Api.DeleteMsg(r.Self_id, r.MessageId)
	}
}
