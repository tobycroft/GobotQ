package Group

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/redis/GroupListRedis"
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
	GroupListRedis.Cac_del(self_id, "*")
	GroupListModel.Api_delete(self_id)
	var gss []GroupListModel.GroupList
	for _, gll := range gl {
		var gs GroupListModel.GroupList
		gs.SelfId = self_id
		gs.GroupId = gll.GroupId
		gs.GroupName = gll.GroupName
		gs.GroupMemo = gll.GroupRemark
		gs.MaxMemberCount = gll.MaxMemberCount
		gs.MemberCount = gll.MemberNum
		gss = append(gss, gs)
		GroupListRedis.Cac_set(self_id, gs.GroupId, gs)
	}
	if len(gss) > 0 {
		GroupListModel.Api_insert_more(gss)
	}

}
