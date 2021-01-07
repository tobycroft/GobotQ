package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/UserTokenModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
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
	token := Calc.GenerateToken()
	UserTokenModel.Api_insert(qq, token, "")
}
