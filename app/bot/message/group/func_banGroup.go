package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func ban_group() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banGroup) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			if gmr.GroupFunction["ban_retract"].(int64) == 1 {
				var rm iapi.RetractMessage
				rm.MessageId = message_id
				rm.SelfId = self_id
				rm.Time = 0
				ps.Publish_struct(types.RetractChannel, rm)
			}
			Group.App_kick_user(self_id, group_id, user_id, true, gmr.GroupFunction, app_default.Default_ban_group)
		}
	}
}
