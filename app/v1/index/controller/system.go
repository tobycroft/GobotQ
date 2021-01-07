package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/SystemAnnounceModel"
	"main.go/tuuz/RET"
)

func SystemController(route *gin.RouterGroup) {
	route.Any("/", index)

	route.Any("announcement", announcement)
}

func announcement(c *gin.Context) {
	data := SystemAnnounceModel.Api_select()
	RET.Success(c, 0, data, nil)
}
