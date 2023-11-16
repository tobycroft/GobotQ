package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"

	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BindController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", bind_bot_add)
}

func bind_bot_add(c *gin.Context) {
	uid := c.GetHeader("uid")
	bot, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	secret, ok := Input.Post("secret", c, false)
	if !ok {
		return
	}
	data := BotModel.Api_find(bot)
	if len(data) > 0 {
		if data["owner"].(int64) != 0 {
			RET.Fail(c, 407, "该机器人已经被绑定", "该机器人已经被绑定")
			return
		}
		if Calc.Any2String(data["secret"]) == secret {
			if BotModel.Api_update_owner(bot, uid) {
				RET.Success(c, 0, "绑定成功", "绑定成功")
			} else {
				RET.Fail(c, 500, "绑定失败", "绑定失败")
			}
		} else {
			RET.Fail(c, 400, "密钥错误", "密钥错误")
		}
	} else {
		RET.Fail(c, 404, "未找到这个机器人", "未找到这个机器人")
	}
}
