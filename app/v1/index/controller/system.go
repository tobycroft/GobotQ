package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/SystemAnnounceModel"
	"main.go/app/v1/index/model/VersionModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func SystemController(route *gin.RouterGroup) {
	route.Any("/", index)

	route.Any("announcement", announcement)
	route.Any("version", version)
}

func announcement(c *gin.Context) {
	data := SystemAnnounceModel.Api_select()
	RET.Success(c, 0, data, nil)
}

func version(c *gin.Context) {
	platform, ok := Input.Post("platform", c, false)
	if !ok {
		return
	}
	version := VersionModel.Api_find(platform)
	if len(version) > 0 {
		RET.Success(c, 0, version, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
