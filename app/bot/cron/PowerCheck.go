package cron

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupMemberModel"
	"time"
)

func PowerCheck() {
	for {
		time.Sleep(2 * time.Hour)
		go power_check()
	}
}

func power_check() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		self_id := bot["self_id"]
		groups := GroupMemberModel.Api_select_groupBy_groupId(self_id)
		for _, group := range groups {
			group_id := group["group_id"]
			role := Group.BotPowerRefresh(group_id, self_id)
			if role == "member" {
				go api.Sendgroupmsg(self_id, group_id, "额，如果以后有需要管理，可以再叫我来啊？", false)
				api.SetGroupLeave(self_id, group_id)
				GroupMemberModel.Api_delete_byGid(self_id, group_id)
			} else if role == "owner" {
				gms := GroupMemberModel.Api_select_admin(self_id, group_id)
				for _, gm := range gms {
					api.SetGroupAdmin(self_id, group_id, gm["user_id"], false)
				}
				api.SetGroupWholeBan(self_id, group_id, true)
				api.SetGroupLeave(self_id, group_id)
			}
		}
	}
}
