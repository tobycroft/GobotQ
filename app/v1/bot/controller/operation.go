package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func OperationController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.ChechBotPower(), gin.Recovery())

}
