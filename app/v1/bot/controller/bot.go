package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BotController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", bot_add)
	route.Any("list", bot_list)
}

func bot_add(c *gin.Context) {
	uid := c.PostForm("uid")
	bot, ok := Input.PostInt("bot", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	secret, ok := Input.Post("secret", c, false)
	if !ok {
		return
	}
	month, ok := Input.PostInt("month", c)
	if !ok {
		return
	}
	if len(BotRequestModel.Api_find(bot)) > 0 {
		RET.Fail(c, 406, nil, "本账号已经启用了，您对您的机器人发送acfur绑定+绑定密码来获取机器人的控制权")
		return
	}
	if len(BotRequestModel.Api_find(bot)) > 0 {
		RET.Fail(c, 406, nil, "本账号已经被提交过了")
		return
	}
	if len(BotRequestModel.Api_select_byUid(uid)) > 3 {
		RET.Fail(c, 406, nil, "你的待通过列表已经有3个账号了，请先等待通过后才可以继续提交")
		return
	}
	var br BotRequestModel.Interface
	br.Db = tuuz.Db()
	if br.Api_insert(uid, bot, password, uid, secret, month*3600*30) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_list(c *gin.Context) {
	uid := c.PostForm("uid")
	data := BotRequestModel.Api_select_byUid(uid)
	RET.Success(c, 0, data, nil)
}
