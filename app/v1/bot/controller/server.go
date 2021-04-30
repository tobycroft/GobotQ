package controller

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/common/BaseController"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/Net"
	"main.go/tuuz/RET"
	"strings"
)

func ServerController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("add", server_add)

}

func server_add(c *gin.Context) {
	uid := c.PostForm("uid")
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	if strings.Contains(address, "http") {
		RET.Fail(c, 400, nil, "请直接填写您服务器的IP即可，无需在前面添加http，请保持")
	}
	port, ok := Input.PostInt64("port", c)
	if !ok {
		return
	}
	secret, ok := Input.Post("secret", c, false)
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
	//ret, err := Net.Post("docker.tuuz.cc:5701/get_status", nil, nil, nil, nil)
	ret, err := Net.Post(address+":"+Calc.Any2String(port)+"/get_login_info", nil, nil, nil, nil)
	if err != nil {
		RET.Fail(c, 300, nil, "无法访问远程服务器，请确认您的机器人接口已经对外开放，请稍后再试")
		return
	} else {
		var ret_struct get_login_info
		jsoniter.UnmarshalFromString(ret, &ret_struct)
		if ret_struct.Retcode != 0 {
			RET.Fail(c, 202, nil, "您的机器人没有准备好，请先登录并按照提示操作后再使用APP绑定")
			return
		} else {
			//todo:完成后执行动作
			BotModel.Api_insert(ret_struct.Data.UserID, ret_struct.Data.Nickname, "remote", uid, secret, "", 1672502399, "http://"+address+Calc.Any2String(":port"))
		}
	}
}

type get_login_info struct {
	Data struct {
		Nickname string `json:"nickname"`
		UserID   int    `json:"user_id"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

func update(c *gin.Context) {

}
