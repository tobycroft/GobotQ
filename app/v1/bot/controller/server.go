package controller

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/api"
	"main.go/common/BaseController"
	"main.go/tuuz/Input"
	"main.go/tuuz/Net"
	"main.go/tuuz/RET"
)

func ServerController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", server_add)

}

func server_add(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	port, ok := Input.PostInt64("port", c)
	if !ok {
		return
	}
	secret, ok := Input.Post("secret", c, false)
	if !ok {
		return
	}
	qq, ok := Input.PostInt64("qq", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	if port <= 0 || port > 65535 {
		RET.Fail(c, 400, nil, nil)
		return
	}
	if len(secret) < 6 || len(secret) > 16 {
		RET.Fail(c, 400, nil, "secret应该大于6位小于16位")
		return
	}
	ret, err := Net.Post("docker.tuuz.cc:5701/get_status", nil, nil, nil, nil)
	if err != nil {
		RET.Fail(c, 300, nil, "无法访问远程服务器，请确认您的机器人接口已经对外开放，请稍后再试")
		return
	} else {
		var ret_struct api.DefaultRetStruct
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		json.UnmarshalFromString(ret, &ret_struct)
		if ret_struct.Retcode != 0 {
			RET.Fail(c, 200, ret_struct.Wording, "您的机器人没有准备好，请先登录并按照提示操作后再使用APP绑定")
			return
		} else {
			//todo:完成后执行动作
		}
	}
}

func update(c *gin.Context) {

}
