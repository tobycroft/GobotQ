package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("user_info", user_info)
}

func user_info(c *gin.Context) {
	uid := c.PostForm("uid")
	user := UserMemberModel.Api_find(uid)
	if len(user) > 0 {
		RET.Success(c, 0, user, nil)
	} else {
		RET.Fail(c, 404, user, "没有找到数据")
	}
}
