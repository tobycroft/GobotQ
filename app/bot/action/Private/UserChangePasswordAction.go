package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
)

func App_userChangePassword(bot int, uid int, text string) {
	if UserMemberModel.Api_update_password(uid, text) {
		api.Sendprivatemsg(bot, uid, "您的密码已被修改为：\r\n↓↓↓↓↓↓↓↓"+text+"↑↑↑↑↑↑↑↑")
	} else {
		api.Sendprivatemsg(bot, uid, app_default.Default_error_alert)
	}
}
