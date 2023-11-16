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
	route.Any("url", edit_change_url)
}

func edit_change_img(c *gin.Context) {
	uid := c.GetHeader("uid")
	self_id, ok := Input.PostInt("self_id", c)
	if !ok {
		return
	}
	img, ok := Input.Post("img", c, true)
	if !ok {
		return
	}
	if BotModel.Api_update_img(uid, self_id, img) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_clear_owner(c *gin.Context) {
	self_id, ok := Input.PostInt("self_id", c)
	if !ok {
		return
	}
	secret, ok := Input.Post("secret", c, false)
	if !ok {
		return
	}
	if BotModel.Api_update_owner(self_id, secret) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_change_secret(c *gin.Context) {
	self_id, ok := Input.PostInt("self_id", c)
	if !ok {
		return
	}
	if BotModel.Api_update_secret(self_id, 0) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_change_password(c *gin.Context) {
	self_id, ok := Input.PostInt("self_id", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if ok {
		return
	}
	if BotModel.Api_update_password(self_id, password) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func edit_change_url(c *gin.Context) {
	self_id, ok := Input.PostInt("self_id", c)
	if !ok {
		return
	}
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
