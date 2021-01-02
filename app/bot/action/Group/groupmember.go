package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupMemberModel"
)

type App_group_member struct {
	Bot interface{}
	Gid interface{}
}

var Chan_refresh_group_member = make(chan App_group_member, 99)

func App_refresh_group_member_chan() {
	for gm := range Chan_refresh_group_member {
		App_refresh_group_member_one(gm.Bot, gm.Gid)
	}
}

func App_refresh_group_member() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		gl, err := api.Getgrouplist(bot["bot"])
		if err != nil {

		} else {
			for _, gll := range gl {
				var apm App_group_member
				apm.Bot = bot["bot"]
				apm.Gid = gll.GIN
				Chan_refresh_group_member <- apm
			}
		}
	}
}

func App_refresh_group_member_one(bot, gid interface{}) {
	gm, err := api.Getgroupmemberlist(bot, gid)
	if err != nil {

	} else {
		GroupMemberModel.Api_delete_byGid(bot, gid)
		var gms []GroupMemberModel.GroupMember
		for _, gmm := range gm {
			var g GroupMemberModel.GroupMember
			g.Bot = bot
			g.Gid = gid
			g.Uid = gmm.UIN
			g.Remark = gmm.Remark
			g.Nickname = gmm.NickName
			g.Age = gmm.Age
			g.Card = gmm.Card
			g.Grouplevel = gmm.GroupLevel
			g.Jointime = gmm.AddGroupTime
			g.Title = gmm.SpecTitle
			g.Lastsend = gmm.LastMsgTime
			gms = append(gms, g)
		}
		GroupMemberModel.Api_insert_more(gms)
	}
}
