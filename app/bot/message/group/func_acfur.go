package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/Group"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func ban_group() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banGroup) {
		var es GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := es.Json

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId

			//the message in json format
			//message := gm.Message

			raw_message := gm.RawMessage

			text := raw_message
			Group.App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_group)
		}
	}

}

func ban_url() {
	Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_url)

}

func ban_wx() {
	go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_weixin)

}

func ban_share() {
	go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_share)

}
