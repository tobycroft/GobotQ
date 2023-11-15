package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func CheckBotPower() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.PostForm("uid")
		bot, ok := Input.PostInt("self_id", c)
		if !ok {
			return
		}
		data := BotModel.Api_find_byOwnerandBot(uid, bot)
		if len(data) > 0 {
			c.Next()
			return
		} else {
			RET.Fail(c, 403, nil, "你并不拥有这个机器人")
			c.Abort()
			return
		}
	}
}

func CheckGroupAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.PostForm("uid")
		gid, ok := Input.PostInt64("group_id", c)
		if !ok {
			return
		}
		data := GroupMemberModel.Api_find(gid, uid)
		if len(data) > 0 {
			if data["role"].(string) == "admin" || data["role"].(string) == "owner" {
				c.Next()
				return
			} else {
				RET.Fail(c, 403, nil, "你不是本群的管理员")
				c.Abort()
				return
			}
		} else {
			RET.Fail(c, 403, nil, "未找到该群，请检查机器人是否有加入本群")
			c.Abort()
			return
		}
	}
}
