package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Any("list", info_list)
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("bind", info_bind)
	route.Any("get", info_get)
	route.Any("unbind", info_unbind)
	route.Any("owned", info_yours_own)
}

func info_list(c *gin.Context) {
	Type, ok := Input.PostIn("type", c, []string{"public", "share"})
	if !ok {
		return
	}
	datas := BotModel.Api_select_public(Type)
	RET.Success(c, 0, datas, nil)
}
func info_unbind(c *gin.Context) {
	unbinds := BotModel.Api_select_byOwner(0)
	RET.Success(c, 0, unbinds, nil)
}

func info_yours_own(c *gin.Context) {
	uid := c.GetHeader("uid")
	bots := BotModel.Api_select_byOwner(uid)
	RET.Success(c, 0, bots, nil)
}

func info_get(c *gin.Context) {
	uid := c.GetHeader("uid")
	bot, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	botinfo := BotModel.Api_find_byOwnerandBot(uid, bot)
	if len(botinfo) > 0 {
		RET.Success(c, 0, botinfo, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_bind(c *gin.Context) {
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
		if Calc.Any2Int64(data["owner"]) != 0 {
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
