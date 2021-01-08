package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func ListController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController(), gin.Recovery())

	route.Any("list_unbind", list_unbind)
}

func list_unbind(c *gin.Context) {

}
