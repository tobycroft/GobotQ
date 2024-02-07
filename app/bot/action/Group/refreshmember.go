package Group

import (
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func App_refreshmember(self_id, group_id int64) {
	var apm App_group_member
	apm.SelfId = self_id
	apm.GroupId = group_id
	Redis.PubSub{}.Publish_struct(types.RefreshGroupMembers, apm)
	//go api.Api{}.Sendgroupmsg(*bot, *gid, "群用户已经刷新", true)
}
