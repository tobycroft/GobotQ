package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/app/v1/bot/model/BotAdminModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func ListController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("unbind", list_unbind)
	route.Any("owned", list_yours_own)
	route.Any("admin", list_your_admin_group)
}

func list_unbind(c *gin.Context) {
	unbinds := BotModel.Api_select_byOwner(0)
	RET.Success(c, 0, unbinds, nil)
}

func list_yours_own(c *gin.Context) {
	uid := c.PostForm("uid")
	bots := BotModel.Api_select_byOwner(uid)
	RET.Success(c, 0, bots, nil)
}

func list_your_admin_group(c *gin.Context) {
	uid := c.PostForm("uid")
	bots := BotAdminModel.Api_select(uid)
	RET.Success(c, 0, bots, nil)
}
