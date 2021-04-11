package cron

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/model/BotModel"
)

func PowerCheck() {
	bots := BotModel.Api_select()
	for _, bot := range bots {

		if Group.BotPowerRefresh(group_id, self_id) == "" {

		}
	}
}
