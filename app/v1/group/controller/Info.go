package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupListModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
}

func info_list(c *gin.Context) {
	bot := c.PostForm("self_id")
	data := GroupListModel.Api_select(bot)
	RET.Success(c, 0, data, nil)
}
