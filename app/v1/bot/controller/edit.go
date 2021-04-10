package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func EditController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.ChechBotPower(), gin.Recovery())

	route.Any("img", change_img)
	route.Any("clear_owner", clear_owner)
	route.Any("secret", change_secret)
	route.Any("password", change_password)
}

func change_img(c *gin.Context) {
	uid := c.PostForm("uid")
	bot := c.PostForm("bot")
	img, ok := Input.Post("img", c, true)
	if !ok {
		return
	}
	if BotModel.Api_update_img(uid, bot, img) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 0, nil, nil)
	}
}

func clear_owner(c *gin.Context) {
	bot := c.PostForm("bot")
	if BotModel.Api_update_owner(bot, 0) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 0, nil, nil)
	}
}

func change_secret(c *gin.Context) {
	bot := c.PostForm("bot")
	if BotModel.Api_update_secret(bot, 0) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 0, nil, nil)
	}
}

func change_password(c *gin.Context) {
	bot := c.PostForm("bot")
	password, ok := Input.Post("password", c, false)
	if ok {
		return
	}
	if BotModel.Api_update_password(bot, password) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 0, nil, nil)
	}
}
