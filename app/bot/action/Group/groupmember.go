package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
)

func App_refresh_group_member() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		gl, err := api.Getgrouplist(bot["bot"])
		if err != nil {

		} else {
			for _, gll := range gl {
				App_refresh_group_member_one(bot["bot"], gll.GIN)
			}
		}
	}
}

func App_refresh_group_member_one(bot, gid interface{}) {
	gm, err := api.Getgroupmemberlist(bot, gid)
	if err != nil {

	} else {

		for _, gmm := range gm {

		}
	}
}
