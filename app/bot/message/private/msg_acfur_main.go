package private

import (
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"strings"
)

func message_fully_attached_with_acfur() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivateValid) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			pm := es.Json
			self_id := pm.SelfId
			user_id := pm.UserId
			group_id := int64(0)
			message := pm.Message
			//raw_message := pm.RawMessage
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
				switch new_text {
				case "ip":
					msg := MessageBuilder.IMessageBuilder{}.Text(es.RemoteAddr)
					iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, true)
					break

				case "app", "下载":
					msg := MessageBuilder.IMessageBuilder{}.Text(app_default.Default_app_download_url)
					iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, true)
					break

				case "help":
					botinfo := BotModel.Api_find(self_id)
					if len(botinfo) > 0 {
						if Calc.Any2Int64(botinfo["owner"]) == user_id {
							msg := MessageBuilder.IMessageBuilder{}.Text(app_default.Default_private_help + app_default.Default_private_help_for_RobotOwner)
							iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
						} else {
							msg := MessageBuilder.IMessageBuilder{}.Text(app_default.Default_private_help)
							iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
						}
					} else {
						msg := MessageBuilder.IMessageBuilder{}.Text(app_default.Default_private_help)
						iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
					}
					break

				case "测试撤回":
					iapi.Api.SendPrivateMsg(self_id, user_id, group_id, MessageBuilder.IMessageBuilder{}.Text("测试撤回"), true)
					break

				case "登录", "登陆", "login":
					Private.App_userLogin(self_id, user_id, group_id, new_text)
					break

				case "清除登录":
					Private.App_userClearLogin(self_id, user_id, group_id)
					break

				case "解绑":
					Private.App_unbind_bot(self_id, user_id, group_id, new_text)
					break

				case "绑定":
					iapi.Api.SendPrivateMsg(self_id, user_id, group_id, MessageBuilder.IMessageBuilder{}.Text("请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人"), false)
					break

				case "绑定群":
					groupbinds := BotGroupAllowModel.Api_select(self_id)
					groups := []string{}
					for _, groupbind := range groupbinds {
						groups = append(groups, Calc.Any2String(groupbind["group_id"]))
					}
					iapi.Api.SendPrivateMsg(self_id, user_id, group_id, MessageBuilder.IMessageBuilder{}.Text("您的机器人可在如下群中使用:\r\n"+strings.Join(groups, ",")+
						"\r\n您可以使用：acfur绑定群:群号，来绑定新群，\r\n使用：acfur解绑群:群号，解绑"), false)
					break

				default:
					ps.Publish(types.MessagePrivateAcfur, c.Payload)
					break
				}
			}

		}
	}
}
