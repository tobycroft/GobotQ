package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
	"main.go/tuuz/Calc"
)

func App_UserLogin(bot int, uid int, text string) {
	rand := Calc.Rand(10000000, 99999999)

	usermember := UserMemberModel.Api_find(uid)
	if len(usermember) > 0 {
		if !UserMemberModel.Api_update_password(uid, rand) {
			api.Sendprivatemsg(bot, uid, "您的登录密码：\r\n"+Calc.Int2String(rand))
			return
		}
	} else {
		if !UserMemberModel.Api_insert(uid, uid, rand) {
			api.Sendprivatemsg(bot, uid, "↓↓↓↓↓您的登录密码↓↓↓↓↓\r\n"+Calc.Int2String(rand)+"\r\n↑↑↑↑↑请在APP中输入↑↑↑↑↑\r\n"+app_default.Default_str_login_text)
		}
	}
}
