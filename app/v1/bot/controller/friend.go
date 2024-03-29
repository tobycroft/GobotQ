package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotFriendAllowModel"
	"main.go/app/bot/model/FriendListModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func FriendController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckBotPower(), gin.Recovery())

	route.Any("white_list", bot_white_friend_list)
	route.Any("white_add", bot_white_friend_add)
	route.Any("white_delete", bot_white_friend_delete)

	route.Any("friend_list", bot_friend_list)
}

func bot_white_friend_list(c *gin.Context) {
	self_id := c.PostForm("self_id")
	datas := BotFriendAllowModel.Api_select(self_id)
	RET.Success(c, 0, datas, nil)
}

func bot_white_friend_add(c *gin.Context) {
	self_id := c.PostForm("self_id")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	data := BotFriendAllowModel.Api_insert(self_id, qq)
	if data {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_white_friend_delete(c *gin.Context) {
	self_id := c.PostForm("self_id")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	data := BotFriendAllowModel.Api_delete(self_id, qq)
	if data {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_friend_list(c *gin.Context) {
	self_id := c.PostForm("self_id")
	data := FriendListModel.Api_select(self_id)
	RET.Success(c, 0, data, nil)
}
