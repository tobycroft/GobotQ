package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/group/controller"
)

func GroupRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.EditController(route.Group("edit"))
	controller.ListController(route.Group("list"))
	controller.BlackController(route.Group("black"))
	controller.AutoreplyController(route.Group("autoreply"))
	controller.AutosendController(route.Group("autosend"))
	controller.BanController(route.Group("ban"))
	controller.WordController(route.Group("word"))

	controller.MemberController(route.Group("member"))
	controller.InfoController(route.Group("info"))
}
