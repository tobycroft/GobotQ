package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupAutoReplyModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func AutoreplyController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(func(c *gin.Context) {
		uid := c.PostForm("uid")
		gid, ok := Input.PostInt64("gid", c)
		if !ok {
			return
		}
		data := GroupMemberModel.Api_find(gid, uid)
		if len(data) > 0 {
			if data["type"].(string) == "admin" || data["type"].(string) == "owner" {
				c.Next()
				return
			} else {
				RET.Fail(c, 403, nil, "你不是本群的管理员")
				c.Abort()
				return
			}
		} else {
			RET.Fail(c, 403, nil, "未找到该群，请检查机器人是否有加入本群")
			c.Abort()
			return
		}
	})
	route.Any("list", autoreply_list)
	route.Any("add", autoreply_add)
	route.Any("delete", autoreply_delete)

	route.Any("full_list", autoreply_full_list)
	route.Any("full_add", autoreply_full_add)
	route.Any("full_delete", autoreply_full_delete)
}

func autoreply_list(c *gin.Context) {
	gid := c.PostForm("gid")
	data := GroupAutoReplyModel.Api_select(gid)
	RET.Success(c, 0, data, nil)
}

func autoreply_add(c *gin.Context) {
	gid := c.PostForm("gid")
	uid := c.PostForm("uid")
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
	if GroupAutoReplyModel.Api_insert(gid, uid, key, value, percent) {
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

}

func autoreply_full_add(c *gin.Context) {

}

func autoreply_full_delete(c *gin.Context) {

}
