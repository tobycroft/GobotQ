package cron

import (
	"main.go/app/bot/apipost"
	"main.go/config/app_conf"
	"time"
)

func Retract() {
	go retract_private()
	retract_instant()
}

func retract_private() {
	for r := range apipost.Retract_chan {
		go func(retract apipost.Struct_Retract) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case apipost.Retract_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_instant() {
	for r := range apipost.Retract_instant {
		apipost.ApiPost{}.DeleteMsg(r.Self_id, r.MessageId)
	}
}
