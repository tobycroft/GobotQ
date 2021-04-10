package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupAutoReplyModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func AutoreplyController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckGroupAdmin(), gin.Recovery())
	route.Any("list", autoreply_list)
	route.Any("add", autoreply_add)
	route.Any("delete", autoreply_delete)

	route.Any("full_list", autoreply_full_list)
}

func autoreply_list(c *gin.Context) {
	gid := c.PostForm("gid")
	data := GroupAutoReplyModel.Api_select_semi(gid)
	RET.Success(c, 0, data, nil)
}

func autoreply_add(c *gin.Context) {
	gid := c.PostForm("gid")
	uid := c.PostForm("uid")
	Type, ok := Input.Post("type", c, false)
	if !ok {
		return
	}
	key, ok := Input.Post("key", c, false)
	if !ok {
		return
	}
	value, ok := Input.Post("value", c, false)
	if !ok {
		return
	}
	percent, ok := Input.PostInt("percent", c)
	if !ok {
		return
	}
	switch Type {
	case "semi":
		break
	case "full":
		break
	default:
		RET.Fail(c, 400, nil, "模式不正确")
		return
	}
	if GroupAutoReplyModel.Api_insert(Type, gid, uid, key, value, percent) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func autoreply_delete(c *gin.Context) {
	gid := c.PostForm("gid")
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	if GroupAutoReplyModel.Api_delete(gid, id) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func autoreply_full_list(c *gin.Context) {
	gid := c.PostForm("gid")
	data := GroupAutoReplyModel.Api_select_full(gid)
	RET.Success(c, 0, data, nil)
}
