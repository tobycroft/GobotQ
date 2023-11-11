package Group

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_groupinfo(self_id, group_id int64) {
	gl, err := iapi.Post{}.Getgrouplist(self_id)
	if err != nil {

	} else {
		GroupListModel.Api_delete_byGid(group_id)
		var gss []GroupListModel.GroupList
		for _, gll := range gl {
			var gs GroupListModel.GroupList
			if gll.GroupID != group_id {
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
		if len(gss) > 0 {
			GroupListModel.Api_insert_more(gss)
		}
	}
}
