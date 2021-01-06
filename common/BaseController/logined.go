package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/UserTokenModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginedController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("S-P-I", c.ClientIP())
		c.Header("S-P-P", app_conf.Project)
		uid, ok := Input.Post("qq", c, false)
		if !ok {
			c.Abort()
			return
		}
		token, ok := Input.Post("token", c, false)
		if !ok {
			c.Abort()
			return
		}
		debug, ok := c.GetPostForm("debug")
		if ok {
			if debug == app_conf.Debug && app_conf.TestMode {
				c.Next()
				return
			}
		}
		if len(UserTokenModel.Api_find_byToken(uid, token)) > 0 {
			c.Next()
			return
		} else {
			RET.Fail(c, -1, "Auth_fail", "未登录")
			c.Abort()
			return
		}
	}
}
