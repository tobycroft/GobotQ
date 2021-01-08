package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/app/v1/bot/model/BotAdminModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func ListController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(func(c *gin.Context) {
		uid := c.PostForm("uid")
		bot, ok := Input.PostInt("bot", c)
		if !ok {
			return
		}
		data := BotModel.Api_find_byOwnerandBot(uid, bot)
		if len(data) > 0 {
			c.Next()
			return
		} else {
			RET.Fail(c, 403, nil, "你并不拥有这个机器人")
			c.Abort()
			return
		}
	})
	route.Any("unbind", list_unbind)
	route.Any("owned", list_yours_own)
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

func list_your_control(c *gin.Context) {
	uid := c.PostForm("uid")
	bots := BotAdminModel.Api_select(uid)
	RET.Success(c, 0, bot, nil)
}
