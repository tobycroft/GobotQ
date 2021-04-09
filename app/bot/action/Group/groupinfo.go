package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_groupinfo(self_id int, gid int64) {
	gl, err := api.Getgrouplist(self_id)
	if err != nil {

	} else {
		GroupListModel.Api_delete_byGid(gid)
		var gss []GroupListModel.GroupList
		for _, gll := range gl {
			var gs GroupListModel.GroupList
			if gll.GroupID != gid {
				continue
			}
			gs.Self_id = self_id
			gs.Group_id = gll.GroupID
			gs.Group_name = gll.GroupName
			gs.Group_memo = gll.GroupMemo
			gs.Member_count = gll.MemberCount
			gs.Max_member_count = gll.MaxMemberCount
			gss = append(gss, gs)
		}
		GroupListModel.Api_insert_more(gss)
	}
}
