package cron

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/event"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"time"
)

func Refresh_group_chan() {
	for data := range event.RefreshGroupChan {
		group_check(data.Uid, data.Bot, data.Gid)
	}
}

func group_check(uid, bot, gid int) {
	groupinfo := GroupListModel.Api_find(bot, gid)
	if len(groupinfo) < 1 {
		Group.App_refresh_groupinfo(&bot, &gid)
	}
	userinfo := GroupMemberModel.Api_find(gid, uid)
	if len(userinfo) < 1 {
		Group.App_refreshmember(&bot, &gid)
	}
}

func Refresh_friend_list() {
	for {
		Private.App_refresh_friend_list_all()
		time.Sleep(3600 * time.Second)
	}
}
