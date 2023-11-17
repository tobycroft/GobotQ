package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/action/GroupListAction"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func MemberController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("list", member_list)
	route.Any("bot", member_bot)
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
	for _, data := range datas {
		data["group_info"] = GroupListAction.App_find_groupList(data["self_id"], data["group_id"])
	}
	RET.Success(c, 0, datas, nil)
}

func member_bot(c *gin.Context) {
	uid := c.GetHeader("uid")
	bots := BotModel.Api_select_byOwner(uid)
	user_ids := []any{}
	for _, data := range bots {
		user_ids = append(user_ids, data["self_id"])
	}
	datas := GroupMemberModel.Api_select_inUids(user_ids, nil)
	RET.Success(c, 0, datas, nil)
}
