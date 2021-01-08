package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func OperationController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("online", operation_online)
	route.Any("offline", operation_offline)

}

func operation_online(c *gin.Context) {

}

func operation_offline(c *gin.Context) {

}
