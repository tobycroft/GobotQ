package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func EditController(route *gin.RouterGroup) {
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
	route.Any("change_img", change_img)
	route.Any("change_name", change_name)
	route.Any("clear_owner", clear_owner)
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

func change_name(c *gin.Context) {
	bot := c.PostForm("bot")
	nickname, ok := Input.Post("nickname", c, false)
	if !ok {
		return
	}
	os, err := api.Setnickname(bot, nickname)
	if err != nil {
		RET.Fail(c, 300, err, err.Error())
		return
	}
	if !os {
		RET.Fail(c, 400, nil, "api修改昵称失败")
		return
	}
	if BotModel.Api_update_cname(bot, nickname) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
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
