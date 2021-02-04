package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func EditController(route *gin.RouterGroup) {
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
				RET.Success(c, 0, data, nil)
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

	route.Any("setting", group_setting_set)
	route.Any("setting_get", group_setting_get)
}

func group_setting_set(c *gin.Context) {
	gid, ok := Input.PostInt64("gid", c)
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
	if GroupFunctionModel.Api_update(gid, key, value) {
		RET.Success(c, 0, nil, "修改成功")
	} else {
		RET.Fail(c, 500, nil, "数据库修改失败")
	}

}
