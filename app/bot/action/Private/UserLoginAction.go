package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
	"main.go/tuuz/Calc"
)

func App_userLogin(bot int, uid, gid int, text string) {
	rand := Calc.Rand(10000000, 99999999)
	uname := ""
	nickname, err := api.Getnickname(bot, uid, true)
	if err != nil {
		uname = Calc.Int2String(uid)
	} else {
		uname = nickname.Ret
	}
	usermember := UserMemberModel.Api_find(uid)
	if len(usermember) > 0 {
		if UserMemberModel.Api_update_all(uid, uname, rand) {
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, "您的登录密码：\r\n"+Calc.Int2String(rand), false)
			} else {
				api.Sendprivatemsg(bot, uid, "您的登录密码：\r\n"+Calc.Int2String(rand), false)
			}
		} else {
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, app_default.Default_error_alert, false)
			} else {
				api.Sendprivatemsg(bot, uid, app_default.Default_error_alert, false)
			}
		}
	} else {
		if UserMemberModel.Api_insert(uid, uname, rand) {
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, "↓↓↓↓↓您的登录密码↓↓↓↓↓\r\n"+Calc.Int2String(rand)+"\r\n↑↑↑↑↑请在APP中输入↑↑↑↑↑\r\n"+app_default.Default_str_login_text, false)
			} else {
				api.Sendprivatemsg(bot, uid, "↓↓↓↓↓您的登录密码↓↓↓↓↓\r\n"+Calc.Int2String(rand)+"\r\n↑↑↑↑↑请在APP中输入↑↑↑↑↑\r\n"+app_default.Default_str_login_text, false)
			}
		} else {
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, app_default.Default_error_alert, false)
			} else {
				api.Sendprivatemsg(bot, uid, app_default.Default_error_alert, false)
			}
		}
	}

}
