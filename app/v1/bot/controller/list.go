package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func ListController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("list_unbind", list_unbind)
}

func list_unbind(c *gin.Context) {
	unbinds := BotModel.Api_select_byOwner(0)
	RET.Success(c, 0, unbinds, nil)
}

func list_yours(c *gin.Context) {
	uid := c.PostForm("uid")
	bots := BotModel.Api_select_byOwner(uid)
	RET.Success(c, 0, bots, nil)
}
