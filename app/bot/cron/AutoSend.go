package cron

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/v1/group/model/AutoSendModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"time"
)

func Cron_auto_send() {
	for {
		auto_send()
		time.Sleep(60 * time.Second)
	}
}

func auto_send() {
	datas := AutoSendModel.Api_select_next_time_up()
	for _, data := range datas {
		db := tuuz.Db()
		db.Begin()
		var ass AutoSendModel.Interface
		ass.Db = db

		timer := data["sep"].(int64)
		next_time := time.Now().Unix() + timer
		ass.Api_update_next_time(data["gid"], data["id"], next_time)

		switch data["type"].(string) {
		case "sep":
			//如果是采用间隔模式，则需要测算下次时间，并count-1
			break

		case "fix":
			//如果采用一次性模式，则直接关闭这个定时
			ass.Api_update_active(data["gid"], data["id"], 0)
			break

		default:
			break
		}
		ass.Api_dec_count(data["id"])
		db.Commit()
		//发送部分
		auto_retract := false
		if data["retract"].(int64) == 1 {
			auto_retract = true
		}
		group := GroupListModel.Api_find(data["gid"])
		if len(group) < 1 {
			return
		}
		api.Sendgroupmsg(group["bot"], data["gid"], Calc.Any2String(data["msg"]), auto_retract)
	}
}