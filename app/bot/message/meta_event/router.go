package meta_event

import (
	"fmt"
	"github.com/bytedance/sonic"
	"log"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/netip"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageMetaEvent) {
		var es EventStruct[MetaEventStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			bot := BotModel.Api_find(es.SelfId)
			if len(bot) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				continue
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if bot["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), es.SelfId)
				continue
			}
			switch es.Json.MetaEventType {
			case "lifecycle":
				_, err := iapi.Api.GetLoginInfo(es.SelfId)
				if err != nil {
					log.Println(err)
				}
				fmt.Println(es.Json, es.SelfId)
				break

			case "heartbeat":
				fmt.Println(es.Json, es.SelfId)
				break

			default:
				fmt.Println("request no route", es)
				break
			}

			//fmt.Println("HeartBeat:", aaa.SelfId, aaa.Status.QqStatus)
		}

	}

}
