package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func GroupController(route *gin.RouterGroup) {
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

	route.Any("white_list", bot_white_group_list)
	route.Any("white_add", bot_white_group_add)
	route.Any("white_delete", bot_white_group_delete)

	route.Any("group_list", bot_group_list)
	route.Any("group_exit", bot_group_exit)
}

func bot_white_group_list(c *gin.Context) {
	bot := c.PostForm("bot")
	data := BotGroupAllowModel.Api_select(bot)
	for k, v := range data {
		groupinfo := GroupListModel.Api_find(v["gid"])
		if len(groupinfo) > 0 {
			data[k]["group_info"] = groupinfo
		}
	}
	RET.Success(c, 0, data, nil)
}

func bot_white_group_add(c *gin.Context) {
	bot := c.PostForm("bot")
	gid, ok := Input.PostInt64("gid", c)
	if !ok {
		return
	}
	data := BotGroupAllowModel.Api_insert(bot, gid)
	if data {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_white_group_delete(c *gin.Context) {
	bot := c.PostForm("bot")
	gid, ok := Input.PostInt64("gid", c)
	if !ok {
		return
	}
	data := BotGroupAllowModel.Api_delete(bot, gid)
	if data {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_group_list(c *gin.Context) {
	bot := c.PostForm("bot")
	data := GroupListModel.Api_select(bot)
	RET.Success(c, 0, data, nil)
}

func bot_group_exit(c *gin.Context) {
	bot := c.PostForm("bot")
	gid, ok := Input.PostInt64("gid", c)
	if !ok {
		return
	}
	ret := api.SetGroupLeave(bot, gid)
	if ret {
		if GroupListModel.Api_delete_byBotandGid(bot, gid) {
			RET.Success(c, 0, nil, nil)
		} else {
			RET.Fail(c, 500, nil, nil)
		}
	} else {
		RET.Fail(c, 200, nil, "故障")
	}
}
