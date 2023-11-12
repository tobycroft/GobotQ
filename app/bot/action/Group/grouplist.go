package Group

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_group_list() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		gl, err := iapi.Api.Getgrouplist(bot["self_id"])
		if err != nil {

		} else {
			App_refresh_group_list_action(bot["self_id"], gl)
		}
	}
}

func App_refresh_group_list_action(self_id any, gl []iapi.GroupList) {
	GroupListModel.Api_delete(self_id)
	var gss []GroupListModel.GroupList
	for _, gll := range gl {
		var gs GroupListModel.GroupList
		gs.Self_id = self_id
		gs.Group_id = gll.GroupId
		gs.Group_name = gll.GroupName
		gs.Group_memo = gll.GroupRemark
		gs.Max_member_count = gll.MaxMemberCount
		gs.Member_count = gll.MemberCount
		gss = append(gss, gs)
	}
	if len(gss) > 0 {
		GroupListModel.Api_insert_more(gss)
	}
}
