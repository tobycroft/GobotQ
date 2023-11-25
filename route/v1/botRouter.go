package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/bot/controller"
)

func BotRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.EditController(route.Group("edit"))
	controller.InfoController(route.Group("info"))
	controller.RequestController(route.Group("request"))
	controller.GroupController(route.Group("group"))
	controller.LogController(route.Group("log"))
	controller.FriendController(route.Group("friend"))

	controller.SettingController(route.Group("setting"))
}
