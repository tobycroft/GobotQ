package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BlackController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckGroupAdmin(), gin.Recovery())

	route.Any("list", black_group_list)
	route.Any("add", black_group_add)
	route.Any("delete", black_group_delete)

}

func black_group_list(c *gin.Context) {
	gid := c.PostForm("group_id")
	data := GroupBlackListModel.Api_select(gid)
	RET.Success(c, 0, data, nil)
}

func black_group_add(c *gin.Context) {
	gid := c.PostForm("group_id")
	uid, ok := Input.PostInt64("uid", c)
	if !ok {
		return
	}
	if GroupBlackListModel.Api_delete(gid, uid) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func black_group_delete(c *gin.Context) {
	gid := c.PostForm("group_id")
	uid, ok := Input.PostInt64("uid", c)
	if !ok {
		return
	}
	if GroupBlackListModel.Api_delete(gid, uid) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
