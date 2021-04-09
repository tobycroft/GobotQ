package cron

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/event"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

type Struct_retract struct {
	Self_id   interface{}
	MessageId interface{}
}

func Retract() {
	go retract_private()
	go retract_group()
	retract_instant()
}

func retract_private() {
	for r := range event.Retract_chan_private {
		go func(retract event.Struct_retract) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case event.Retract_chan_private_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_group() {
	for r := range event.Retract_chan_group {
		go func(retract event.Struct_retract) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case event.Retract_chan_group_instant <- retract:

			case <-time.After(5 * time.Second):
				Log.Errs(errors.New("retract_group失败"), tuuz.FUNCTION_ALL())
				return
			}
		}(r)
	}
}

func retract_instant() {
	for r := range event.Retract_chan_private_instant {
		api.DeleteMsg(r.Self_id, r.MessageId)
	}
}
