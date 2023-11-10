package main

import (
	"main.go/app/bot/api"
	"main.go/config/app_conf"
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

type Send struct {
	Action string `json:"action"`
	Params struct {
		UserId  int    `json:"user_id"`
		Message string `json:"message"`
	} `json:"params"`
	Echo string `json:"echo"`
}

func main() {
	api.Getgrouplist(123)
	//cron.BotInfoCron()
	//go cron.BaseCron()
	//go cron.Refresh_friend_list()
	//
	//go Group.App_refresh_group_member_chan()
	//
	//go cron.Refresh_group_chan()
	////go cron.Refresh_group_chan()
	////go cron.Refresh_group_chan()
	////go cron.Refresh_group_chan()
	////go cron.Refresh_group_chan()
	//
	//go cron.Retract()
	//go cron.Send_private()
	//go cron.Send_group()
	//
	//go cron.GroupMsgRecv()
	//go cron.PrivateMsgRecv()
	//
	//go cron.Cron_auto_send()
	//
	//go cron.BanPermenentCheck()
	//
	//go cron.PowerCheck()
	//
	//Calc.RefreshBaseNum()
	//mainroute := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	//mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	//mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	//route.OnRoute(mainroute)
	//mainroute.Run(":80")

}
