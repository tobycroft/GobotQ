package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/BotRequestModel"
	"main.go/app/bot/model/SystemParamModel"
	"main.go/app/v1/user/action/BalanceAction"
	"main.go/common/BaseController"
	"main.go/tuuz"
	"main.go/tuuz/Date"

	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func BotController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("info", bot_info)
	route.Any("add", bot_add)
	route.Any("list", bot_list)
	route.Any("del", bot_delete)
}

func bot_info(c *gin.Context) {
	uid := c.GetHeader("uid")
	bot, ok := Input.PostInt64("self_id", c)
	if !ok {
		return
	}
	botinfo := BotModel.Api_find_byOwnerandBot(uid, bot)
	if len(botinfo) > 0 {
		RET.Success(c, 0, botinfo, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func bot_add(c *gin.Context) {
	uid := c.GetHeader("uid")
	bot, ok := Input.PostInt("bot", c)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	secret, ok := Input.Post("secret", c, false)
	if !ok {
		return
	}
	month, ok := Input.PostInt("month", c)
	if !ok {
		return
	}
	if len(BotRequestModel.Api_find(bot)) > 0 {
		RET.Fail(c, 406, nil, "本账号已经启用了，您对您的机器人发送acfur绑定+绑定密码来获取机器人的控制权")
		return
	}
	if len(BotRequestModel.Api_find(bot)) > 0 {
		RET.Fail(c, 406, nil, "本账号已经被提交过了")
		return
	}
	if len(BotRequestModel.Api_select_byUid(uid)) >= 3 {
		RET.Fail(c, 406, nil, "你的待通过列表已经有3个账号了，请先等待通过后才可以继续提交")
		return
	}
	price := SystemParamModel.Api_value("price")
	var br BotRequestModel.Interface
	db := tuuz.Db()
	db.Begin()
	br.Db = db
	end_date := Date.Date_offset_month_todayWithTimeZero(month)
	if br.Api_insert(uid, bot, password, uid, secret, end_date) {
		var ba BalanceAction.Interface
		ba.Db = db
		err := ba.App_single_balance(uid, nil, -float64(month)*Calc.Any2Float64(price), "预定了"+Calc.Any2String("month")+"月的服务")
		if err != nil {
			db.Rollback()
			RET.Fail(c, 400, err.Error(), err.Error())
		} else {
			db.Commit()
			RET.Success(c, 0, nil, nil)
		}
	} else {
		db.Rollback()
		RET.Fail(c, 500, nil, nil)
	}
}

func bot_list(c *gin.Context) {
	uid := c.GetHeader("uid")
	data := BotRequestModel.Api_select_byUid(uid)
	RET.Success(c, 0, data, nil)
}

func bot_delete(c *gin.Context) {
	uid := c.GetHeader("uid")
	bot, ok := Input.PostInt("bot", c)
	if !ok {
		return
	}
	var br BotRequestModel.Interface
	br.Db = tuuz.Db()
	if br.Api_delete_byUid(uid, bot) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}
