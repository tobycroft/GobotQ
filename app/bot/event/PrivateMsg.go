package event

import (
	"errors"
	"fmt"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotDefaultReplyModel"
	"main.go/app/bot/model/BotGroupAllowModel"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"net"

	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"strings"
	"sync"
	"time"
)

type PrivateMessageStruct struct {
	remoteaddr  net.Addr
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int64  `json:"message_id"`
	TargetId    int64  `json:"target_id"`
	PeerId      int64  `json:"peer_id"`
	UserId      int64  `json:"user_id"`
	Message     []struct {
		Data struct {
			Text string `json:"text"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"message"`
	RawMessage string        `json:"raw_message"`
	Font       int64         `json:"font"`
	Sender     PrivateSender `json:"sender"`
}
type PrivateSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Role     string `json:"role"`
	Title    string `json:"title"`
	Level    string `json:"level"`
}

var PrivateMsgChan = make(chan PrivateMessageStruct, 99)

func (pm PrivateMessageStruct) PrivateMsg() {
	PrivateMsgChan <- pm
	selfId := pm.SelfId
	user_id := pm.UserId
	group_id := int64(0)
	user_idString := Calc.Int642String(user_id)
	message := pm.RawMessage
	rawMessage := pm.RawMessage

	if Redis.CheckExists("PrivateMsg:" + user_idString) {
		return
	}

	Redis.String_set("PrivateMsg:"+user_idString, Calc.Md5(message), 1*time.Second)

	PrivateHandle(selfId, user_id, group_id, message, rawMessage, pm.remoteaddr.String())
}

func PrivateHandle(selfId, user_id, group_id int64, message, rawMessage, remoteip string) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(message)
	new_text := reg.ReplaceAllString(message, "")

	botinfo := BotModel.Api_find(selfId)
	if botinfo["url"] == nil {
		return
	}

	if strings.Contains(remoteip, botinfo["url"].(string)) {
		return
	}

	if len(botinfo) < 1 {
		Log.Crrs(errors.New("bot_not_found"), Calc.Any2String(selfId))
		return
	}
	if botinfo["end_time"].(int64) < time.Now().Unix() {
		iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, app_default.Default_over_time, false)
		return
	}
	if active {
		fmt.Println("privateHandle_acfur")
		privateHandle_acfur(selfId, user_id, group_id, new_text, message)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		if private_default_reply(selfId, user_id, group_id, message) {
			return
		}
		auto_reply := PrivateAutoReplyModel.Api_find_byKey(message)
		if len(auto_reply) > 0 {
			if auto_reply["value"] == nil {
				return
			}
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, auto_reply["value"].(string), false)
		} else {
			private_auto_reply(selfId, user_id, group_id, message)
		}
	}
}

func private_default_reply(selfId, user_id, group_id int64, message string) bool {
	default_reply := BotDefaultReplyModel.Api_select()
	for _, auto_reply := range default_reply {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, auto_reply["value"].(string), false)
			return true
		}
	}
	return false
}

func private_auto_reply(selfId, user_id, group_id int64, message string) {
	auto_replys := PrivateAutoReplyModel.Api_select_semi(selfId)
	for _, auto_reply := range auto_replys {
		if auto_reply["key"] == nil {
			continue
		}
		if strings.Contains(message, auto_reply["key"].(string)) {
			if auto_reply["value"] == nil {
				continue
			}
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, auto_reply["value"].(string), true)
			break
		}
	}
}

func privateHandle_acfur(selfId, user_id, group_id int64, message, origin_text string) {
	switch message {

	case "app", "下载":
		iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, app_default.Default_app_download_url, true)
		break

	case "help":
		botinfo := BotModel.Api_find(selfId)
		if len(botinfo) > 0 {
			if botinfo["owner"].(int64) == int64(user_id) {
				iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, app_default.Default_private_help+app_default.Default_private_help_for_RobotOwner, false)
			} else {
				iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, app_default.Default_private_help, false)
			}
		} else {
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, app_default.Default_private_help, false)
		}
		break

	case "登录", "登陆", "login":
		Private.App_userLogin(selfId, user_id, group_id, message)
		break

	case "清除登录":
		Private.App_userClearLogin(selfId, user_id, group_id)
		break

	case "解绑":
		Private.App_unbind_bot(selfId, user_id, group_id, message)
		break

	case "绑定":
		iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "请使用\"acfur绑定(+)本机器人密码\"来绑定您的机器人", false)
		break

	case "绑定群":
		groupbinds := BotGroupAllowModel.Api_select(selfId)
		groups := []string{}
		for _, groupbind := range groupbinds {
			groups = append(groups, Calc.Any2String(groupbind["group_id"]))
		}
		iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "您的机器人可在如下群中使用:\r\n"+strings.Join(groups, ",")+
			"\r\n您可以使用：acfur绑定群:群号，来绑定新群，\r\n使用：acfur解绑群:群号，解绑", false)
		break

	default:
		privateHandle_acfur_middle(selfId, user_id, group_id, message, origin_text)
		break
	}
}

const private_function_number = 5

var private_function_type = []string{"unknow", "密码", "修改密码", "绑定群", "解绑群", "绑定"}

func privateHandle_acfur_middle(selfId, user_id, group_id int64, message, origin_text string) {
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
	privateHandle_acfur_other(private_function_type[function_route], selfId, user_id, group_id, new_text[function_route])
}

func privateHandle_acfur_other(Type string, selfId int64, user_id, group_id int64, message string) {
	botinfo := BotModel.Api_find(selfId)
	switch Type {
	case "密码":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_userChangePassword(selfId, user_id, group_id, message)
		} else {
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "绑定":
		Private.App_bind_robot(selfId, user_id, group_id, message)
		break

	case "修改密码":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_change_bot_secret(selfId, user_id, group_id, message)
		} else {
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "绑定群":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_bind_group(selfId, user_id, group_id, message)
		} else {
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "解绑群":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_unbind_group(selfId, user_id, group_id, message)
		} else {
			iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	default:
		iapi.Api{}.Sendprivatemsg(selfId, user_id, group_id, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		break
	}
}
