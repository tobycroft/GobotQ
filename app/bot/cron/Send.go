package cron

import "main.go/app/bot/iapi"

func Send_private() {
	iapi.Api.Send_private()
}

func Send_group() {
	iapi.Api.Send_group()
}
