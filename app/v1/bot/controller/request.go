package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func RequestController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("join", request_join)
	route.Any("out", request_out)
	route.Any("allow", request_out)
}

func request_join(c *gin.Context) {

}

func request_out(c *gin.Context) {

}

func request_allow(c *gin.Context) {

}
