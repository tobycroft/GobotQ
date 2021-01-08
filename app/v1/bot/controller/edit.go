package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func EditController(route *gin.RouterGroup) {

	route.Any("change_img", change_img)
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
