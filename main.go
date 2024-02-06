package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/cron"
	"main.go/app/bot/logs"
	event "main.go/app/bot/message"
	"main.go/common/BaseController"
	"main.go/config/app_conf"
	"main.go/route"
	"os"
	"time"
)

func init() {
	time.Local = app_conf.TimeZone
	if app_conf.TestMode == false {
		s, err := os.Stat("./log/")

		if err != nil {
			os.Mkdir("./log", 0755)
		} else if s.IsDir() {
			os.Mkdir("./log", 0755)
		}
	}
}

func main() {

	/*Cron which no needed
	cron.BotInfoCron()
	go cron.BaseCron()
	go cron.Refresh_friend_list()
	*/
	go logs.LogsInit()
	go event.MainRouter()

	go Group.App_refresh_group_member_chan()
	go cron.Refresh_group_chan()

	go cron.Retract()
	go cron.Send_private()
	go cron.Send_group()

	go cron.Cron_auto_send()

	//go cron.BanPermenentCheck()

	//go cron.PowerCheck()

	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.Use(BaseController.CommonController())
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
