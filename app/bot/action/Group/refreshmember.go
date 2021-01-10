package Group

import "main.go/app/bot/api"

func App_refreshmember(bot *int, gid *int, uid *int) {
	var apm App_group_member
	apm.Bot = *bot
	apm.Gid = *gid
	Chan_refresh_group_member <- apm
	api.Sendgroupmsg(*bot, *gid, "群用户已经刷新")
}
