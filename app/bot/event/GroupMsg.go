package event

import (
	"errors"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"math"
	"regexp"
	"sync"
	"time"
)

type RefreshGroupStruct struct {
	UserId  int64
	SelfId  int64
	GroupId int64
}

var RefreshGroupChan = make(chan RefreshGroupStruct, 20)

type GM struct {
	Anonymous   interface{} `json:"anonymous"`
	Font        int64       `json:"font"`
	GroupID     int64       `json:"group_id"`
	Message     string      `json:"message"`
	MessageID   int64       `json:"message_id"`
	MessageSeq  int64       `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfID      int64       `json:"self_id"`
	Sender      _Sender     `json:"sender"`
	SubType     string      `json:"sub_type"`
	Time        int64       `json:"time"`
	UserID      int64       `json:"user_id"`
}

type _Sender struct {
	Age      int64  `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserID   int64  `json:"user_id"`
}

var GroupMsgChan = make(chan GM, 99)

func GroupMsg(gm GM) {
	GroupMsgChan <- gm
	is_self := false

	self_id := gm.SelfID
	user_id := gm.UserID
	group_id := gm.GroupID
	message_id := gm.MessageID
	message := gm.Message
	raw_message := gm.RawMessage

	if user_id == self_id {
		is_self = true
	}

	if !is_self {
		var group RefreshGroupStruct
		group.GroupId = group_id
		group.SelfId = self_id
		group.UserId = user_id
		RefreshGroupChan <- group
		botinfo := BotModel.Api_find(self_id)
		if len(botinfo) < 1 {
			Log.Crrs(errors.New("bot_not_found"), Calc.Any2String(self_id))
			return
		}
		GroupHandle(self_id, group_id, user_id, message_id, message, raw_message, gm.Sender)
	} else {

	}
}

func GroupHandle(self_id, group_id, user_id, message_id int64, message, raw_message string, sender _Sender) {
	text := message
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	groupmember := GroupMemberModel.Api_find(group_id, user_id)
	groupfunction := GroupFunctionModel.Api_find(group_id)
	if len(groupfunction) < 1 {
		GroupFunctionModel.Api_insert(group_id)
		groupfunction = GroupFunctionModel.Api_find(group_id)
	}
	botinfo := BotModel.Api_find(self_id)

	if active {
		if botinfo["end_time"].(int64) < time.Now().Unix() {
			api.Sendgroupmsg(self_id, group_id, app_default.Default_over_time, true)
			return
		}
		groupHandle_acfur(self_id, group_id, user_id, message_id, new_text, raw_message, sender, groupmember, groupfunction)
	} else {
		if botinfo["end_time"].(int64) < time.Now().Unix() {
			return
		}
		//在未激活acfur的情况下应该对原始内容进行还原
		groupHandle_acfur_middle(self_id, group_id, user_id, message_id, message, raw_message, sender, groupmember, groupfunction)
	}
}

