package private

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/config/app_default"
	"strings"
)

func active_main_function(self_id, user_id, group_id int64, message, origin_text string) {
	switch message {

	case "ip":
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, pm.RemoteAddr, true)
		break

	case "app", "下载":
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, app_default.Default_app_download_url, true)
		break

	case "help":
		botinfo := BotModel.Api_find(self_id)
		if len(botinfo) > 0 {
			if botinfo["owner"].(int64) == user_id {
				iapi.Api.Sendprivatemsg(self_id, user_id, group_id, app_default.Default_private_help+app_default.Default_private_help_for_RobotOwner, false)
			} else {
				iapi.Api.Sendprivatemsg(self_id, user_id, group_id, app_default.Default_private_help, false)
			}
		} else {
			iapi.Api.Sendprivatemsg(self_id, user_id, group_id, app_default.Default_private_help, false)
		}
		break

	case "测试撤回":
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, "测试撤回", true)
		break

	case "登录", "登陆", "login":
		Private.App_userLogin(self_id, user_id, group_id, message)
		break

	case "清除登录":
		Private.App_userClearLogin(self_id, user_id, group_id)
		break

	case "解绑":
		Private.App_unbind_bot(self_id, user_id, group_id, message)
		break

	case "绑定":
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
		break

	case "绑定群":
		groupbinds := BotGroupAllowModel.Api_select(self_id)
		groups := []string{}
		for _, groupbind := range groupbinds {
			groups = append(groups, Calc.Any2String(groupbind["group_id"]))
		}
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, "您的机器人可在如下群中使用:\r\n"+strings.Join(groups, ",")+
			"\r\n您可以使用：acfur绑定群:群号，来绑定新群，\r\n使用：acfur解绑群:群号，解绑", false)
		break

	default:
		privateHandle_acfur_middle(self_id, user_id, group_id, message, origin_text)
		break
	}
}
