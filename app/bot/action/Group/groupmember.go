package Group

import (
	"fmt"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/tuuz"
)

type App_group_member struct {
	SelfId  any
	GroupId any
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
			App_refresh_group_member_action(bot["self_id"], gl)
		}
	}
}

func App_refresh_group_member_action(self_id any, grouplist []gorose.Data) {
	for _, gll := range grouplist {
		var apm App_group_member
		apm.SelfId = self_id
		apm.GroupId = gll["group_id"]
		Chan_refresh_group_member <- apm
	}
}

func App_refresh_group_member_one(self_id, group_id any) {
	GroupMemberModel.Api_delete_byGid(self_id, group_id)
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
	for _, gmm := range gm {
		var g GroupMemberModel.GroupMember
		g.SelfId = self_id
		g.GroupId = gmm.GroupId
		g.UserID = gmm.UserId
		g.Nickname = gmm.Nickname
		g.Card = gmm.UserDisplayname
		g.Level = gmm.Level
		g.JoinTime = gmm.JoinTime
		g.Title = gmm.Title
		g.LastSentTime = gmm.LastSentTime
		g.Role = gmm.Role
		if len(GroupMemberModel.Api_find(gmm.GroupId, gmm.UserId)) > 0 {
			GroupMemberModel.Api_update(gmm.GroupId, gmm.UserId, g)
		} else {
			GroupMemberModel.Api_insert(g)
		}
	}
}
