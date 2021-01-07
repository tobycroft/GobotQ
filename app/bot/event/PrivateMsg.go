package event

import (
	"fmt"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/app/bot/model/PrivateMsgModel"
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

func PrivateMsg(pm PM) {
	PrivateMsgModel.Api_insert(pm.LogonQQ, pm.FromQQ.UIN, pm.Msg.Text, pm.Msg.Req, pm.Msg.Seq, pm.Msg.Type, pm.Msg.SubType, pm.File.ID,
		pm.File.MD5, pm.File.Name, pm.File.Size)

	bot := pm.LogonQQ
	uid := pm.FromQQ.UIN
	uid_string := Calc.Int2String(uid)
	text := pm.Msg.Text

	text_exists := Redis.CheckExists("PrivateMsg_" + uid_string)
	if text_exists {
		return
	}

	Redis.SetRaw("PrivateMsg_"+uid_string, Calc.Md5(text), 2)

	PrivateHandle(bot, uid, text)
}

func PrivateHandle(bot int, uid int, text string) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	if active {
		privateHandle_acfur(bot, uid, new_text)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		if private_default_reply(bot, uid, text) {
			return
		}
		auto_reply := PrivateAutoReplyModel.Api_find_byKey(text)
		if len(auto_reply) > 0 {
			if auto_reply["value"] == nil {
				return
			}
			api.Sendprivatemsg(bot, uid, auto_reply["value"].(string))
		} else {
			private_auto_reply(bot, uid, text)
		}
	}
}

func private_default_reply(bot int, uid int, text string) bool {
	default_reply := BotDefaultReplyModel.Api_select()
	for _, auto_reply := range default_reply {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(text, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			api.Sendprivatemsg(bot, uid, auto_reply["value"].(string))
			return true
		}
	}
	return false
}

func private_auto_reply(bot int, uid int, text string) {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(bot)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(text, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			api.Sendprivatemsg(bot, uid, auto_reply["value"].(string))
			break
		}
	}
}

const private_function_number = 1

var private_function_type = []string{"unknow", "password"}

func privateHandle_acfur(bot int, uid int, text string) {
	switch text {
	case "help":
		api.Sendprivatemsg(bot, uid, app_default.Default_private_help)
		break

	case "登录", "登陆", "login":
		Private.App_userLogin(bot, uid, text)
		break

	case "清除登录":
		Private.App_userClearLogin(bot, uid)
		break

	default:

		function := make([]bool, private_function_number+1, private_function_number+1)
		new_text := make([]string, private_function_number+1, private_function_number+1)
		fmt.Println(function)
		var wg sync.WaitGroup
		wg.Add(private_function_number)

		go func(idx int, wg *sync.WaitGroup) {
			defer wg.Done()
			str, ok := service.Serv_text_match(text, []string{"密码", "password"})
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

		privateHandle_acfur_other(private_function_type[function_route], bot, uid, new_text[function_route])

		break
	}
}

func privateHandle_acfur_other(Type string, bot int, uid int, text string) {
	switch Type {
	case "password":
		Private.App_userChangePassword(bot, uid, text)
		break

	default:
		api.Sendprivatemsg(bot, uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp")
		break
	}
}
