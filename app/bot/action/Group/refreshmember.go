package Group

func App_refreshmember(self_id, group_id *int64) {
	var apm App_group_member
	apm.SelfId = *self_id
	apm.GroupId = *group_id
	Chan_refresh_group_member <- apm
	//api.Sendgroupmsg(*bot, *gid, "群用户已经刷新", true)
}
