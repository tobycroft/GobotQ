package event

import (
	"fmt"
	"log"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogErrorModel"
	"net"
	"net/netip"
)

type MetaEventStruct struct {
	remoteaddr    net.Addr
	Time          int64  `json:"time"`
	SelfId        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
	Status        struct {
		Self struct {
			Platform string `json:"platform"`
			UserId   int64  `json:"user_id"`
		} `json:"self"`
		Online   bool   `json:"online"`
		Good     bool   `json:"good"`
		QqStatus string `json:"qq.status"`
	} `json:"status"`
	Interval int64 `json:"interval"`
}

func (self MetaEventStruct) MetaEvent() {
	bot := BotModel.Api_find(self.SelfId)
	if len(bot) < 1 {
		LogErrorModel.Api_insert("bot bot found", self.remoteaddr.String())
		return
	}
	ip := netip.MustParseAddrPort(self.remoteaddr.String())
	if bot["allow_ip"] != ip.Addr().String() {
		LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), self.SelfId)
		return
	}
	switch self.MetaEventType {
	case "lifecycle":
		_, err := iapi.Api.GetLoginInfo(self.SelfId)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(self.MetaEventType, self.SelfId)
		break

	case "heartbeat":
		fmt.Println(self.MetaEventType, self.SelfId)
		break

	default:
		fmt.Println("request no route", self)
		break
	}
}
