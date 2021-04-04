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
		switch data["type"].(string) {
		case "sep":
			timer := data["sep"].(int64)
			next_time := time.Now().Unix() + timer
			ass.Api_update_next_time(data["gid"], data["id"], next_time)
			break

		case "fix":
			timer := data["sep"].(int64)
			next_time := time.Now().Unix() + timer
			ass.Api_update_next_time(data["gid"], data["id"], next_time)
			ass.Api_update_active(data["gid"], data["id"], 0)
			break

		default:
			break
		}
		var gss api.GroupSendStruct
		if data["retract"].(int64) == 1 {
			gss.AutoRetract = true
		} else {
			gss.AutoRetract = false
		}
		group := GroupListModel.Api_find(data["gid"])
		if len(group) < 1 {
			return
		}
		ass.Api_dec_count(data["id"])
		db.Commit()
		gss.Fromqq = group["bot"]
		gss.Text = Calc.Any2String(data["msg"])
		gss.Togroup = data["group"]
		api.Group_send_chan <- gss
	}
}

func auto_send_sep() {

}

func auto_send_fix() {

}
