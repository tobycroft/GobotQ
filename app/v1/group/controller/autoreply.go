package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func AutoreplyController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("list", autoreply_list)
	route.Any("add", autoreply_add)
	route.Any("delete", autoreply_delete)

	route.Any("full_list", autoreply_full_list)
	route.Any("full_add", autoreply_full_add)
	route.Any("full_delete", autoreply_full_delete)
}

func autoreply_list(c *gin.Context) {

}

func autoreply_add(c *gin.Context) {

}

func autoreply_delete(c *gin.Context) {

}

func autoreply_full_list(c *gin.Context) {

}

func autoreply_full_add(c *gin.Context) {

}

func autoreply_full_delete(c *gin.Context) {

}
