package Group

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_groupinfo(self_id, group_id int64) {
	gl, err := iapi.Api.Getgrouplist(self_id)
	if err != nil {

	} else {
		GroupListModel.Api_delete_byGid(group_id)
		var gss []GroupListModel.GroupList
		for _, gll := range gl {
			var gs GroupListModel.GroupList
			if gll.GroupId != group_id {
				continue
			}
			gs.SelfId = self_id
			gs.GroupId = gll.GroupId
			gs.GroupName = gll.GroupName
			gs.GroupMemo = gll.GroupRemark
			gs.MemberCount = gll.MemberCount
			gs.MaxMemberCount = gll.MaxMemberCount
			gss = append(gss, gs)
		}
		if len(gss) > 0 {
			GroupListModel.Api_insert_more(gss)
		}
	}
}
