package event

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/app/bot/model/LogsModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"net"
	"sync"
)

var ClientToConn = new(sync.Map)
var ConnToClient = new(sync.Map)

func EventListener() {
	for c := range Net.WsServer_ReadChannel {
		if c.Status {
			var es EventStruct
			es.json = string(c.Message)
			err := sonic.UnmarshalString(es.json, &es)
			if err != nil {
				go LogErrorModel.Api_insert(err.Error(), tuuz.FUNCTION_ALL())
				continue
			}
			if c.Status {
				ClientToConn.Store(es.SelfId, c.Conn)
				ConnToClient.Store(c.Conn, es.SelfId)
			}
			es.remoteaddr = c.Conn.RemoteAddr()
			es.EventRouter()
		} else {
			fmt.Println(c.Conn.RemoteAddr(), "断开链接")
			client, ok := ConnToClient.Load(c.Conn)
			if ok {
				ClientToConn.Delete(client)
			}
		}
	}
}

type EventStruct struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	json        string
	remoteaddr  net.Addr
}

func (es EventStruct) EventRouter() {
	go LogsModel.Api_insert(es.json, "main", es.remoteaddr.String())

	switch es.PostType {
	case "message":
		message_type := es.MessageType
		switch message_type {
		case "private":
			var pm PrivateMessageStruct
			pm.remoteaddr = es.remoteaddr
			err := sonic.UnmarshalString(es.json, &pm)
			if err != nil {
				fmt.Println(err, es.json)
			} else {
				pm.PrivateMsg()
			}
			break

		case "group":
			var gm GroupMessageStruct
			gm.remoteaddr = es.remoteaddr
			err := sonic.UnmarshalString(es.json, &gm)
			if err != nil {
				fmt.Println(err, es.json)
			} else {
				gm.GroupMsg()
			}
			break

		default:
			Log.Crrs(errors.New("undefine route"), es.json)
			break
		}
		break

	case "notice":
		var notice Notice
		notice.remoteaddr = es.remoteaddr
		err := sonic.UnmarshalString(es.json, &notice)
		if err != nil {
			fmt.Println(err)
		} else {
			notice.NoticeMsg()
		}
		break

	case "request":
		var req Request
		req.remoteaddr = es.remoteaddr
		err := sonic.UnmarshalString(es.json, &req)
		if err != nil {
			fmt.Println(err)
		} else {
			req.RequestMsg()
		}
		break

	case "meta_event":
		//trigger the
		var aaa MetaEventStruct
		aaa.remoteaddr = es.remoteaddr
		err := sonic.UnmarshalString(es.json, &aaa)
		if err != nil {
			fmt.Println(err)
		} else {
			aaa.MetaEvent()
			//fmt.Println("HeartBeat:", aaa.SelfId, aaa.Status.QqStatus)
		}
		break

	default:
		LogRecvModel.Api_insert(es.json)
		break
	}
}
