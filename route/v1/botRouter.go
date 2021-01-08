package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/bot/controller"
)

func BotRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.BindController(route.Group("bind"))
	controller.EditController(route.Group("edit"))
	controller.OperationController(route.Group("operation"))
	controller.RequestController(route.Group("request"))
	controller.ShareController(route.Group("share"))
}
