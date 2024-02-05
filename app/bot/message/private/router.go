package private

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"net/netip"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivate) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			pm := es.Json
			bot := BotModel.Api_find(pm.SelfId)
			if len(bot) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				return
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if bot["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), pm.SelfId)
				return
			}
			PrivateMsgChan <- pm
			selfId := pm.SelfId
			user_id := pm.UserId
			group_id := int64(0)
			message := pm.RawMessage
			rawMessage := pm.RawMessage

			pm.PrivateHandle(selfId, user_id, group_id, message, rawMessage)
		}

	}

}
