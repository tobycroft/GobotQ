package Private

import (
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotGroupAllowModel"
)

func App_bind_group(self_id, user_id, group_id int64, message string) {
	if len(message) < 2 {
		msg := MessageBuilder.IMessageBuilder{}.Text("请使用\"acfur绑定群群号\"，来绑定新群")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	if len(BotGroupAllowModel.Api_find(self_id, message)) > 0 {
		msg := MessageBuilder.IMessageBuilder{}.Text("群号已经被绑定：" + message)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	if BotGroupAllowModel.Api_insert(self_id, message) {
		msg := MessageBuilder.IMessageBuilder{}.Text("绑定群已经增加：" + message)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	} else {
		msg := MessageBuilder.IMessageBuilder{}.Text("绑定群增加失败：" + message)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	}
}

func App_unbind_group(self_id int64, user_id, group_id int64, message string) {
	if len(message) < 2 {
		msg := MessageBuilder.IMessageBuilder{}.Text("请使用\"acfur绑定群群号\"，来绑定新群")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		return
	}
	if BotGroupAllowModel.Api_insert(self_id, message) {
		msg := MessageBuilder.IMessageBuilder{}.Text("绑定群已经删除：" + message)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	} else {
		msg := MessageBuilder.IMessageBuilder{}.Text("绑定群删除失败：" + message)
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	}
}
