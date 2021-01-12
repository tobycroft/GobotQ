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
		api.Sendprivatemsg(bot, uid, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
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

func App_unbind_bot(bot int, uid int, text string) {
	data := BotModel.Api_find(bot)
	if len(data) < 1 {
		api.Sendprivatemsg(bot, uid, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		return
	}
	if data["owner"] != uid {
		api.Sendprivatemsg(bot, uid, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		return
	}
	if len(text) < 2 {
		api.Sendprivatemsg(bot, uid, "请使用\"acfur解除绑定机器人(+)密码\"来绑定您的机器人", false)
		return
	}
	if BotModel.Api_update_owner(bot, 0) {
		api.Sendprivatemsg(bot, uid, "取消绑定成功", false)
	} else {
		api.Sendprivatemsg(bot, uid, "取消绑定失败", false)
	}
}

func Api_change_bot_password(bot int, uid int, text string) {
	data := BotModel.Api_find(bot)
	if len(data) < 1 {
		api.Sendprivatemsg(bot, uid, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		return
	}
	if len(text) < 2 {
		api.Sendprivatemsg(bot, uid, "请使用\"acfur修改密码(+)密码\"来修改您机器人的绑定密码", false)
		return
	}
	if data["owner"] != uid {
		api.Sendprivatemsg(bot, uid, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		return
	}
	if data["secret"] != text {
		api.Sendprivatemsg(bot, uid, "机器人密码错误，请重新输入", true)
		return
	}

	if BotModel.Api_update_password(bot, text) {
		api.Sendprivatemsg(bot, uid, "修改机器人密码成功", false)
	} else {
		api.Sendprivatemsg(bot, uid, "修改机器人密码失败", false)
	}
}
