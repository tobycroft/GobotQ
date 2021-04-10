package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func EditController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckGroupAdmin(), gin.Recovery())

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
