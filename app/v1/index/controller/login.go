package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func LoginController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController())

	route.Any("login", login)
}

func login() {
	qq
}
