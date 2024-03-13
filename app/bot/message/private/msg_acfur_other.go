package private

import (
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"strings"
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
			new_text := reg.ReplaceAllString(text, "")
			if active {
				botinfo := BotModel.Api_find(selfId)
				if msg, ok := service.Serv_text_match(new_text, []string{"密码"}); ok {
					if int64(user_id) == Calc.Any2Int64(botinfo["owner"]) {
						Private.App_userChangePassword(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else if msg, ok := service.Serv_text_match(new_text, []string{"绑定"}); ok {
					Private.App_bind_robot(selfId, user_id, group_id, msg)
				} else if msg, ok := service.Serv_text_match(new_text, []string{"修改密码"}); ok {
					if int64(user_id) == Calc.Any2Int64(botinfo["owner"]) {
						Private.App_change_bot_secret(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else if msg, ok := service.Serv_text_match(new_text, []string{"绑定群"}); ok {
					if int64(user_id) == Calc.Any2Int64(botinfo["owner"]) {
						Private.App_bind_group(selfId, user_id, group_id, msg)
					} else {
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
					}
				} else if msg, ok := service.Serv_text_match(new_text, []string{"解绑群"}); ok {
					if int64(user_id) == Calc.Any2Int64(botinfo["owner"]) {
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
