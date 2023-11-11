package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func ListController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController())

	route.Any("control", group_control)
	route.Any("joined", group_joined)
	route.Any("member", group_member)
	route.Any("setting", group_setting_get)
}

func group_control(c *gin.Context) {
	uid := c.PostForm("uid")
	con_group := GroupMemberModel.Api_select_byUid(uid, []any{"owner", "admin"})
	gids := []any{}
	for _, data := range con_group {
		gids = append(gids, data["group_id"])
	}
	gls := []gorose.Data{}
	if len(gids) > 0 {
		gls = GroupListModel.Api_select_InGid(gids)
	}
	RET.Success(c, 0, gls, nil)
}

func group_joined(c *gin.Context) {
	uid := c.PostForm("uid")
	con_group := GroupMemberModel.Api_select_byUid(uid, []any{"member"})
	gids := []any{}
	for _, data := range con_group {
		gids = append(gids, data["group_id"])
	}
	gls := []gorose.Data{}
	if len(gids) > 0 {
		gls = GroupListModel.Api_select_InGid(gids)
	}
	RET.Success(c, 0, gls, nil)
}

func group_member(c *gin.Context) {
	uid := c.PostForm("uid")
	usergroup := GroupMemberModel.Api_select_byUid(uid, []any{"owner", "admin"})
	RET.Success(c, 0, usergroup, nil)
}

func group_setting_get(c *gin.Context) {
	uid := c.PostForm("uid")
	gid, ok := Input.PostInt64("group_id", c)
	if !ok {
		return
	}
	if len(GroupMemberModel.Api_find(gid, uid)) > 0 {
		group_setting := GroupFunctionModel.Api_find(gid)
		if len(group_setting) < 1 {
			RET.Fail(c, 404, nil, "本群不存在")
		} else {
			function := GroupFunctionDetailModel.Api_select_kv()
			arr := make(map[string]map[string]any)
			for k, v := range group_setting {
				if function[k] != nil {
					function[k]["value"] = v
					arr[k] = function[k]
				}
			}
			RET.Success(c, 0, arr, nil)
		}
	} else {
		RET.Fail(c, 403, nil, "你不在这个群")
	}

}
