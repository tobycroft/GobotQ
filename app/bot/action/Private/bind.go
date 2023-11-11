package Private

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/config/app_default"
	"main.go/tuuz"
)

func App_bind_robot(self_id, user_id, group_id int64, message string) {
	if len(message) < 2 {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
		return
	}
	data := BotModel.Api_find(self_id)
	if len(data) > 0 {
		if data["owner"].(int64) != 0 {
			iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "本机器人已经被绑定，如果需要清除绑定，请让号主解除本机器人的绑定", true)
			return
		}
		db := tuuz.Db()
		db.Begin()
		var botreq BotRequestModel.Interface
		botreq.Db = db
		if !botreq.Api_delete(self_id) {
			db.Rollback()
			return
		}
		if data["secret"].(string) != message {
			db.Rollback()
			iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "绑定密码不正确", false)
			return
		}
		if BotModel.Api_update_owner(self_id, user_id) {
			db.Commit()
			iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "你已经成功绑定这个机器人咯！", false)
		} else {
			db.Rollback()
			iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "机器人绑定失败"+app_default.Default_error_alert, false)
		}
	} else {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "未找到这个机器人，也许机器人的密码有错？", true)
	}
}

func App_unbind_bot(self_id int64, user_id, group_id int64, message string) {
	data := BotModel.Api_find(self_id)
	if len(data) < 1 {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		return
	}
	if data["owner"].(int64) != int64(user_id) {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		return
	}
	if BotModel.Api_update_owner(self_id, 0) {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "取消绑定成功", false)
	} else {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "取消绑定失败", false)
	}
}

func App_change_bot_secret(self_id int64, user_id, group_id int64, message string) {
	data := BotModel.Api_find(self_id)
	if len(data) < 1 {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "未找到当前机器人的信息，请稍后再试"+app_default.Default_error_alert, false)
		return
	}
	if len(message) < 2 {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "请使用\"acfur修改密码(+)密码\"来修改您机器人的绑定密码", false)
		return
	}
	if data["owner"].(int64) != int64(user_id) {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "对不起您不是当前机器人的拥有人，请联系拥有人先行解绑", true)
		return
	}
	if BotModel.Api_update_password(self_id, message) {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "修改机器人密码成功，机器人当前的密码为："+message, false)
	} else {
		iapi.Post{}.Sendprivatemsg(self_id, user_id, group_id, "修改机器人密码失败", false)
	}
}
