package service

import "main.go/app/bot/api"

func Not_admin(bot, gid, uid interface{}) {
	api.Sendgroupmsg(bot, gid, "你不是本群的管理员，无法使用本功能"+Serv_at(uid), true)
}
