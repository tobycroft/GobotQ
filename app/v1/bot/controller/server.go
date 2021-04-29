package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func ServerController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

}

func add(c *gin.Context) {

}

func update(c *gin.Context) {

}
