package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/config/app_default"
	"main.go/tuuz"
	"time"
)

func App_bind_robot(bot int, uid int, text string) {
	if len(text) < 2 {
		api.Sendprivatemsg(bot, uid, "请使用\"acfur绑定账号:密码\"来绑定您的机器人", false)
		return
	}
	data := BotModel.Api_find(bot)
	if len(data) > 0 {
		if data["owner"].(int64) != 0 {
			api.Sendprivatemsg(bot, uid, "本机器人已经被绑定，如果需要清除绑定，请让号主解除本机器人的绑定", true)
			return
		}
		db := tuuz.Db()
		db.Begin()
		var botreq BotRequestModel.Interface
		botreq.Db = db
		if !botreq.Api_delete(bot) {
			db.Rollback()
			return
		}
		if data["secret"] != text {
			api.Sendprivatemsg(bot, uid, "绑定密码不正确", false)
			return
		}
		if BotModel.Api_insert(bot, bot, "private", uid, data["secret"], data["password"], time.Now().Unix()+data["time"].(int64)) {
			db.Commit()
			api.Sendprivatemsg(bot, uid, "你已经成功绑定这个机器人咯！", false)
		} else {
			db.Rollback()
			api.Sendprivatemsg(bot, uid, "机器人绑定失败"+app_default.Default_error_alert, false)
		}
	} else {
		api.Sendprivatemsg(bot, uid, "未找到这个机器人，也许机器人的密码有错？", true)
	}
}
