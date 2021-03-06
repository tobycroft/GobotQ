package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/config/app_default"
	"main.go/tuuz"
)

func App_bind_robot(bot int, uid, gid int, text string) {
	if len(text) < 2 {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
		} else {
			api.Sendprivatemsg(bot, uid, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
		}
		return
	}
	data := BotModel.Api_find(bot)
	if len(data) > 0 {
		if data["owner"].(int64) != 0 {
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, "本机器人已经被绑定，如果需要清除绑定，请让号主解除本机器人的绑定", true)
			} else {
				api.Sendprivatemsg(bot, uid, "本机器人已经被绑定，如果需要清除绑定，请让号主解除本机器人的绑定", true)
			}
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
		if data["secret"].(string) != text {
			db.Rollback()
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, "绑定密码不正确", false)
			} else {
				api.Sendprivatemsg(bot, uid, "绑定密码不正确", false)
			}
			return
		}
		if BotModel.Api_update_owner(bot, uid) {
			db.Commit()
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, "你已经成功绑定这个机器人咯！", false)
			} else {
				api.Sendprivatemsg(bot, uid, "你已经成功绑定这个机器人咯！", false)
			}
		} else {
			db.Rollback()
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, "机器人绑定失败"+app_default.Default_error_alert, false)
			} else {
				api.Sendprivatemsg(bot, uid, "机器人绑定失败"+app_default.Default_error_alert, false)
			}
		}
	} else {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "未找到这个机器人，也许机器人的密码有错？", true)
		} else {
			api.Sendprivatemsg(bot, uid, "未找到这个机器人，也许机器人的密码有错？", true)
		}
	}
}

func App_unbind_bot(bot int, uid, gid int, text string) {
	data := BotModel.Api_find(bot)
	if len(data) < 1 {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		} else {
			api.Sendprivatemsg(bot, uid, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		}
		return
	}
	if data["owner"].(int64) != int64(uid) {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		} else {
			api.Sendprivatemsg(bot, uid, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		}
		return
	}
	if BotModel.Api_update_owner(bot, 0) {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "取消绑定成功", false)
		} else {
			api.Sendprivatemsg(bot, uid, "取消绑定成功", false)
		}
	} else {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "取消绑定失败", false)
		} else {
			api.Sendprivatemsg(bot, uid, "取消绑定失败", false)
		}
	}
}

func App_change_bot_secret(bot int, uid, gid int, text string) {
	data := BotModel.Api_find(bot)
	if len(data) < 1 {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		} else {
			api.Sendprivatemsg(bot, uid, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		}
		return
	}
	if len(text) < 2 {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "请使用\"acfur修改密码(+)密码\"来修改您机器人的绑定密码", false)
		} else {
			api.Sendprivatemsg(bot, uid, "请使用\"acfur修改密码(+)密码\"来修改您机器人的绑定密码", false)
		}
		return
	}
	if data["owner"].(int64) != int64(uid) {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		} else {
			api.Sendprivatemsg(bot, uid, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		}
		return
	}
	if BotModel.Api_update_password(bot, text) {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "修改机器人密码成功，机器人当前的密码为："+text, false)
		} else {
			api.Sendprivatemsg(bot, uid, "修改机器人密码成功，机器人当前的密码为："+text, false)
		}
	} else {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, "修改机器人密码失败", false)
		} else {
			api.Sendprivatemsg(bot, uid, "修改机器人密码失败", false)
		}
	}
}
