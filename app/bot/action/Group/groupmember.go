package Group

import (
	"fmt"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/redis/GroupMemberRedis"
	"main.go/tuuz"
)

type App_group_member struct {
	SelfId  int64
	GroupId int64
}

var Chan_refresh_group_member = make(chan App_group_member, 99)

func App_refresh_group_member_chan() {
	for gm := range Chan_refresh_group_member {
		App_refresh_group_member_one(gm.SelfId, gm.GroupId)
	}
}

func App_refresh_group_member() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		gl := GroupListModel.Api_select(bot["self_id"])
		if len(gl) > 0 {
			App_refresh_group_member_action(Calc.Any2Int64(bot["self_id"]), gl)
		}
	}
}

func App_refresh_group_member_action(self_id int64, grouplist []gorose.Data) {
	for _, gll := range grouplist {
		var apm App_group_member
		apm.SelfId = self_id
		apm.GroupId = gll["group_id"].(int64)
		Chan_refresh_group_member <- apm
	}
}

func App_refresh_group_member_one(self_id, group_id int64) {
	GroupMemberModel.Api_delete_byGid(self_id, group_id)
	GroupMemberRedis.Cac_del(self_id, "*", group_id)
	gm, err := iapi.Api.Getgroupmemberlist(self_id, group_id)
	if err != nil {
		fmt.Println(tuuz.FUNCTION_ALL(), err)
	} else {
		if len(gm) > 0 {
			App_refresh_group_member_one_action(self_id, gm)
		}
	}
}

func App_refresh_group_member_one_action(self_id any, gm []iapi.GroupMemberList) {
	var gss []GroupMemberModel.GroupMember
	for _, gmm := range gm {
		var gs GroupMemberModel.GroupMember
		gs.SelfId = self_id
		gs.GroupId = gmm.GroupId
		gs.UserId = gmm.UserId
		gs.Nickname = gmm.Nickname
		gs.Card = gmm.UserDisplayname
		gs.Level = gmm.Level
		gs.JoinTime = gmm.JoinTime
		gs.Title = gmm.Title
		gs.LastSentTime = gmm.LastSentTime
		gs.Role = gmm.Role
		gss = append(gss, gs)
		GroupMemberRedis.Cac_set(self_id, gs.UserId, gs.GroupId, gs)
	}
	GroupMemberModel.Api_insert_more(gss)
}
