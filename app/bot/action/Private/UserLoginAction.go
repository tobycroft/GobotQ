package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
	"main.go/tuuz/Calc"
)

func App_userLogin(self_id int64, uid, gid int64, message string) {
	rand := Calc.Rand(10000000, 99999999)
	uname := ""
	nickname, err := api.GetStrangerInfo(self_id, uid, false)
	if err != nil {
		uname = Calc.Int642String(uid)
	} else {
		uname = nickname.Nickname
	}
	usermember := UserMemberModel.Api_find(uid)
	if len(usermember) > 0 {
		if UserMemberModel.Api_update_all(uid, uname, rand) {
			api.Sendprivatemsg(self_id, uid, "您的登录密码：\r\n"+Calc.Int2String(rand), false)
		} else {
			api.Sendprivatemsg(self_id, uid, app_default.Default_error_alert, false)
		}
	} else {
		if UserMemberModel.Api_insert(uid, uname, rand) {
			api.Sendprivatemsg(self_id, uid, "↓↓↓↓↓您的登录密码↓↓↓↓↓\r\n"+Calc.Int2String(rand)+"\r\n↑↑↑↑↑请在APP中输入↑↑↑↑↑\r\n"+app_default.Default_str_login_text, false)
		} else {
			api.Sendprivatemsg(self_id, uid, app_default.Default_error_alert, false)
		}
	}

}
