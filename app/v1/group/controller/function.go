package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func FunctionController(route *gin.RouterGroup) {
	route.Any("get", function_get)
	route.Any("detail", function_detail)

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("edit", function_edit)
}

func function_get(c *gin.Context) {
	group_id, ok := Input.PostInt64("group_id", c)
	if !ok {
		return
	}
	data := GroupFunctionModel.Api_find(group_id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func function_detail(c *gin.Context) {
	datas := GroupFunctionDetailModel.Api_select()
	RET.Success(c, 0, datas, nil)
}

func function_edit(c *gin.Context) {
	uid := c.GetHeader("uid")
	group_id, ok := Input.PostInt64("group_id", c)
	if !ok {
		return
	}
	data := GroupMemberModel.Api_find_byRoles(group_id, uid, []any{"admin", "owner"})
	if len(data) < 1 {
		RET.Fail(c, 403, nil, "你不是管理员，没有权限修改群设定")
		return
	}

	mp := Input.NewModelPost(c)
	mp.PostBool("sign")
	mp.PostBool("sign_send_private")
	mp.PostBool("sign_send_retract")
	mp.PostBool("adblock")
	mp.PostBool("ad_retract")
	mp.PostBool("all_send_private")
	mp.PostBool("ban_retract")
	mp.PostBool("ban_group")
	mp.PostBool("ban_wx")
	mp.PostInt64("ban_time")
	mp.PostInt64("ban_limit")
	mp.PostBool("auto_kick_out")
	mp.PostString("ban_words")
	mp.PostString("kick_words")
	mp.PostBool("ban_url")
	mp.PostBool("auto_welcome")
	mp.PostString("welcome_word")
	mp.PostBool("welcome_at")
	mp.PostBool("ban_share")
	mp.PostInt64("word_limit")
	mp.PostBool("auto_retract")
	mp.PostBool("auto_join")
	mp.PostBool("auto_verify")
	mp.PostBool("auto_hold")
	mp.PostBool("join_alert")
	mp.PostBool("exit_alert")
	mp.PostBool("ban_repeat")
	mp.PostInt64("repeat_time")
	mp.PostInt64("repeat_count")
	mp.PostBool("kick_to_black")
	mp.PostBool("exit_to_black")
	mp.PostBool("auto_card")
	mp.PostString("auto_card_value")
	mp.PostString("auto_verify_word")
	mp.PostBool("auto_card_insert")
	err, errs := mp.Error()
	if err != nil {
		RET.Fail(c, 0, errs, err.Error())
		return
	}
	data := mp.Select()
	if GroupFunctionModel.Api_update_more(group_id, data) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
