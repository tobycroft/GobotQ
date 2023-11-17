package Group

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupAdminModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/redis/GroupListRedis"
	"main.go/tuuz/Jsong"
)

func App_refresh_groupinfo(self_id, group_id int64) {
	gl, err := iapi.Api.Getgrouplist(self_id)
	if err != nil {

	} else {
		GroupListRedis.Cac_del(self_id, "*")
		GroupListModel.Api_delete(self_id)
		GroupAdminModel.Api_delete_bySelfIdAndGroupId(self_id, nil)
		var gss []GroupListModel.GroupList
		var gas []GroupAdminModel.GroupAdmins
		for _, gll := range gl {
			var gs GroupListModel.GroupList
			gs.SelfId = self_id
			gs.GroupId = gll.GroupId
			gs.GroupName = gll.GroupName
			gs.GroupMemo = gll.GroupRemark
			gs.MemberCount = gll.MemberCount
			gs.MaxMemberCount = gll.MaxMemberCount
			gs.Admins, _ = Jsong.Encode(gll.Admins)
			gss = append(gss, gs)
			for _, admin := range gll.Admins {
				gas = append(gas, GroupAdminModel.GroupAdmins{
					SelfId:  self_id,
					GroupId: gll.GroupId,
					UserId:  admin,
				})
			}
		}
		if len(gss) > 0 && len(gas) > 0 {
			GroupListModel.Api_insert_more(gss)
			GroupAdminModel.Api_insert_more(gas)
		}
	}
}
