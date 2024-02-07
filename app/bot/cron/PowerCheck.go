package cron

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
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
		self_id := Calc.Any2Int64(bot["self_id"])
		groups := GroupMemberModel.Api_select_groupBy_groupId(self_id)
		for _, group := range groups {
			group_id := Calc.Any2Int64(group["group_id"])
			role := Group.BotPowerRefresh(group_id, self_id)
			if role == "member" {
				go iapi.Api.SendGroupMsg(self_id, group_id, "额，如果以后有需要管理，可以再叫我来啊？", false)
				iapi.Api.SetGroupLeave(self_id, group_id)
				GroupMemberModel.Api_delete_byGid(self_id, group_id)
			} else if role == "owner" {
				gms := GroupMemberModel.Api_select_admin(self_id, group_id)
				for _, gm := range gms {
					iapi.Api.SetGroupAdmin(self_id, group_id, Calc.Any2Int64(gm["user_id"]), false)
				}
				iapi.Api.SetGroupWholeBan(self_id, group_id, true)
				iapi.Api.SetGroupLeave(self_id, group_id)
			}
		}
	}
}
