package main

import (
	"fmt"
	"main.go/app/v1/user/action/BalanceAction"
	"main.go/tuuz"
)

func main() {

	//go cron.BaseCron()
	//go cron.Refresh_friend_list()
	//
	//go Group.App_refresh_group_member_chan()
	//
	//go cron.Refresh_group_chan()
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
	//mainroute := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	//route.OnRoute(mainroute)
	//mainroute.Run(":81")
	var bl BalanceAction.Interface
	bl.Db = tuuz.Db()
	bl.Db.Begin()
	err := bl.App_single_balance(0, "test1", 1, "test1")
	fmt.Println(err)
	//err2 := bl.App_single_balance(1, "test2", -1, "test2")
	//fmt.Println(err2)
	bl.Db.Rollback()
}
