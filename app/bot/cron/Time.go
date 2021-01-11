package cron

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupListModel"
	"time"
)

func BaseCron() {
	for {
		bots := BotModel.Api_select()
		for _, bot := range bots {
			gl, err := api.Getgrouplist(bot["bot"])
			if err != nil {

			} else {
				GroupListModel.Api_delete(bot["bot"])
				var gss []GroupListModel.GroupList
				for _, gll := range gl {
					var gs GroupListModel.GroupList
					gs.Bot = bot["bot"]
					gs.Gid = gll.GIN
					gs.Group_name = gll.StrGroupName
					gs.Group_memo = gll.StrGroupMemo
					gs.Owner = gll.DwGroupOwnerUin
					gs.Number = gll.DwMemberNum
					gss = append(gss, gs)
					var gm Group.App_group_member
					gm.Gid = gll.GIN
					gm.Bot = bot["bot"]
					gm.Owner = gll.DwGroupOwnerUin
					Group.Chan_refresh_group_member <- gm
					if len(GroupFunctionModel.Api_find(gll.GIN)) < 1 {
						GroupFunctionModel.Api_insert(gll.GIN)
					}
				}
				GroupListModel.Api_insert_more(gss)
			}
		}
		time.Sleep(3600 * time.Second)
	}
}
