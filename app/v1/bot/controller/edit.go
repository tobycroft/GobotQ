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
	route.Use(BaseController.CheckBotPower(), gin.Recovery())

	route.Any("img", edit_change_img)
	route.Any("clear_owner", edit_clear_owner)
	route.Any("secret", edit_change_secret)
	route.Any("password", edit_change_password)
	route.Any("password", edit_change_password)
}

func edit_change_img(c *gin.Context) {
	uid := c.PostForm("uid")
	bot := c.PostForm("self_id")
	img, ok := Input.Post("img", c, true)
	if !ok {
		return
	}
	if BotModel.Api_update_img(uid, bot, img) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_clear_owner(c *gin.Context) {
	bot := c.PostForm("self_id")
	if BotModel.Api_update_owner(bot, 0) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_change_secret(c *gin.Context) {
	bot := c.PostForm("self_id")
	if BotModel.Api_update_secret(bot, 0) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_change_password(c *gin.Context) {
	bot := c.PostForm("self_id")
	password, ok := Input.Post("password", c, false)
	if ok {
		return
	}
	if BotModel.Api_update_password(bot, password) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_change_url(c *gin.Context) {
	self_id := c.PostForm("self_id")
	url, ok := Input.Post("url", c, false)
	if !ok {
		return
	}
	if BotModel.Api_update_url(self_id, url) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
