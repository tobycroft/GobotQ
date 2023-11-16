package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func BotController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	//route.Any("bot_list", bot_list)
}

func bot_list(c *gin.Context) {
	uid := c.GetHeader("uid")
	bots := BotModel.Api_select_byOwner(uid)
	RET.Success(c, 0, bots, nil)
}
