package route

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/event"
	v1 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("", func(context *gin.Context) {
		data, _ := context.GetRawData()
		//fmt.Println(string(data))
		event.EventRouter(string(data))
		context.String(200, "ok")
	})
	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}
		user := version1.Group("user")
		{
			v1.UserRouter(user)
		}
		bot := version1.Group("bot")
		{
			v1.BotRouter(bot)
		}
		group := version1.Group("group")
		{
			v1.GroupRouter(group)
		}
	}
}
