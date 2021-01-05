package main

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/cron"
	"main.go/route"
)

func main() {

	go Group.App_refresh_group_member_chan()

	go cron.Retract()

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
