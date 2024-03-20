package Private

import (
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/common/BaseModel/TokenModel"
)

func App_userClearLogin(self_id, user_id, group_id int64) {
	if TokenModel.Api_delete(user_id) {
		msg := MessageBuilder.IMessageBuilder{}.New().Text("您的登录状态已经全部清空，如需再次登录请发送acfur登录")
		iapi.Api.SendPrivateMsg(self_id, user_id, group_id, msg, false)
	}
}
