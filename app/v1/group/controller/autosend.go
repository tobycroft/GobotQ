package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupAutoReplyModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func AutosendController(route *gin.RouterGroup) {
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

}

func autosend_list(c *gin.Context) {
	gid := c.PostForm("gid")
	data := GroupAutoReplyModel.Api_select_semi(gid)
	RET.Success(c, 0, data, nil)
}
