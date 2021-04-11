package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BanController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckGroupAdmin(), gin.Recovery())

	route.Any("list", ban_group_list)
	route.Any("add", ban_group_add)
	route.Any("delete", ban_group_delete)

}

func ban_group_list(c *gin.Context) {
	gid := c.PostForm("group_id")
	data := GroupBanPermenentModel.Api_select_byGroupId(gid)
	for i, datum := range data {
		datum["user_info"] = GroupMemberModel.Api_find(datum["group_id"], datum["user_id"])
		data[i] = datum
	}
	RET.Success(c, 0, data, nil)
}

func ban_group_add(c *gin.Context) {
	gid := c.PostForm("group_id")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	if GroupBanPermenentModel.Api_delete(gid, qq) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func ban_group_delete(c *gin.Context) {
	gid := c.PostForm("group_id")
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	if GroupBanPermenentModel.Api_delete(gid, qq) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
