package cron

import "main.go/app/bot/apipost"

func Send_private() {
	apipost.ApiPost{}.Send_private()
}

func Send_group() {
	apipost.ApiPost{}.Send_group()
}