func groupHandle_acfur(self_id, group_id, user_id int64, message_id int64, new_text, raw_message string, sender _Sender, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	admin := false
	owner := false
	if len(groupmember) > 0 {
		if groupmember["role"].(string) == "admin" {
			admin = true
		}
		if groupmember["role"].(string) == "owner" {
			admin = true
			owner = true
		}
	}
	auto_retract := true
	if groupfunction["auto_retract"].(int64) == 0 {
		auto_retract = false
	}
	switch new_text {
	case "help":
		api.Sendgroupmsg(self_id, group_id, app_default.Default_group_help, auto_retract)
		break

	case "app":
		api.Sendgroupmsg(self_id, group_id, app_default.Default_app_download_url, auto_retract)
		break

	case "设定":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_group_function_get_all(self_id, group_id, user_id, &new_text)
		break

	case "刷新":
		api.Sendgroupmsg(self_id, group_id, "可以使用“刷新人数”以及“刷新群信息”来控制刷新", auto_retract)
		break

	case "刷新人数":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(self_id, group_id, user_id)
				return
			}
		}
		Group.App_refreshmember(self_id, group_id)
		api.Sendgroupmsg(self_id, group_id, "群用户已经刷新", auto_retract)
		break

	case "刷新群信息":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_refresh_groupinfo(self_id, group_id)
		api.Sendgroupmsg(self_id, group_id, "群信息刷新完成", true)
		break

	case "测试撤回":
		var ret api.Struct_Retract
		ret.MessageId = message_id
		ret.Self_id = self_id
		if !admin {
			return
		}
		api.Retract_chan_instant <- ret
		break

	case "测试拼音":
		py, err := service.Serv_pinyin(new_text)
		if err != nil {

		} else {
			api.Sendgroupmsg(self_id, group_id, py, auto_retract)
		}
		break

	case "测试自动撤回":
		api.Sendgroupmsg(self_id, group_id, "自动撤回测试中……预计"+Calc.Int2String(app_conf.Retract_time_second+3)+"秒后撤回", auto_retract)
		break

	case "屏蔽":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		api.Sendgroupmsg(self_id, group_id, app_default.Default_str_ban_word, auto_retract)
		break

	case "屏蔽词":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_group_ban_word_list(self_id, group_id, user_id, new_text, 1, groupmember, groupfunction)
		break

	case "T出词":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_group_ban_word_list(self_id, group_id, user_id, new_text, 2, groupmember, groupfunction)
		break

	case "撤回词":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_group_ban_word_list(self_id, group_id, user_id, new_text, 3, groupmember, groupfunction)
		break

	case "清除小黑屋":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		GroupBanModel.Api_delete(group_id)
		break

	default:
		api.Sendgroupmsg(self_id, group_id, "Hi~我是Acfur，有任何问题可以发送acfurhelp哦~", true)
		break
	}
}

const group_function_number = 11

var group_function_type = []string{"unknow", "ban_group", "url_detect", "ban_weixin", "ban_share", "ban_word", "setting", "sign", "威望查询", "威望排行", "长度限制", "自动回复"}

func groupHandle_acfur_middle(self_id, group_id, user_id, message_id int64, message, raw_message string, sender _Sender, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	function := make([]bool, group_function_number+1, group_function_number+1)
	new_text := make([]string, group_function_number+1, group_function_number+1)
	var wg sync.WaitGroup
	wg.Add(group_function_number)

	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_group(raw_message)
		new_text[idx] = raw_message
		function[idx] = ok
	}(1, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_url_detect(raw_message)
		new_text[idx] = raw_message
		function[idx] = ok
	}(2, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_weixin(raw_message)
		new_text[idx] = raw_message
		function[idx] = ok
	}(3, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_share(raw_raw_message)
		new_text[idx] = raw_raw_message
		function[idx] = ok
	}(4, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(raw_message, []string{"acfur屏蔽"})
		new_text[idx] = str
		function[idx] = ok
	}(5, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(raw_message, []string{"acfur设定"})
		new_text[idx] = str
		function[idx] = ok
	}(6, &wg)
	//签到(直接)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(raw_message, []string{"签到"})
		new_text[idx] = str
		function[idx] = ok
	}(7, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(raw_message, []string{"积分查询", "威望查询"})
		new_text[idx] = str
		function[idx] = ok
	}(8, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		_, ok := service.Serv_text_match_all(raw_message, []string{"积分排行", "威望排行"})
		new_text[idx] = ""
		function[idx] = ok
	}(9, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		if int64(len(raw_message)) > groupfunction["word_limit"].(int64) {
			new_text[idx] = raw_message
			function[idx] = true
		}
	}(10, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_auto_reply(group_id, raw_message)
		new_text[idx] = str
		function[idx] = ok
	}(11, &wg)
	wg.Wait()
	function_route := 0
	for i := range function {
		if function[i] == true {
			function_route = i
			break
		}
	}
	groupHandle_acfur_other(group_function_type[function_route], self_id, group_id, user_id, message_id, new_text[function_route], raw_message, sender, groupmember, groupfunction)
}

