package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotSettingModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func SettingController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Use(BaseController.CheckBotPower(), gin.Recovery())

	route.Any("get", setting_get)
	route.Any("edit", setting_edit)
}

func setting_get(c *gin.Context) {
	self_id, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	data := BotSettingModel.Api_find(self_id)
	if len(data) == 0 {
		RET.Fail(c, 404, nil, nil)
	} else {
		RET.Success(c, 0, data, nil)
	}
}

func setting_edit(c *gin.Context) {
	self_id, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	mp := Input.NewModelPost(c)
	mp.PostInt64("add_friend")
	mp.PostInt64("add_group")
	data := mp.Select()
	if err, errs := mp.Error(); err != nil {
		RET.Fail(c, 400, errs, err.Error())
		return
	}
	if BotSettingModel.Api_update(self_id, data) {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
