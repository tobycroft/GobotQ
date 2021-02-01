package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_groupinfo(bot int, gid int) {
	gl, err := api.Getgrouplist(bot)
	if err != nil {

	} else {
		GroupListModel.Api_delete_byGid(gid)
		var gss []GroupListModel.GroupList
		for _, gll := range gl {
			var gs GroupListModel.GroupList
			if gll.GIN != gid {
				continue
			}
			gs.Bot = bot
			gs.Gid = gll.GIN
			gs.Group_name = gll.StrGroupName
			gs.Group_memo = gll.StrGroupMemo
			gs.Owner = gll.DwGroupOwnerUin
			gs.Number = gll.DwMemberNum
			gss = append(gss, gs)
		}
		GroupListModel.Api_insert_more(gss)
		api.Sendgroupmsg(bot, gid, "群信息刷新完成", true)
	}
}
