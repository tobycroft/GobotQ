package event

import (
	"fmt"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/action/group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/GroupMsgModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/app/bot/model/groupAutoReplyModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"regexp"
	"strings"
	"sync"
)

type GM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN       int    `json:"UIN"`
		Card      string `json:"Card"`
		SpecTitle string `json:"SpecTitle"`
		Pos       struct {
			Lo int `json:"Lo"`
			La int `json:"La"`
		} `json:"Pos"`
	} `json:"FromQQ"`
	LogonQQ   int `json:"LogonQQ"`
	TimeStamp struct {
		Recv int `json:"Recv"`
		Send int `json:"Send"`
	} `json:"TimeStamp"`
	FromGroup struct {
		GIN  int    `json:"GIN"`
		Name string `json:"name"`
	} `json:"FromGroup"`
	Msg struct {
		Req       int    `json:"Req"`
		Random    int    `json:"Random"`
		SubType   int    `json:"SubType"`
		AppID     int    `json:"AppID"`
		Text      string `json:"Text"`
		TextReply string `json:"Text_Reply"`
		BubbleID  int    `json:"BubbleID"`
	} `json:"Msg"`
	File struct {
		ID   string `json:"ID"`
		MD5  string `json:"MD5"`
		Name string `json:"Name"`
		Size int64  `json:"Size"`
	} `json:"File"`
}

func GroupMsg(gm GM) {
	GroupMsgModel.Api_insert(gm.LogonQQ, gm.FromQQ.UIN, gm.FromGroup.GIN, gm.Msg.Text, gm.Msg.Req, gm.Msg.Random, gm.File.ID, gm.File.MD5,
		gm.File.Name, gm.File.Size)
	is_self := false

	text := gm.Msg.Text
	bot := gm.LogonQQ
	uid := gm.FromQQ.UIN
	gid := gm.FromGroup.GIN
	retract := gm.Msg.Random

	if gm.LogonQQ == gm.FromQQ.UIN {
		is_self = true
	}

	if !is_self {
		GroupHandle(bot, gid, uid, text, gm.Msg.Req, retract)
	}
}

func GroupHandle(bot, gid, uid int, text string, req int, random int) {
	reg := regexp.MustCompile("(?i)^acfur")
	call_bot := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	if call_bot {
		api.Sendgroupmsg(bot, uid, "Hi我是Acfur！"+new_text)
		return
	}
	is_sign, _ := regexp.MatchString("^签到$", new_text)
	if is_sign {
		Group.App_group_sign(bot, gid, uid)
		return
	}
	is_lunpan, _ := regexp.MatchString("^轮盘$", new_text)
	if is_lunpan {

	}

	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	if active {
		groupHandle_acfur(&bot, &gid, &uid, new_text)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		groupHandle_acfur_middle(&bot, &gid, &uid, &text)
	}
}

const group_function_number = 1

var group_function_type = []string{"unknow", "sign"}

func groupHandle_acfur(bot *int, gid *int, uid *int, text string) {
	switch text {
	case "help":
		api.Sendgroupmsg(*bot, *uid, app_default.Default_private_help)
		break

	case "签到":
		Group.App_group_sign(bot, gid, uid)
		break

	case "设定":
		Group.App_group_sign(bot, gid, uid)
		break

	default:
		groupHandle_acfur_middle(bot, gid, uid, &text)
		break
	}
}

func groupHandle_acfur_middle(bot *int, gid *int, uid *int, text *string) {
	function := make([]bool, group_function_number+1, group_function_number+1)
	new_text := make([]string, group_function_number+1, group_function_number+1)
	var wg sync.WaitGroup
	wg.Add(group_function_number)

	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(*text, []string{"密码", "password"})
		fmt.Println(str, ok)
		new_text[idx] = str
		function[idx] = ok
	}(1, &wg)
	wg.Wait()
	function_route := 0
	for i := range function {
		if function[i] == true {
			function_route = i
			break
		}
	}
	groupHandle_acfur_other(group_function_type[function_route], bot, gid, uid, new_text[function_route])
}

func groupHandle_acfur_other(Type string, bot *int, gid *int, uid *int, text string) {
	switch Type {

	default:
		api.Sendgroupmsg(*bot, *uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp")
		break
	}
}
