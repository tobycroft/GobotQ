package private

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"regexp"
)

func PrivateHandle() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivate) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			selfId := pm.SelfId
			user_id := pm.UserId
			group_id := int64(0)
			message := pm.RawMessage
			rawMessage := pm.RawMessage

			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(message)
			new_text := reg.ReplaceAllString(message, "")

			if active {
				active_main_function(selfId, user_id, group_id, new_text, message)
			} else {
				//在未激活acfur的情况下应该对原始内容进行还原
				if private_default_reply(selfId, user_id, group_id, message) {
					return
				}
				auto_reply := PrivateAutoReplyModel.Api_find_byKey(message)
				if len(auto_reply) > 0 {
					if auto_reply["value"] == nil {
						return
					}
					iapi.Api.Sendprivatemsg(selfId, user_id, group_id, auto_reply["value"].(string), false)
				} else {
					private_auto_reply(selfId, user_id, group_id, message)
				}
			}
		}
	}
}
