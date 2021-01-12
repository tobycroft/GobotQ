package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
	"main.go/tuuz"
)

func App_userChangePassword(bot int, uid int, text string) {
	if len(text) < 1 {
		api.Sendprivatemsg(bot, uid, "密码长度应该大于1位", true)
		return
	}
	if len(text) > 16 {
		api.Sendprivatemsg(bot, uid, "密码长度应该小于等于16位", true)
		return
	}
	if UserMemberModel.Api_update_password(uid, text) {
		api.Sendprivatemsg(bot, uid, "您的密码已被修改为：【"+text+"】", false)
	} else {
		LogErrorModel.Api_insert("修改密码错误", tuuz.FUNCTION_ALL())
		api.Sendprivatemsg(bot, uid, app_default.Default_error_alert, true)
	}
}
