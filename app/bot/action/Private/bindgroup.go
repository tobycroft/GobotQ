package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotGroupAllowModel"
)

func App_bind_group(self_id, user_id int64, message string) {
	if len(message) < 2 {
		api.Sendprivatemsg(self_id, user_id, "请使用\"acfur绑定群群号\"，来绑定新群", false)
		return
	}
	if len(BotGroupAllowModel.Api_find(self_id, message)) > 0 {
		api.Sendprivatemsg(self_id, user_id, "群号已经被绑定："+message, false)
		return
	}
	if BotGroupAllowModel.Api_insert(self_id, message) {
		api.Sendprivatemsg(self_id, user_id, "绑定群已经增加："+message, false)
	} else {
		api.Sendprivatemsg(self_id, user_id, "绑定群增加失败："+message, false)
	}
}

func App_unbind_group(self_id int64, user_id int64, message string) {
	if len(message) < 2 {
		api.Sendprivatemsg(self_id, user_id, "请使用\"acfur绑定群群号\"，来绑定新群", false)
		return
	}
	if BotGroupAllowModel.Api_insert(self_id, message) {
		api.Sendprivatemsg(self_id, user_id, "绑定群已经删除："+message, false)
	} else {
		api.Sendprivatemsg(self_id, user_id, "绑定群删除失败："+message, false)
	}
}
