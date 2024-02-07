package service

import (
	"main.go/app/bot/iapi"
)

func Not_admin(bot, gid, uid int64) {
	go iapi.Api.SendGroupMsg(bot, gid, "你不是本群的管理员，无法使用本功能"+Serv_at(uid), true)
}

func Not_owner(bot, gid, uid int64) {
	go iapi.Api.SendGroupMsg(bot, gid, "本功能仅限群主执行"+Serv_at(uid), true)
}
