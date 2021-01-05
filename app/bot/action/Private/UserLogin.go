package Private

import (
	"fmt"
	"main.go/app/bot/api"
	"main.go/tuuz/Calc"
)

func UserLogin(bot int, uid int, text string) {
	rand := Calc.Rand(10000000, 99999999)
	token := Calc.GenerateToken()
	api.Sendprivatemsg(bot, uid, "您的登录密码：\r\n"+Calc.Int2String(rand))

	fmt.Println(token)
}
