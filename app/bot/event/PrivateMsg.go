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
	"time"
)

type PM struct {
	Font        int    `json:"font"`
	Message     string `json:"message"`
	MessageID   int64  `json:"message_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	RawMessage  string `json:"raw_message"`
	SelfID      int64  `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		GroupID  int64  `json:"group_id"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserID   int64  `json:"user_id"`
	} `json:"sender"`
	SubType    string `json:"sub_type"`
	TempSource int    `json:"temp_source"`
	Time       int64  `json:"time"`
	UserID     int64  `json:"user_id"`
}

var PrivateMsgChan = make(chan PM, 99)

func PrivateMsg(pm PM) {
	/*
		message：消息事件
		notice：通知事件
		request：请求事件
		meta_event：元事件
	*/
	PrivateMsgChan <- pm
	bot := pm.SelfID
	uid := pm.UserID
	gid := pm.Sender.GroupID
	uid_string := Calc.Int642String(uid)
	message := pm.Message
	raw_message := pm.RawMessage

	text_exists := Redis.CheckExists("PrivateMsg:" + uid_string)
	if text_exists {
		return
	}

	//Redis.SetRaw("PrivateMsg:"+uid_string, Calc.Md5(message), 1)

	PrivateHandle(bot, uid, gid, message, raw_message)
}

func PrivateHandle(bot int64, uid, gid int64, message, raw_message string) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(message)
	new_text := reg.ReplaceAllString(message, "")

	botinfo := BotModel.Api_find(bot)
	if botinfo["end_time"].(int64) < time.Now().Unix() {
		if gid != 0 {
			api.Sendgrouptempmsg(bot, gid, uid, app_default.Default_over_time, false)
		} else {
			api.Sendprivatemsg(bot, uid, app_default.Default_over_time, false)
		}
		return
	}

	if active {
		privateHandle_acfur(&bot, &uid, &gid, new_text, message)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		if private_default_reply(&bot, &uid, &gid, &message) {
			return
		}
		auto_reply := PrivateAutoReplyModel.Api_find_byKey(message)
		if len(auto_reply) > 0 {
			if auto_reply["value"] == nil {
				return
			}
			if gid != 0 {
				api.Sendgrouptempmsg(bot, gid, uid, auto_reply["value"].(string), false)
			} else {
				api.Sendprivatemsg(bot, uid, auto_reply["value"].(string), false)
			}

		} else {
			private_auto_reply(&bot, &uid, &gid, &message)
		}
	}
}

func private_default_reply(bot *int64, uid, gid *int64, message *string) bool {
	default_reply := BotDefaultReplyModel.Api_select()
	for _, auto_reply := range default_reply {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(*message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, auto_reply["value"].(string), false)
			} else {
				api.Sendprivatemsg(*bot, *uid, auto_reply["value"].(string), false)
			}
			return true
		}
	}
	return false
}

func private_auto_reply(bot *int64, uid, gid *int64, message *string) {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(*bot)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(*message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, auto_reply["value"].(string), false)
			} else {
				api.Sendprivatemsg(*bot, *uid, auto_reply["value"].(string), false)
			}
			break
		}
	}
}

func privateHandle_acfur(bot *int64, uid, gid *int64, message, origin_text string) {
	switch message {
	case "help":
		botinfo := BotModel.Api_find(*bot)
		if len(botinfo) > 0 {
			if botinfo["owner"].(int64) == int64(*uid) {
				if *gid != 0 {
					api.Sendgrouptempmsg(*bot, *gid, *uid, app_default.Default_private_help+app_default.Default_private_help_for_RobotOwner, false)
				} else {
					api.Sendprivatemsg(*bot, *uid, app_default.Default_private_help+app_default.Default_private_help_for_RobotOwner, false)
				}
			} else {
				if *gid != 0 {
					api.Sendgrouptempmsg(*bot, *gid, *uid, app_default.Default_private_help, false)
				} else {
					api.Sendprivatemsg(*bot, *uid, app_default.Default_private_help, false)
				}
			}
		} else {
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, app_default.Default_private_help, false)
			} else {
				api.Sendprivatemsg(*bot, *uid, app_default.Default_private_help, false)
			}
		}
		break

	case "登录", "登陆", "login":
		Private.App_userLogin(*bot, *uid, *gid, message)
		break

	case "清除登录":
		Private.App_userClearLogin(*bot, *uid, *gid)
		break

	case "解绑":
		Private.App_unbind_bot(*bot, *uid, *gid, message)
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
		privateHandle_acfur_middle(bot, uid, gid, message, origin_text)
		break
	}
}

const private_function_number = 5

var private_function_type = []string{"unknow", "密码", "修改密码", "绑定群", "解绑群", "绑定"}

func privateHandle_acfur_middle(bot *int64, uid, gid *int64, message, origin_text string) {
	function := make([]bool, private_function_number+1, private_function_number+1)
	new_text := make([]string, private_function_number+1, private_function_number+1)
	var wg sync.WaitGroup
	wg.Add(private_function_number)

	go func(idx int64, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"密码", "password"})
		new_text[idx] = str
		function[idx] = ok
	}(1, &wg)
	go func(idx int64, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"修改密码", "change_secret"})
		new_text[idx] = str
		function[idx] = ok
	}(2, &wg)
	go func(idx int64, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"绑定群", "bindgroup"})
		new_text[idx] = str
		function[idx] = ok
	}(3, &wg)
	go func(idx int64, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"解绑群", "unbindgroup"})
		new_text[idx] = str
		function[idx] = ok
	}(4, &wg)
	go func(idx int64, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"绑定", "bind"})
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
	privateHandle_acfur_other(private_function_type[function_route], bot, uid, gid, new_text[function_route])
}

func privateHandle_acfur_other(Type string, bot *int64, uid, gid *int64, message string) {
	botinfo := BotModel.Api_find(*bot)
	switch Type {
	case "密码":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_userChangePassword(*bot, *uid, message)
		} else {
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			} else {
				api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			}
		}
		break

	case "绑定":
		Private.App_bind_robot(*bot, *uid, *gid, message)
		break

	case "修改密码":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_change_bot_secret(*bot, *uid, *gid, message)
		} else {
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			} else {
				api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			}
		}
		break

	case "绑定群":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_bind_group(*bot, *uid, message)
		} else {
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			} else {
				api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			}
		}
		break

	case "解绑群":
		if int64(*uid) == botinfo["owner"].(int64) {
			Private.App_unbind_group(*bot, *uid, message)
		} else {
			if *gid != 0 {
				api.Sendgrouptempmsg(*bot, *gid, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			} else {
				api.Sendprivatemsg(*bot, *uid, "您未拥有这个机器人的权限，请先绑定机器人", true)
			}
		}
		break

	default:
		if *gid != 0 {
			api.Sendgrouptempmsg(*bot, *gid, *uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		} else {
			api.Sendprivatemsg(*bot, *uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		}
		break
	}
}
