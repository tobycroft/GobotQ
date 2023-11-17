package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func MemberController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("list", member_list)
}

func member_list(c *gin.Context) {
	role, ok := Input.PostArray[any]("role", c)
	if !ok {
		return
	}
	user_id, ok := Input.PostInt64("user_id", c)
	if !ok {
		return
	}
	datas := GroupMemberModel.Api_select_byUid(user_id, role)
	RET.Success(c, 0, datas, nil)
}
