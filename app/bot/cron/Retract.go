package cron

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
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
		fmt.Println(c.Payload)
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
		go retract_and_wait_for_exec(es.SelfId, es.MessageId, es.Time)
	}
}

func retract_and_wait_for_exec(self_id, message_id int64, duration time.Duration) {
	if duration.Seconds() > 0 {
		time.Sleep(duration)
	}
	fmt.Println("开始撤回:", Calc.Any2String(self_id), Calc.Any2String(message_id), duration)
	iapi.Api.DeleteMsg(self_id, message_id)
}
