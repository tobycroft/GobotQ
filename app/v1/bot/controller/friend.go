package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotFriendAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/FriendListModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func FriendController(route *gin.RouterGroup) {
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

	route.Any("white_list", bot_white_friend_list)
	route.Any("white_add", bot_white_friend_add)
	route.Any("white_delete", bot_white_friend_delete)

	route.Any("friend_list", bot_friend_list)
	route.Any("friend_add", bot_friend_add)
	route.Any("friend_delete", bot_friend_delete)
}

func bot_white_friend_list(c *gin.Context) {
	bot := c.PostForm("bot")
	data := BotFriendAllowModel.Api_select(bot)

	for k, v := range data {
		user_info := FriendListModel.Api_find(v["uid"])
		if len(user_info) > 0 {
			data[k]["user_info"] = user_info
		} else {
			ui, err := api.Queryuserinfo(bot, v["uid"])
			if err != nil {

			} else {
				FriendListModel.Api_insert(bot, v["uid"], ui.NickName, ui.Remark, ui.Email)
				data[k]["user_info"] = FriendListModel.Api_find(v["uid"])
			}
		}
	}
	RET.Success(c, 0, data, nil)
}

func bot_white_friend_add(c *gin.Context) {
	bot := c.PostForm("bot")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	data := BotFriendAllowModel.Api_insert(bot, qq)
	if data {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_white_friend_delete(c *gin.Context) {
	bot := c.PostForm("bot")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	data := BotFriendAllowModel.Api_delete(bot, qq)
	if data {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_friend_list(c *gin.Context) {
	bot := c.PostForm("bot")
	data := FriendListModel.Api_select(bot)
	RET.Success(c, 0, data, nil)
}

func bot_friend_add(c *gin.Context) {
	bot := c.PostForm("bot")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	text, ok := Input.Post("text", c, false)
	if !ok {
		return
	}
	_, ret, err := api.Addfriend(bot, qq, text, nil)
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, ret.Retmsg, ret.Retmsg)
	}
}

func bot_friend_delete(c *gin.Context) {
	bot := c.PostForm("bot")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	_, ret, err := api.Deletefriend(bot, qq)
	if ret.Retcode == 0 {
		if FriendListModel.Api_delete_byUid(bot, qq) {
			RET.Success(c, 0, ret.Retmsg, ret.Retmsg)
		} else {
			RET.Fail(c, 500, nil, nil)
		}
	} else {
		RET.Fail(c, 200, err.Error(), err.Error())
	}
}
