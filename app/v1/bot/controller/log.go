package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func LogController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("common", log_common_list)
}

func log_common_list(c *gin.Context) {

}

func log_all(c *gin.Context) {

}

func log_list(c *gin.Context) {

}
