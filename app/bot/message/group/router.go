package group

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/netip"
)

func Router() {
	go message_main_handler()
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroup) {
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := es.Json
			bot := BotModel.Api_find(gm.SelfId)
			if len(bot) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				continue
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if bot["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), gm.SelfId)
				continue
			}
			is_self := false

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			//message := gm.Message
			message := gm.RawMessage
			raw_message := gm.RawMessage

			if user_id == self_id {
				is_self = true
			}

			if !is_self {
				botinfo := BotModel.Api_find(self_id)
				if len(botinfo) < 1 {
					Log.Crrs(errors.New("bot_not_found"), Calc.Any2String(self_id))
					break
				}

				has1 := Redis.CheckExists("__groupinfo__" + Calc.Int642String(group_id) + "_" + Calc.Int642String(user_id))
				has2 := Redis.CheckExists("__userinfo__" + Calc.Int642String(group_id) + "_" + Calc.Int642String(user_id))
				if !has1 || !has2 {
					ps.Publish_struct(types.OperationRefreshGroup, RefreshGroupStruct{
						GroupId: group_id,
						SelfId:  self_id,
						UserId:  user_id,
					})
				}
				ps.Publish(types.MessageGroupValid, c.Payload)
			} else {

			}
		}

	}

}
