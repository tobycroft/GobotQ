package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func set_setting() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + setting) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
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
			if !admin && !owner {
				if len(groupmember) > 0 {
					service.Not_admin(self_id, group_id, user_id)
				} else {
					if str, ok := service.Serv_text_match(raw_message, []string{"acfur设定"}); ok {
						GroupFunction.App_group_function_set(self_id, group_id, user_id, str, groupfunction)
					}
				}
			}

		}
	}
}
