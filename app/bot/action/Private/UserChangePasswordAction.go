package Private

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/config/app_default"
	"main.go/tuuz"
)

func App_userChangePassword(self_id, user_id, group_id int64, message string) {
	if len(message) < 1 {
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, "密码长度应该大于1位", true)
		return
	}
	if len(message) > 16 {
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, "密码长度应该小于等于16位", true)
		return
	}
	if UserMemberModel.Api_update_password(user_id, message) {
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, "您的密码已被修改为：【"+message+"】", false)
	} else {
		LogErrorModel.Api_insert("修改密码错误", tuuz.FUNCTION_ALL())
		iapi.Api.Sendprivatemsg(self_id, user_id, group_id, app_default.Default_error_alert, true)
	}
}
