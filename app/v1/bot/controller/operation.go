package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/api"
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
		ok, err := api.Logoutqq(bot)
		if err != nil {
			RET.Fail(c, 300, err.Error(), err.Error())
		} else if ok {
			RET.Success(c, 0, nil, nil)
		} else {
			RET.Fail(c, 300, nil, "下线失败，机器人可能本来就没有在线？")
		}

	} else {
		RET.Fail(c, 403, nil, "你并不拥有这个机器人")
	}
}

func operation_offline(c *gin.Context) {
	uid := c.PostForm("uid")
	bot, ok := Input.PostInt("bot", c)
	if !ok {
		return
	}
	data := BotModel.Api_find_byOwnerandBot(uid, bot)
	if len(data) > 0 {
		ok, err := api.Loginqq(bot)
		if err != nil {
			RET.Fail(c, 300, err.Error(), err.Error())
		} else if ok {
			RET.Success(c, 0, nil, nil)
		} else {
			RET.Fail(c, 300, nil, "上线失败，机器人的密码可能已经被修改了？")
		}

	} else {
		RET.Fail(c, 403, nil, "你并不拥有这个机器人")
	}
}
