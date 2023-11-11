package cron

import (
	apipost2 "main.go/app/bot/iapi/apipost"
	"main.go/config/app_conf"
	"time"
)

func Retract() {
	go retract_private()
	retract_instant()
}

func retract_private() {
	for r := range apipost2.Retract_chan {
		go func(retract apipost2.Struct_Retract) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case apipost2.Retract_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_instant() {
	for r := range apipost2.Retract_instant {
		apipost2.Api{}.DeleteMsg(r.Self_id, r.MessageId)
	}
}
