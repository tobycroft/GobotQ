package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/tuuz"
	"strings"
	"time"
)

func App_bind_robot(bot int, uid int, text string) {
	strs := strings.Split(text, ":")
	if len(strs) < 2 {
		api.Sendprivatemsg(bot, uid, "请使用\"acfur绑定账号:密码\"来绑定您的机器人")
	}
	data := BotRequestModel.Api_find(bot, text)
	if len(data) > 0 {
		db := tuuz.Db()
		db.Begin()
		var botreq BotRequestModel.Interface
		botreq.Db = db
		if !botreq.Api_delete(bot) {
			db.Rollback()
			return
		}
		if data["secret"] {

		}
		if BotModel.Api_insert(bot, bot, "private", uid, data["secret"], data["password"], time.Now().Unix()+data["time"].(int64)) {
			db.Commit()
			api.Sendprivatemsg(bot, uid, "你已经成功绑定这个机器人咯！", false)
		} else {
			db.Rollback()
			api.Sendprivatemsg(bot, uid)
		}
	} else {
		api.Sendprivatemsg(bot, uid, "未找到这个机器人", true)
	}
}
