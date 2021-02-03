package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/bot/model/UserMemberModel"
	"main.go/app/v1/user/model/UserBalanceModel"
	"main.go/app/v1/user/model/UserBalanceRecordModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())
	route.Any("info", user_info)
	route.Any("balance", user_balance)
	route.Any("balance_record", user_balance_record)
}

func user_info(c *gin.Context) {
	uid := c.PostForm("uid")
	user := UserMemberModel.Api_find(uid)
	if len(user) > 0 {
		delete(user, "password")
		RET.Success(c, 0, user, nil)
	} else {
		RET.Fail(c, 404, user, "没有找到数据")
	}
}

func user_balance(c *gin.Context) {
	uid := c.PostForm("uid")
	ub := UserBalanceModel.Api_find(uid)
	if len(ub) > 0 {
		RET.Success(c, 0, ub, nil)
	} else {
		if UserBalanceModel.Api_insert(uid, 0) {
			ub = UserBalanceModel.Api_find(uid)
			RET.Success(c, 0, ub, nil)
		} else {
			RET.Fail(c, 500, nil, "数据错误")
		}
	}
}

func user_balance_record(c *gin.Context) {
	uid := c.PostForm("uid")
	balances := UserBalanceRecordModel.Api_select(uid)
	RET.Success(c, 0, balances, nil)
}
