package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func AutoreplyController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

}

func autoreply_add(c *gin.Context) {

}

func autoreply_list(c *gin.Context) {

}

func autoreply_delete(c *gin.Context) {

}

func autoreply_add_full(c *gin.Context) {

}
