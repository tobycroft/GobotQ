package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"net/netip"
	"regexp"
	"time"
)

func Router() {

	go group_message_acfur_when_fully_matched()
	go group_message_acfur_semi_match()

	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroup) {
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
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
			gm := es.Json
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

			text := message
			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(text)
			//new_text := reg.ReplaceAllString(text, "")
			groupmember := GroupMemberModel.Api_find(group_id, user_id)
			groupfunction := GroupFunctionModel.Api_find(group_id)
			if len(groupfunction) < 1 {
				GroupFunctionModel.Api_insert(group_id)
				groupfunction = GroupFunctionModel.Api_find(group_id)
			}
			if botinfo["end_date"].(time.Time).Before(time.Now()) {
				Group.AutoMessage(self_id, group_id, user_id, app_default.Default_over_time, groupfunction)
				continue
			}
			sender := gm.Sender
			if !active || !service.Serv_is_at_me(self_id, message) {
				//在未激活acfur的情况下应该对原始内容进行还原
				go groupHandle_acfur_middle(self_id, group_id, user_id, message_id, message, raw_message, sender, groupmember, groupfunction)
				ps.Publish(types.MessageGroupValid, c.Payload)
			}

		}

	}

}
