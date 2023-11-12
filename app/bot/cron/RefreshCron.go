package cron

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/event"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"

	"main.go/tuuz/Redis"
	"time"
)

func Refresh_group_chan() {
	for data := range event.RefreshGroupChan {
		group_check(data.SelfId, data.UserId, data.GroupId)
	}
}

func group_check(self_id, user_id, group_id int64) {
	groupinfo := GroupListModel.Api_find(group_id)
	if len(groupinfo) < 1 {
		Group.App_refresh_groupinfo(self_id, group_id)
	} else {
		Redis.String_set("__groupinfo__"+Calc.Int642String(group_id)+"_"+Calc.Int642String(user_id), groupinfo, 60*time.Second)
	}
	userinfo := GroupMemberModel.Api_find(group_id, user_id)
	if len(userinfo) < 1 {
		Group.App_refreshmember(self_id, group_id)
	} else {
		Redis.String_set("__userinfo__"+Calc.Int642String(group_id)+"_"+Calc.Int642String(user_id), groupinfo, 60*time.Second)
	}
}

func Refresh_friend_list() {
	for {
		Private.App_refresh_friend_list_all()
		time.Sleep(3600 * time.Second)
	}
}
