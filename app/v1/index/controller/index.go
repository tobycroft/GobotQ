package controller

import (
	"github.com/gin-gonic/gin"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("/", index)
}

func index(c *gin.Context) {
	c.String(0, "index")
}