func groupHandle_acfur_other(Type string, self_id, group_id, user_id, message_id int64, message, raw_message string, sender _Sender, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	admin := false
	owner := false
	if len(groupmember) > 0 {
		if groupmember["role"].(string) == "admin" {
			admin = true
		}
		if groupmember["role"].(string) == "owner" {
			admin = true
			owner = true
		}
	}
	auto_retract := true
	if groupfunction["auto_retract"].(int64) == 0 {
		auto_retract = false
	}
	var ret api.Struct_Retract
	ret.MessageId = group_id
	ret.Self_id = self_id

	switch Type {
	case "sign":
		Group.App_group_sign(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		break

	case "setting":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(self_id, group_id, user_id)
				return
			}
		}
		Group.App_group_function_set(self_id, group_id, user_id, message, message_id, groupmember, groupfunction)
		break

	case "ban_word":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(self_id, group_id, user_id)
				return
			}
		}
		Group.App_group_ban_word_set(self_id, group_id, user_id, message, message_id, groupmember, groupfunction)
		break

	case "url_detect":
		if groupfunction["ban_url"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			api.Sendgroupmsg(self_id, group_id, app_default.Default_ban_url, auto_retract)
			time := GroupBanModel.Api_count(group_id, user_id)
			GroupBanModel.Api_insert(group_id, user_id)
			api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "ban_group":
		if groupfunction["ban_group"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			api.Sendgroupmsg(self_id, group_id, app_default.Default_ban_group, auto_retract)
			time := GroupBanModel.Api_count(group_id, user_id)
			GroupBanModel.Api_insert(group_id, user_id)
			api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "ban_weixin":
		if groupfunction["ban_wx"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			api.Sendgroupmsg(self_id, group_id, app_default.Default_ban_weixin, auto_retract)
			time := GroupBanModel.Api_count(group_id, user_id)
			GroupBanModel.Api_insert(group_id, user_id)
			api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "ban_share":
		if groupfunction["ban_share"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			api.Sendgroupmsg(self_id, group_id, app_default.Default_ban_share, auto_retract)
			time := GroupBanModel.Api_count(group_id, user_id)
			GroupBanModel.Api_insert(group_id, user_id)
			api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "威望查询":
		Group.App_check_balance(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		break

	case "威望排行":
		Group.App_check_rank(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		break

	case "长度限制":
		api.Retract_chan_instant <- ret
		api.Sendgroupmsg(self_id, group_id, app_default.Default_length_limit+"本群消息长度限制为："+Calc.Int642String(groupfunction["word_limit"].(int64)), auto_retract)
		time := GroupBanModel.Api_count(group_id, user_id)
		GroupBanModel.Api_insert(group_id, user_id)
		api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		break

	case "自动回复":
		api.Sendgroupmsg(self_id, group_id, message, auto_retract)
		break

	default:
		if groupfunction["ban_repeat"].(int64) == 1 {
			num, err := Redis.GetInt(Calc.Md5(Calc.Any2String(user_id) + "_" + raw_message))
			if err != nil {
				Log.Crrs(err, tuuz.FUNCTION_ALL())

			}
			Redis.SetRaw(Calc.Md5(Calc.Any2String(user_id)+"_"+raw_message), num+1, int(groupfunction["repeat_time"].(int64)))
			if int64(num) > groupfunction["repeat_count"].(int64) {
				gb := GroupBanModel.Api_count(group_id, user_id)
				GroupBanModel.Api_insert(group_id, user_id)
				api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(gb)))
				api.Sendgroupmsg(self_id, group_id, "请不要在"+Calc.Any2String(groupfunction["repeat_time"])+"秒内重复发送相同内容", auto_retract)
				//gms := GroupMsgModel.Api_select(group_id, user_id, int(groupfunction["repeat_count"].(int64))+2)
				//for _, gm := range gms {
				//
				//}
			} else if int64(num)+1 > groupfunction["repeat_count"].(int64) {
				api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+Calc.Any2String(groupfunction["repeat_time"])+"秒内请勿重复发送相同内容", auto_retract)
			}
		}

		break
	}
}
