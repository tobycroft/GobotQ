package private

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/config/types"
	"main.go/extend/Aigc"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
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
			message := pm.RawMessage

			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(message)
			if !active {
				//在未激活acfur的情况下应该对原始内容进行还原
				if private_default_reply(selfId, user_id, group_id, message) {
					continue
				}
				auto_reply := PrivateAutoReplyModel.Api_find_byKey(message)
				if len(auto_reply) > 0 {
					if auto_reply["value"] == nil {
						continue
					}
					iapi.Api.SendPrivateMsg(selfId, user_id, group_id, auto_reply["value"].(string), false)
					continue
				} else {
					private_auto_reply(selfId, user_id, group_id, message)
				}
				ai_reply, err := Aigc.Aigc_gemini_text(message)
				if err != nil {
					continue
				}
				iapi.Api.SendPrivateMsg(selfId, user_id, group_id, ai_reply.Data.Text, false)
				continue
			}
		}
	}
}
