package event

import (
	"main.go/app/bot/action/Private"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Redis"
	"regexp"
	"strings"
	"sync"
)

type PM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN      int    `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"FromQQ"`
	LogonQQ   int `json:"LogonQQ"`
	TimeStamp struct {
		Recv int `json:"Recv"`
		Send int `json:"Send"`
	} `json:"TimeStamp"`
	FromGroup struct {
		GIN int `json:"GIN"`
	} `json:"FromGroup"`
	Msg struct {
		Req         int    `json:"Req"`
		Seq         int64  `json:"Seq"`
		Type        int    `json:"Type"`
		SubType     int    `json:"SubType"`
		SubTempType int    `json:"SubTempType"`
		Text        string `json:"Text"`
		BubbleID    int    `json:"BubbleID"`
	} `json:"Msg"`
	Hb struct {
		Type int `json:"Type"`
	} `json:"Hb"`
	File struct {
		ID   string `json:"ID"`
		MD5  string `json:"MD5"`
		Name string `json:"Name"`
		Size int    `json:"Size"`
	} `json:"File"`
}

var PrivateMsgChan = make(chan PM, 99)

func PrivateMsg(pm PM) {
	PrivateMsgChan <- pm

	bot := pm.LogonQQ
	uid := pm.FromQQ.UIN
	uid_string := Calc.Int2String(uid)
	text := pm.Msg.Text

	text_exists := Redis.CheckExists("PrivateMsg:" + uid_string)
	if text_exists {
		return
	}

	Redis.SetRaw("PrivateMsg:"+uid_string, Calc.Md5(text), 1)

	PrivateHandle(bot, uid, text)
}

func PrivateHandle(bot int, uid int, text string) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	if active {
		privateHandle_acfur(&bot, &uid, new_text, text)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		if private_default_reply(&bot, &uid, &text) {
			return
		}
		auto_reply := PrivateAutoReplyModel.Api_find_byKey(text)
		if len(auto_reply) > 0 {
			if auto_reply["value"] == nil {
				return
			}
			api.Sendprivatemsg(bot, uid, auto_reply["value"].(string), false)
		} else {
			private_auto_reply(&bot, &uid, &text)
		}
	}
}

func private_default_reply(bot *int, uid *int, text *string) bool {
	default_reply := BotDefaultReplyModel.Api_select()
	for _, auto_reply := range default_reply {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(*text, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			api.Sendprivatemsg(*bot, *uid, auto_reply["value"].(string), false)
			return true
		}
	}
	return false
}

func private_auto_reply(bot *int, uid *int, text *string) {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(*bot)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(*text, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			api.Sendprivatemsg(*bot, *uid, auto_reply["value"].(string), false)
			break
		}
	}
}

func privateHandle_acfur(bot *int, uid *int, text, origin_text string) {
	switch text {
	case "help":
		botinfo := BotModel.Api_find(*bot)
		if len(botinfo) > 0 {
			if botinfo["owner"].(int64) == int64(*uid) {
				api.Sendprivatemsg(*bot, *uid, app_default.Default_private_help+app_default.Default_private_help_for_RobotOwner, false)
			} else {
				api.Sendprivatemsg(*bot, *uid, app_default.Default_private_help, false)
			}
		} else {
			api.Sendprivatemsg(*bot, *uid, app_default.Default_private_help, false)
		}
		break

	case "登录", "登陆", "login":
		Private.App_userLogin(*bot, *uid, text)
		break

	case "清除登录":
		Private.App_userClearLogin(*bot, *uid)
		break

	case "解绑":
		Private.App_unbind_bot(*bot, *uid, text)
		break

	case "绑定":
		api.Sendprivatemsg(*bot, *uid, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
		break

	case "绑定群":
		groupbinds := BotGroupAllowModel.Api_select(*bot)
		groups := []string{}
		for _, groupbind := range groupbinds {
			groups = append(groups, Calc.Any2String(groupbind["gid"]))
		}
		api.Sendprivatemsg(*bot, *uid, "您的机器人可在如下群中使用:\r\n"+strings.Join(groups, ",")+
			"\r\n您可以使用：acfur绑定群:群号，来绑定新群，\r\n使用：acfur解绑群:群号，解绑", false)
		break

	default:
		privateHandle_acfur_middle(bot, uid, text, origin_text)
		break
	}
}

const private_function_number = 5

var private_function_type = []string{"unknow", "密码", "修改密码", "绑定群", "解绑群", "绑定"}

func privateHandle_acfur_middle(bot *int, uid *int, text, origin_text string) {
	function := make([]bool, private_function_number+1, private_function_number+1)
	new_text := make([]string, private_function_number+1, private_function_number+1)
	var wg sync.WaitGroup
	wg.Add(private_function_number)

	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(text, []string{"密码", "password"})
		new_text[idx] = str
		function[idx] = ok
	}(1, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(text, []string{"修改密码", "change_secret"})
		new_text[idx] = str
		function[idx] = ok
	}(2, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(text, []string{"绑定群", "bindgroup"})
		new_text[idx] = str
		function[idx] = ok
	}(3, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(text, []string{"解绑群", "unbindgroup"})
		new_text[idx] = str
		function[idx] = ok
	}(4, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(text, []string{"绑定", "bind"})
		new_text[idx] = str
		function[idx] = ok
	}(5, &wg)
	wg.Wait()
	function_route := 0
	for i := range function {
		if function[i] == true {
			function_route = i
			break
		}
	}
	privateHandle_acfur_other(private_function_type[function_route], bot, uid, new_text[function_route])
}

func privateHandle_acfur_other(Type string, bot *int, uid *int, text string) {
	botinfo := BotModel.Api_find(*bot)
	switch Type {
	case "密码":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_userChangePassword(*bot, *uid, text)
		} else {
			api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "绑定":
		Private.App_bind_robot(*bot, *uid, text)
		break

	case "修改密码":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_change_bot_secret(*bot, *uid, text)
		} else {
			api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "绑定群":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_bind_group(*bot, *uid, text)
		} else {
			api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "解绑群":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_unbind_group(*bot, *uid, text)
		} else {
			api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	default:
		api.Sendprivatemsg(*bot, *uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		break
	}
}
