package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/Net"
)

func ServerController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

}

func add(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}

	ret, err := Net.Post("docker.tuuz.cc:5701/get_status", nil, nil, nil, nil)
	if err != nil {

	} else {

	}
}

func update(c *gin.Context) {

}
