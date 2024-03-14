package Private

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
)

func App_userLogin(self_id int64, user_id, group_id int64, message string) {
	rand := Calc.Rand(10000000, 99999999)
	uname := ""
	nickname, err := iapi.Api.GetStrangerInfo(self_id, user_id, false)
	if err != nil {
		uname = Calc.Int642String(user_id)
	} else {
		uname = nickname.Nickname
	}
	usermember := UserMemberModel.Api_find(user_id)
	if len(usermember) > 0 {
		if UserMemberModel.Api_update_all(user_id, uname, rand) {
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text("您的登录密码：\r\n" + Calc.Int2String(rand))
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		} else {
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text(app_default.Default_error_alert)
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		}
	} else {
		if UserMemberModel.Api_insert(user_id, uname, rand) {
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text("↓↓↓↓↓您的登录密码↓↓↓↓↓\r\n" + Calc.Int2String(rand) + "\r\n↑↑↑↑↑请在APP中输入↑↑↑↑↑\r\n" + app_default.Default_str_login_text)
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		} else {
			msg := MessageBuilder.IMessageBuilder{}.New().New().Text(app_default.Default_error_alert)
			iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
		}
	}

}
