package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotGroupAllowModel"
)

func App_bind_group(bot int, uid int, text string) {
	if len(text) < 2 {
		api.Sendprivatemsg(bot, uid, "请使用\"acfur绑定群群号\"，来绑定新群", false)
		return
	}
	if len(BotGroupAllowModel.Api_find(bot, text)) > 0 {
		api.Sendprivatemsg(bot, uid, "群号已经被绑定："+text, false)
		return
	}
	if BotGroupAllowModel.Api_insert(bot, text) {
		api.Sendprivatemsg(bot, uid, "绑定群已经增加："+text, false)
	} else {
		api.Sendprivatemsg(bot, uid, "绑定群增加失败："+text, false)
	}
}

func App_unbind_group(bot int, uid int, text string) {
	if len(text) < 2 {
		api.Sendprivatemsg(bot, uid, "请使用\"acfur绑定群群号\"，来绑定新群", false)
		return
	}
	if BotGroupAllowModel.Api_insert(bot, text) {
		api.Sendprivatemsg(bot, uid, "绑定群已经删除："+text, false)
	} else {
		api.Sendprivatemsg(bot, uid, "绑定群删除失败："+text, false)
	}
}
