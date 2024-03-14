package private

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/config/types"
	"main.go/extend/Aigc"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"strings"
	"unicode/utf8"
)

func message_main_handler() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivateValid) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			pm := es.Json
			selfId := pm.SelfId
			user_id := pm.UserId
			group_id := int64(0)
			message := pm.Message

			normal_text := strings.Builder{}
			for _, msg := range message {
				switch msg.Type {
				case "at":
					break

				case "text":
					normal_text.WriteString(msg.Data["text"])
					break
				}
			}
			text := normal_text.String()

			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(text)
			if !active {
				//在未激活acfur的情况下应该对原始内容进行还原
				if private_default_reply(selfId, user_id, group_id, text) {
					continue
				}
				auto_reply := PrivateAutoReplyModel.Api_find_byKey(message)
				if len(auto_reply) > 0 {
					if auto_reply["value"] != nil {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.Text(auto_reply["value"].(string)), false)
						continue
					}
				} else {
					private_auto_reply(selfId, user_id, group_id, text)
				}
				if utf8.RuneCountInString(text) > 2 {
					ai_reply, err := Aigc.Aigc_bing_text(text)
					if err != nil {
						fmt.Println(err)
						Log.Crrs(err, tuuz.FUNCTION_ALL())
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.Text(err.Error()), false)
						continue
					}
					iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.Text(ai_reply.Echo), false)
					continue
				}
			}
		}
	}
}
