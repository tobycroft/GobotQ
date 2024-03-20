package service

import (
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
)

func Not_admin(bot, gid, uid int64) {
	msg := MessageBuilder.IMessageBuilder{}.New().Text("你不是本群的管理员，无法使用本功能").At(uid)
	go iapi.Api.SendGroupMsg(bot, gid, msg, true)
}

func Not_owner(bot, gid, uid int64) {
	msg := MessageBuilder.IMessageBuilder{}.New().Text("本功能仅限群主执行").At(uid)
	go iapi.Api.SendGroupMsg(bot, gid, msg, true)
}
