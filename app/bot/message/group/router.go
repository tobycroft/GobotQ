package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/netip"
	"time"
)

func Router() {

	go group_message_acfur_when_fully_matched()
	go group_message_acfur_semi_match()
	go group_message_normal()

	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroup) {
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			botinfo := BotModel.Api_find(es.SelfId)
			if len(botinfo) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				continue
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if botinfo["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", botinfo["allow_ip"], ip.Addr().String()), es.SelfId)
				continue
			}
			if botinfo["end_date"].(time.Time).Before(time.Now()) {
				continue
			}

			gm := es.Json
			is_self := false

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			//message_id := gm.MessageId
			//message := gm.Message
			//message := gm.RawMessage
			//raw_message := gm.RawMessage

			if user_id == self_id {
				is_self = true
			}

			if !is_self {
				has1 := Redis.CheckExists("__groupinfo__" + Calc.Int642String(group_id) + "_" + Calc.Int642String(user_id))
				has2 := Redis.CheckExists("__userinfo__" + Calc.Int642String(group_id) + "_" + Calc.Int642String(user_id))
				if !has1 || !has2 {
					ps.Publish_struct(types.OperationRefreshGroup, RefreshGroupStruct{
						GroupId: group_id,
						SelfId:  self_id,
						UserId:  user_id,
					})
				}
			}

			ps.Publish(types.MessageGroupValid, c.Payload)

			//text := raw_message
			//reg := regexp.MustCompile("(?i)^acfur")
			//active := reg.MatchString(text)
			//new_text := reg.ReplaceAllString(text, "")
			//groupmember := GroupMemberModel.Api_find(group_id, user_id)
			//groupfunction := GroupFunctionModel.Api_find(group_id)
			//if len(groupfunction) < 1 {
			//	GroupFunctionModel.Api_insert(group_id)
			//	groupfunction = GroupFunctionModel.Api_find(group_id)
			//}

			//sender := gm.Sender
			//if !active || !service.Serv_is_at_me(self_id, raw_message) {
			//在未激活acfur的情况下应该对原始内容进行还原
			//}

		}

	}

}
