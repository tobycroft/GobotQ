package Private

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/config/app_default"
	"main.go/tuuz"
)

func App_bind_robot(self_id, user_id, group_id int64, message string) {
	if len(message) < 2 {
		msg := MessageBuilder.IMessageBuilder{}.New()
		msg.Text("请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	data := BotModel.Api_find(self_id)
	if len(data) > 0 {
		if Calc.Any2Int64(data["owner"]) != 0 {
			msg := MessageBuilder.IMessageBuilder{}.New()
			msg.Text("本机器人已经被绑定，如果需要清除绑定，请让号主解除本机器人的绑定")
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, true)
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
		if Calc.Any2String(data["secret"]) != message {
			db.Rollback()
			msg := MessageBuilder.IMessageBuilder{}.New()
			msg.Text("绑定密码不正确")
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
			return
		}
		if BotModel.Api_update_owner(self_id, user_id) {
			db.Commit()
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text("你已经成功绑定这个机器人咯！")
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		} else {
			db.Rollback()
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text("机器人绑定失败" + app_default.Default_error_alert)
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		}
	} else {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("未找到这个机器人，也许机器人的密码有错？")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, true)
	}
}

func App_unbind_bot(self_id int64, user_id, group_id int64, message string) {
	data := BotModel.Api_find(self_id)
	if len(data) < 1 {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("未找到当前机器人的信息，请稍后再试" + app_default.Default_error_alert)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	if Calc.Any2Int64(data["owner"]) != user_id {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("对不起您不是当前机器人的拥有人，请联系拥有人先行解绑")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, true)
		return
	}
	if BotModel.Api_update_owner(self_id, 0) {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("取消绑定成功")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	} else {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("取消绑定失败")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	}
}

func App_change_bot_secret(self_id int64, user_id, group_id int64, message string) {
	data := BotModel.Api_find(self_id)
	if len(data) < 1 {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("未找到当前机器人的信息，请稍后再试" + app_default.Default_error_alert)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	if len(message) < 2 {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("请使用\"acfur修改密码(+)密码\"来修改您机器人的绑定密码")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	if Calc.Any2Int64(data["owner"]) != int64(user_id) {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("对不起您不是当前机器人的拥有人，请联系拥有人先行解绑")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, true)
		return
	}
	if BotModel.Api_update_password(self_id, message) {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("修改机器人密码成功，机器人当前的密码为：" + message)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	} else {
		msg := MessageBuilder.IMessageBuilder{}.New().New().Text("修改机器人密码失败")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	}
}
