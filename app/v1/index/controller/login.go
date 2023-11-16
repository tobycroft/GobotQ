package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/TokenModel"

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
	user := UserMemberModel.Api_find_byQqandPassword(qq, password)
	if len(user) > 0 {
		token := Calc.GenerateToken()
		TokenModel.Api_insert(qq, token, c.ClientIP())
		RET.Success(c, 0, map[string]any{
			"uid":   qq,
			"token": token,
			"uname": user["uname"],
		}, "登录成功")
	} else {
		RET.Fail(c, -1, nil, "登录信息不存在")
	}
}
