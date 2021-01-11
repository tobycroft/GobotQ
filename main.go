package main

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/cron"
	"main.go/route"
)

func main() {

	go cron.BaseCron()

	go Group.App_refresh_group_member_chan()

	go cron.Retract()
	go cron.Send_private()
	go cron.Send_group()
	go cron.Send_temp()

	go cron.GroupMsgRecv()
	go cron.PrivateMsgRecv()

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
