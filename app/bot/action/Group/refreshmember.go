package Group

func App_refreshmember(bot *int, gid *int) {
	var apm App_group_member
	apm.Bot = *bot
	apm.Gid = *gid
	Chan_refresh_group_member <- apm
	//api.Sendgroupmsg(*bot, *gid, "群用户已经刷新", true)
}
