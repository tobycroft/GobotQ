package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("get", info_get)
	route.Any("unbind", info_unbind)
	route.Any("owned", info_yours_own)
}

func info_unbind(c *gin.Context) {
	unbinds := BotModel.Api_select_byOwner(0)
	RET.Success(c, 0, unbinds, nil)
}

func info_yours_own(c *gin.Context) {
	uid := c.GetHeader("uid")
	bots := BotModel.Api_select_byOwner(uid)
	RET.Success(c, 0, bots, nil)
}

func info_get(c *gin.Context) {
	uid := c.GetHeader("uid")
	bot, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	botinfo := BotModel.Api_find_byOwnerandBot(uid, bot)
	if len(botinfo) > 0 {
		RET.Success(c, 0, botinfo, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
