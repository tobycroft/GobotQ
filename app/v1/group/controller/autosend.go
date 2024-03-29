package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/group/action/AutoSendAction"
	"main.go/app/v1/group/model/AutoSendModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"time"
)

func AutosendController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckGroupAdmin(), gin.Recovery())

	route.Any("list", autosend_list)
	route.Any("get", autosend_get)
	route.Any("add", autosend_add)
	route.Any("update", autosend_update)
	route.Any("delete", autosend_delete)
}

func autosend_list(c *gin.Context) {
	gid := c.PostForm("group_id")
	data := AutoSendModel.Api_select(gid)
	RET.Success(c, 0, data, nil)
}

func autosend_get(c *gin.Context) {
	gid := c.PostForm("group_id")
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	data := AutoSendModel.Api_find(gid, id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func autosend_delete(c *gin.Context) {
	gid := c.PostForm("group_id")
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	if AutoSendModel.Api_delete(gid, id) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func autosend_add(c *gin.Context) {
	gid := c.PostForm("group_id")
	uid := c.GetHeader("uid")
	ident, ok := Input.Post("ident", c, true)
	if !ok {
		return
	}
	msg, ok := Input.Post("msg", c, false)
	if !ok {
		return
	}
	Type, ok := Input.Post("type", c, true)
	if !ok {
		return
	}
	sep, ok := Input.PostInt64("sep", c)
	if !ok {
		return
	}
	count, ok := Input.PostInt("count", c)
	if !ok {
		return
	}
	retract, ok := Input.PostBool("retract", c)
	if !ok {
		return
	}
	err := AutoSendAction.App_autosend_verify(sep, count, Type)
	if err != nil {
		RET.Fail(c, 400, err.Error(), err.Error())
		return
	}
	next_time := int64(0)
	switch Type {
	case "sep":
		next_time = sep + time.Now().Unix()
		break

	case "fix":
		count = 1
		next_time = sep
		break

	default:
		Type = "sep"
		count = 1
		sep = 60
		break
	}

	if AutoSendModel.Api_insert(gid, uid, ident, msg, Type, sep, count, next_time, retract) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func autosend_update(c *gin.Context) {
	gid := c.PostForm("group_id")
	uid := c.GetHeader("uid")
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	ident, ok := Input.Post("ident", c, true)
	if !ok {
		return
	}
	msg, ok := Input.Post("msg", c, false)
	if !ok {
		return
	}
	Type, ok := Input.Post("type", c, true)
	if !ok {
		return
	}
	sep, ok := Input.PostInt64("sep", c)
	if !ok {
		return
	}
	count, ok := Input.PostInt("count", c)
	if !ok {
		return
	}
	retract, ok := Input.PostBool("retract", c)
	if !ok {
		return
	}
	next_time := int64(0)
	switch Type {
	case "sep":
		next_time = sep + time.Now().Unix()
		break

	case "fix":
		count = 1
		next_time = sep
		break

	default:
		Type = "sep"
		count = 1
		sep = 60
		break
	}

	if AutoSendModel.Api_update(gid, uid, id, ident, msg, Type, sep, count, next_time, retract) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
