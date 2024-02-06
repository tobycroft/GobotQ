package cron

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/iapi"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"time"
)

func Retract() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.RetractChannel) {
		var es iapi.RetractMessage
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
		go retract_and_wait_for_exec(es.SelfId, es.MessageId, es.Time)
	}
}

func retract_and_wait_for_exec(self_id, message_id any, duration time.Duration) {
	time.Sleep(duration)
	iapi.Api.DeleteMsg(self_id, message_id)
}

//func retract_private() {
//	for r := range iapi.Retract_chan {
//		fmt.Println("retract_private", r)
//		go func(retract iapi.Struct_Retract) {
//			time.Sleep(app_conf.Retract_time_second * time.Second)
//			select {
//			case iapi.Retract_instant <- retract:
//
//			case <-time.After(5 * time.Second):
//				return
//			}
//		}(r)
//	}
//}

//func retract_instant() {
//	for r := range iapi.Retract_instant {
//		iapi.Api.DeleteMsg(r.SelfId, r.MessageId)
//	}
//}
