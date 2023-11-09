package cron

import (
	"errors"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

func BaseCron() {
	for {
		bots := BotModel.Api_select()
		for _, bot := range bots {
			gl, err := api.Getgrouplist(bot["self_id"])
			if err != nil {

			} else {
				GroupListModel.Api_delete(bot["self_id"])
				var gss []GroupListModel.GroupList
				for _, gll := range gl {
					var gs GroupListModel.GroupList
					gs.Self_id = bot["self_id"]
					gs.Group_id = gll.GroupID
					gs.Group_name = gll.GroupName
					gs.Group_memo = gll.GroupMemo
					gs.Member_count = gll.MemberCount
					gs.Max_member_count = gll.MaxMemberCount
					gss = append(gss, gs)
					var gm Group.App_group_member
					gm.GroupId = gll.GroupID
					gm.SelfId = bot["self_id"]
					Group.Chan_refresh_group_member <- gm
					if len(GroupFunctionModel.Api_find(gll.GroupID)) < 1 {
						GroupFunctionModel.Api_insert(gll.GroupID)
					}
				}
				if len(gss) > 0 {
					GroupListModel.Api_insert_more(gss)
				}
			}
		}
		time.Sleep(3600 * time.Second)
	}
}

func BotInfoCron() {
	for {
		bots := BotModel.Api_select()
		for _, bot := range bots {
			bot_info, err := api.GetLoginInfo(bot["self_id"])
			if err != nil {

			} else {
				self_id := bot_info.UserID
				name := bot_info.Nickname
				if !BotModel.Api_update_cname(self_id, name) {
					Log.Crrs(errors.New("机器人用户名无法更新"), tuuz.FUNCTION_ALL())
				}
			}
		}
		time.Sleep(30 * time.Minute)
	}
}
