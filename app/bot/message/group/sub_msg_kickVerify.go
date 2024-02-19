package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func kick_verify() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur) {
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
			raw_message := gm.RawMessage

			groupmember := gmr.GroupMember
			groupfunction := gmr.GroupFunction
			admin := false
			owner := false
			if len(groupmember) > 0 {
				if groupmember["role"].(string) == "admin" {
					admin = true
				}
				if groupmember["role"].(string) == "owner" {
					admin = true
					owner = true
				}
			}

			if str, ok := service.Serv_text_match(raw_message, []string{"acfur踢出验证"}); ok {
				if !admin && !owner {
					if len(groupmember) > 0 {
						service.Not_admin(self_id, group_id, user_id)
					} else {
						GroupFunction.App_reverify_death(self_id, group_id, user_id, message_id, str, groupmember, groupfunction)
					}
				}
			}
		}
	}
}
