package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/UserTokenModel"
)

func App_userClearLogin(bot int, uid int) {
	if UserTokenModel.Api_delete(uid) {
		api.Sendprivatemsg(bot, uid, "您的登录状态已经全部清空，如需再次登录请发送acfur登录", false)
	}
}
