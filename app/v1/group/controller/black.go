package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BlackController(route *gin.RouterGroup) {
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

	route.Any("list", black_group_list)
	route.Any("add", black_group_add)
	route.Any("delete", black_group_delete)

}

func black_group_list(c *gin.Context) {
	gid := c.PostForm("gid")
	data := GroupBlackListModel.Api_select(gid)
	RET.Success(c, 0, data, nil)
}

func black_group_add(c *gin.Context) {

}

func black_group_delete(c *gin.Context) {
	gid := c.PostForm("gid")
	user_id, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	if GroupBlackListModel.Api_delete(gid, user_id) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
