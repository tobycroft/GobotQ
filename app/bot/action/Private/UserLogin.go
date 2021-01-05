package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/UserTokenModel"
	"main.go/tuuz/Calc"
)

func UserLogin(bot int, uid int, text string) {
	rand := Calc.Rand(10000000, 99999999)
	token := Calc.GenerateToken()
	api.Sendprivatemsg(bot, uid, "您的登录密码：\r\n"+Calc.Int2String(rand))
	UserTokenModel.Api_insert(uid, token, "")
}
