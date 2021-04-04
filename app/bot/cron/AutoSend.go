package cron

import (
	"main.go/app/bot/api"
	"main.go/app/v1/group/model/AutoSendModel"
	"main.go/tuuz"
	"time"
)

func Cron_auto_send() {
	datas := AutoSendModel.Api_select_next_time_up()
	for _, data := range datas {
		db := tuuz.Db()
		db.Begin()
		switch data["type"].(string) {
		case "sep":
			timer := data["sep"].(int64)
			next_time := time.Now().Unix() + timer
			var ass AutoSendModel.Interface
			ass.Db = db
			ass.Api_update_next_time(data["id"], next_time)
			var gss api.GroupSendStruct
			if data["retract"].(int64) == 1 {
				gss.AutoRetract = true
			} else {
				gss.AutoRetract = false
			}

			api.Group_send_chan <- gss
			break

		case "fix":

			break

		default:
			break
		}
		db.Commit()

	}
}

func auto_send_sep() {

}

func auto_send_fix() {

}
