package Group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/redis/GroupMemberRedis"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

type App_group_member struct {
	SelfId  int64 `json:"self_id"`
	GroupId int64 `json:"group_id"`
}

//var Chan_refresh_group_member = make(chan App_group_member, 99)

func App_refresh_group_member_chan() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.RetractChannel) {
		var gm App_group_member
		err := sonic.UnmarshalString(c.Payload, &gm)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			continue
		}
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
		apm.GroupId = Calc.Any2Int64(gll["group_id"])
		Redis.PubSub{}.Publish_struct(types.RefreshGroupMembers, apm)
	}
}

func App_refresh_group_member_one(self_id, group_id int64) {
	GroupMemberModel.Api_delete_byGid(self_id, group_id)
	GroupMemberRedis.Cac_del(self_id, "*", group_id)
	gm, err := iapi.Api.GetGroupMemberList(self_id, group_id)
	if err != nil {
		fmt.Println(tuuz.FUNCTION_ALL(), err)
	} else {
		if len(gm) > 0 {
			App_refresh_group_member_one_action(self_id, gm)
		}
	}
}

func App_refresh_group_member_one_action(self_id int64, gm []iapi.GroupMemberList) {
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
