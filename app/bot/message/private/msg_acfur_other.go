package private

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/Private"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"regexp"
	"sync"
)

const private_function_number = 5

var private_function_type = []string{"unknow", "密码", "修改密码", "绑定群", "解绑群", "绑定"}

func private_message_setting_change_with_acfur() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivateAcfur) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			pm := es.Json
			self_id := pm.SelfId
			user_id := pm.UserId
			group_id := int64(0)
			message := pm.RawMessage

			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(message)
			new_text := reg.ReplaceAllString(message, "")
			if active {

			}
		}
	}
}

func acfur_password() {

	service.Serv_text_match(message, []string{"密码", "password"})
}

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
			iapi.Api.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "绑定":
		Private.App_bind_robot(selfId, user_id, group_id, message)
		break

	case "修改密码":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_change_bot_secret(selfId, user_id, group_id, message)
		} else {
			iapi.Api.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "绑定群":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_bind_group(selfId, user_id, group_id, message)
		} else {
			iapi.Api.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	case "解绑群":
		if int64(user_id) == botinfo["owner"].(int64) {
			Private.App_unbind_group(selfId, user_id, group_id, message)
		} else {
			iapi.Api.Sendprivatemsg(selfId, user_id, group_id, "您未拥有这个机器人的权限，请先绑定机器人", true)
		}
		break

	default:
		iapi.Api.Sendprivatemsg(selfId, user_id, group_id, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		break
	}
}
