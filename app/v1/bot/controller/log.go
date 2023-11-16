package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func LogController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("common", log_list_common)
	route.Any("all", log_list_all)
	route.Any("get", log_get)
}

func log_list_common(c *gin.Context) {
	//self_id := c.GetHeader("uid")
	//page, ok := Input.PostInt("page", c)
	//if !ok {
	//	return
	//}
	//limit, ok := Input.PostInt("limit", c)
	//if !ok {
	//	return
	//}
	//datas := BotLogModel.Api_select(self_id, page, limit)

}

func log_list_all(c *gin.Context) {

}

func log_get(c *gin.Context) {

}
