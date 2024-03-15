package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/cron"
	"main.go/app/bot/logs"
	event "main.go/app/bot/message"
	"main.go/common/BaseController"
	"main.go/config/app_conf"
	"main.go/extend/STT"
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
	url := "https://27.44.121.148:443/?ver=2&rkey=3062020101045b3059020101020100020447e8852f042431387a4f784437627245326c4c334c6137756e366638236b39767558725f63475f414249020465f47d17041f0000000866696c6574797065000000013000000005636f64656300000001310400&voice_codec=1&filetype=0&client_proto=qq&client_appid=537182769&client_type=android&client_ver=8.9.88&client_down_type=auto&client_aio_type=unk"
	str, err := STT.Audio{}.New().SpeechToText(url)
	fmt.Println(str, err)
}
