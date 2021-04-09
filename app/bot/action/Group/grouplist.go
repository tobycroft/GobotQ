package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_group_list() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		gl, err := api.Getgrouplist(bot["bot"])
		if err != nil {

		} else {
			GroupListModel.Api_delete(bot["bot"])
			var gss []GroupListModel.GroupList
			for _, gll := range gl {
				var gs GroupListModel.GroupList
				gs.Self_id = bot["bot"]
				gs.Group_id = gll.GroupID
				gs.Group_name = gll.GroupName
				gs.Group_memo = gll.GroupMemo
				gs.Max_member_count = gll.MaxMemberCount
				gs.Member_count = gll.MemberCount
				gss = append(gss, gs)
			}
			GroupListModel.Api_insert_more(gss)
		}
	}
}
