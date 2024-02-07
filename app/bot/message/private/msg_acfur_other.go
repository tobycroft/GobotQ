package private

import (
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
)

func message_setting_change_with_acfur() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivateAcfur) {
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
			new_text := reg.ReplaceAllString(message, "")
			if active {
				botinfo := BotModel.Api_find(selfId)
				if msg, ok := service.Serv_text_match(new_text, []string{"密码"}); ok {
					if int64(user_id) == botinfo["owner"].(int64) {
						Private.App_userChangePassword(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else if msg, ok := service.Serv_text_match(new_text, []string{"绑定"}); ok {
					Private.App_bind_robot(selfId, user_id, group_id, msg)
				} else if msg, ok := service.Serv_text_match(new_text, []string{"修改密码"}); ok {
					if int64(user_id) == botinfo["owner"].(int64) {
						Private.App_change_bot_secret(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else if msg, ok := service.Serv_text_match(new_text, []string{"绑定群"}); ok {
					if int64(user_id) == botinfo["owner"].(int64) {
						Private.App_bind_group(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else if msg, ok := service.Serv_text_match(new_text, []string{"解绑群"}); ok {
					if int64(user_id) == botinfo["owner"].(int64) {
						Private.App_unbind_group(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else {
					iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
				}
			}
		}
	}
}
