package cron

import "main.go/app/bot/api"

func Send_private() {
	api.Send_private()
}

func Send_group() {
	api.Send_group()
}

func Send_temp() {
	api.Send_GroupTempMsg()
}
