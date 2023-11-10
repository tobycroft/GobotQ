package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
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

type Send struct {
	Action string `json:"action"`
	Params struct {
		UserId  int    `json:"user_id"`
		Message string `json:"message"`
	} `json:"params"`
	Echo string `json:"echo"`
}

func main() {
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
	go func() {
		for c := range Net.WsServer_ReadChannel {
			fmt.Println("ws:", string(c.Message))
		}
	}()

	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}

type T struct {
	Time          int    `json:"time"`
	SelfId        int    `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
	Status        struct {
		Self struct {
			Platform string `json:"platform"`
			UserId   int    `json:"user_id"`
		} `json:"self"`
		Online   bool   `json:"online"`
		Good     bool   `json:"good"`
		QqStatus string `json:"qq.status"`
	} `json:"status"`
	Interval int `json:"interval"`
}

type T2 struct {
	Time        int    `json:"time"`
	Time        int    `json:"time"`
	SelfId      int    `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int    `json:"message_id"`
	TargetId    int    `json:"target_id"`
	PeerId      int    `json:"peer_id"`
	UserId      int    `json:"user_id"`
	Message     []struct {
		Data struct {
			Text string `json:"text"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"message"`
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	Sender     struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
		Title    string `json:"title"`
		Level    string `json:"level"`
	} `json:"sender"`

	SelfId      int    `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int    `json:"message_id"`
	TargetId    int    `json:"target_id"`
	PeerId      int    `json:"peer_id"`
	UserId      int    `json:"user_id"`
	Message     []struct {
		Data struct {
			Text string `json:"text"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"message"`
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	Sender     struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
		Title    string `json:"title"`
		Level    string `json:"level"`
	} `json:"sender"`
}
