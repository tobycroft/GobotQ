package event

import (
	"fmt"
	"net"
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
	switch self.MetaEventType {
	case "lifecycle":
		//self.SelfId
		break

	case "heartbeat":
		break

	default:
		fmt.Println("request no route", self)
		break

	}

}
