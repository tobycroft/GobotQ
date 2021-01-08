package controller

import "github.com/gin-gonic/gin"

func RequestController(route *gin.RouterGroup) {

	route.Any("join", request_join)
	route.Any("out", request_out)
}

func request_join(c *gin.Context) {

}

func request_out(c *gin.Context) {

}
