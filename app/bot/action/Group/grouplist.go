package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
)

func App_fresh_group_list() {
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
			}
			GroupListModel.Api_insert_more(gss)
		}
	}
}
