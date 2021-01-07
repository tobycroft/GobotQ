package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/app/bot/model/UserTokenModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController())

	route.Any("login", login)
}

func login(c *gin.Context) {
	qq, ok := Input.PostInt("qq", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	if len(UserMemberModel.Api_find_byQqandPassword(qq, password)) > 0 {
		token := Calc.GenerateToken()
		UserTokenModel.Api_insert(qq, token, c.ClientIP())
		RET.Success(c, 0, map[string]interface{}{
			"uid":   qq,
			"token": token,
		}, "登录成功")
	} else {
		RET.Fail(c, 0, nil, "登录信息不存在")
	}
}
