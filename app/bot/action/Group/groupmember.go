package Group

import (
	"fmt"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupMemberModel"
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
		gl, err := api.Getgrouplist(bot["self_id"])
		if err != nil {

		} else {
			for _, gll := range gl {
				var apm App_group_member
				apm.SelfId = bot["self_id"].(int64)
				apm.GroupId = gll.GroupID
				Chan_refresh_group_member <- apm
			}
		}
	}
}

func App_refresh_group_member_one(self_id, group_id int64) {
	gm, err := api.Getgroupmemberlist(self_id, group_id)
	if err != nil {
		fmt.Println(tuuz.FUNCTION_ALL(), err)
	} else {
		GroupMemberModel.Api_delete_byGid(self_id, group_id)
		var gms []GroupMemberModel.GroupMember
		for _, gmm := range gm {
			var g GroupMemberModel.GroupMember
			g.SelfId = self_id
			g.GroupID = group_id
			g.UserID = gmm.UserID
			g.Nickname = gmm.Nickname
			g.Card = gmm.Card
			g.Level = gmm.Level
			g.JoinTime = gmm.JoinTime
			g.Title = gmm.Title
			g.LastSentTime = gmm.LastSentTime
			gms = append(gms, g)
		}
		GroupMemberModel.Api_insert_more(gms)
	}
}
