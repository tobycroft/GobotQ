package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/cron"
	"main.go/app/bot/logs"
	event "main.go/app/bot/message"
	"main.go/common/BaseController"
	"main.go/config/app_conf"
	"main.go/route"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
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

	go GroupFunction.App_refresh_group_member_chan()
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

func main2() {
	//rs := Redis.StreamNew("test")
	//fmt.Println(rs.Publish(map[string]any{"sss": "bbb"}))
	//fmt.Println(rs.XRange())
	chunk_test()
}

func chunk_test() {
	db := tuuz.Db().Table("bot_default_reply")
	db.Where("id", 1)
	db.ChunkWG(1, 1, func(data []gorose.Data) error {
		fmt.Println(Jsong.Encode(data))
		return nil
	})

}
