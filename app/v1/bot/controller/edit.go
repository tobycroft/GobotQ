package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/tuuz/Input"
)

func EditController(route *gin.RouterGroup) {

	route.Any("change_img", change_img)
}

func change_img(c *gin.Context) {
	uid := c.PostForm("uid")
	img, ok := Input.Post("img", c, true)
	if !ok {
		return
	}

}
