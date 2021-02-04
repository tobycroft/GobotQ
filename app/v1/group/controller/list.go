package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func ListController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController())

	route.Any("control", group_control)
	route.Any("joined", group_control)
	route.Any("member", group_member)
}

func group_control(c *gin.Context) {
	uid := c.PostForm("uid")
	con_group := GroupMemberModel.Api_select_byUid(uid, []interface{}{"owner", "admin"})
	gids := []interface{}{}
	for _, data := range con_group {
		gids = append(gids, data["gid"])
	}
	gls := GroupListModel.Api_select_InGid(gids)
	RET.Success(c, 0, gls, nil)
}

func group_joined() {

}

func group_member(c *gin.Context) {
	uid := c.PostForm("uid")
	usergroup := GroupMemberModel.Api_select_byUid(uid, []interface{}{"owner", "admin"})
	RET.Success(c, 0, usergroup, nil)
}
