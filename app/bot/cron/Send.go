package cron

import (
	"main.go/app/bot/iapi/apipost"
)

func Send_private() {
	apipost.Api{}.Send_private()
}

func Send_group() {
	apipost.Api{}.Send_group()
}
