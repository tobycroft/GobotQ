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

}
