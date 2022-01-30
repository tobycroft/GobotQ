package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/GroupBanWordModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func WordController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckGroupAdmin(), gin.Recovery())

	route.Any("list", word_list)
	route.Any("add", word_add)
	route.Any("delete", word_delete)

}

func word_list(c *gin.Context) {
	gid := c.PostForm("group_id")
	data := GroupBanWordModel.Api_select(gid)
	RET.Success(c, 0, data, nil)
}

func word_add(c *gin.Context) {
	gid := c.PostForm("group_id")
	uid, ok := Input.PostInt64("uid", c)
	if !ok {
		return
	}
	word, ok := Input.Post("word", c, false)
	if !ok {
		return
	}
	mode, ok := Input.PostInt("mode", c)
	if !ok {
		return
	}
	is_kick, ok := Input.PostBool("is_kick", c)
	if !ok {
		return
	}
	is_ban, ok := Input.PostBool("is_ban", c)
	if !ok {
		return
	}
	is_retract, ok := Input.PostBool("is_retract", c)
	if !ok {
		return
	}
	share, ok := Input.PostBool("share", c)
	if !ok {
		return
	}
	if GroupBanWordModel.Api_insert(gid, uid, word, mode, is_kick, is_ban, is_retract, share) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func word_delete(c *gin.Context) {
	gid := c.PostForm("group_id")
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	if GroupBanWordModel.Api_delete_byId(gid, id) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
