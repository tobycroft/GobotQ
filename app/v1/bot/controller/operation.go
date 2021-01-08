package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func OperationController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("online", operation_online)
	route.Any("offline", operation_offline)

}

func operation_online(c *gin.Context) {
	uid := c.PostForm("uid")
	bot, ok := Input.PostInt("bot", c)
	if !ok {
		return
	}
	data := BotModel.Api_find_byOwnerandBot(uid, bot)
	if len(data) > 0 {

	} else {
		RET.Fail(c, 403, nil, "你并不拥有这个机器人")
	}

}

func operation_offline(c *gin.Context) {

}
