package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Input"
)

func BotController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

}

func add(c *gin.Context) {
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
	time, ok := Input.PostInt("time", c)
	var br BotRequestModel.Interface
	br.Db = tuuz.Db()
	if br.Api_insert(uid, bot, password, uid, secret, time*3600*30) {

	}
}
