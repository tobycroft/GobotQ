package private

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"net/netip"
	"time"
)

func Router() {
	go private_main_handler()
	go private_message_fully_attached_with_acfur()
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivate) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			pm := es.Json
			botinfo := BotModel.Api_find(pm.SelfId)
			if len(botinfo) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				continue
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if botinfo["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", botinfo["allow_ip"], ip.Addr().String()), pm.SelfId)
				continue
			}
			//if botinfo["allow_ip"] == nil {
			//	return
			//}
			//if !strings.Contains(remoteip, botinfo["allow_ip"].(string)) {
			//	Log.Errs(errors.New(fmt.Sprint(remoteip, botinfo["allow_ip"].(string))), "不允许的ip")
			//	return
			//}
			if botinfo["end_date"].(time.Time).Before(time.Now()) {
				iapi.Api.Sendprivatemsg(pm.SelfId, pm.UserId, 0, app_default.Default_over_time, false)
				continue
			}

			ps.Publish(types.MessagePrivateValid, c)
		}

	}

}
