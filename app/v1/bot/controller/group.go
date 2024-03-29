package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func GroupController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckBotPower(), gin.Recovery())

	route.Any("white_list", bot_white_group_list)
	route.Any("white_add", bot_white_group_add)
	route.Any("white_delete", bot_white_group_delete)

	route.Any("group_list", bot_group_list)
	route.Any("group_exit", bot_group_exit)
}

func bot_white_group_list(c *gin.Context) {
	bot := c.PostForm("self_id")
	data := BotGroupAllowModel.Api_select(bot)
	for k, v := range data {
		groupinfo := GroupListModel.Api_find(v["group_id"])
		if len(groupinfo) > 0 {
			data[k]["group_info"] = groupinfo
		}
	}
	RET.Success(c, 0, data, nil)
}

func bot_white_group_add(c *gin.Context) {
	bot := c.PostForm("self_id")
	gid, ok := Input.PostInt64("group_id", c)
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
	bot := c.PostForm("self_id")
	gid, ok := Input.PostInt64("group_id", c)
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
	bot := c.PostForm("self_id")
	data := GroupListModel.Api_select(bot)
	RET.Success(c, 0, data, nil)
}

func bot_group_exit(c *gin.Context) {
	bot, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	gid, ok := Input.PostInt64("group_id", c)
	if !ok {
		return
	}
	ret, _ := iapi.Api.SetGroupLeave(bot, gid)
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
