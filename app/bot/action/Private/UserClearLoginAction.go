package Private

import (
	"main.go/app/bot/apipost"
	"main.go/app/bot/model/UserTokenModel"
)

func App_userClearLogin(self_id, user_id, group_id int64) {
	if UserTokenModel.Api_delete(user_id) {
		apipost.ApiPost{}.Sendprivatemsg(self_id, user_id, group_id, "您的登录状态已经全部清空，如需再次登录请发送acfur登录", false)
	}
}
