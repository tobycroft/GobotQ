package event

import (
	"errors"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"sync"
	"time"
)

type RefreshGroupStruct struct {
	UserId  int64
	SelfId  int64
	GroupId int64
}

var RefreshGroupChan = make(chan RefreshGroupStruct, 100)

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

func GroupMsg(gm GM, remoteip string) {
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
		botinfo := BotModel.Api_find(self_id)
		if len(botinfo) < 1 {
			Log.Crrs(errors.New("bot_not_found"), Calc.Any2String(self_id))
			return
		}
		has1 := Redis.CheckExists("__groupinfo__" + Calc.Int642String(group_id) + "_" + Calc.Int642String(user_id))
		has2 := Redis.CheckExists("__userinfo__" + Calc.Int642String(group_id) + "_" + Calc.Int642String(user_id))
		if !has1 || !has2 {
			var group RefreshGroupStruct
			group.GroupId = group_id
			group.SelfId = self_id
			group.UserId = user_id
			RefreshGroupChan <- group
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

	if active || service.Serv_is_at_me(self_id, message) {
		if botinfo["end_time"].(int64) < time.Now().Unix() {
			Group.AutoMessage(self_id, group_id, user_id, app_default.Default_over_time, groupfunction)
			return
		}
		groupHandle_acfur(self_id, group_id, user_id, message_id, new_text, message, raw_message, sender, groupmember, groupfunction)
	} else {
		if botinfo["end_time"].(int64) < time.Now().Unix() {
			return
		}
		//在未激活acfur的情况下应该对原始内容进行还原
		groupHandle_acfur_middle(self_id, group_id, user_id, message_id, message, raw_message, sender, groupmember, groupfunction)
	}
}

func groupHandle_acfur(self_id, group_id, user_id int64, message_id int64, new_text, message, raw_message string, sender _Sender, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
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
	switch new_text {

	case "":
		api.Sendgroupmsg(self_id, group_id, app_default.Default_welcome, true)
		break

	case "道具", "商店", "商城":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_daoju, groupfunction)
		break

	case "轮盘":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_lunpan_help, groupfunction)
		break

	case "help":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_group_help, groupfunction)
		break

	case "app":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_app_download_url, groupfunction)
		break

	case "设定":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_group_function_get_all(self_id, group_id, user_id, new_text, groupfunction)
		break

	case "刷新":
		Group.AutoMessage(self_id, group_id, user_id, "可以使用“刷新人数”以及“刷新群信息”来控制刷新", groupfunction)
		break

	case "权限":
		Group.AutoMessage(self_id, group_id, user_id, "我当前的权限为："+Group.BotPowerRefresh(group_id, self_id), groupfunction)
		break

	case "随机数测试":
		rand1 := Calc.Rand(1, 100)
		rand2 := Calc.Rand(1, 100)
		Group.AutoMessage(self_id, group_id, user_id, "随机数1："+Calc.Any2String(rand1)+"\n随机数2："+Calc.Any2String(rand2), groupfunction)
		break

	case "刷新人数":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(self_id, group_id, user_id)
				return
			}
		}
		Group.App_refreshmember(self_id, group_id)
		Group.AutoMessage(self_id, group_id, user_id, "群用户已经刷新", groupfunction)
		break

	case "刷新群信息":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.App_refresh_groupinfo(self_id, group_id)
		Group.AutoMessage(self_id, group_id, user_id, "群信息刷新完成", groupfunction)
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

	case "测试T出测试":
		Group.App_kick_user(self_id, group_id, user_id, true, groupfunction, "测试")
		break

	case "测试禁言测试":
		Group.App_ban_user(self_id, group_id, user_id, true, groupfunction, "测试")
		break

	case "测试拼音":
		py, err := service.Serv_pinyin(new_text)
		if err != nil {

		} else {
			Group.AutoMessage(self_id, group_id, user_id, py, groupfunction)
		}
		break

	case "测试自动撤回":
		Group.AutoMessage(self_id, group_id, user_id, "自动撤回测试中……预计"+Calc.Int2String(app_conf.Retract_time_second+3)+"秒后撤回", groupfunction)
		break

	case "屏蔽":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			return
		}
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_str_ban_word, groupfunction)
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

		if GroupBanPermenentModel.Api_delete_byGroupId(group_id) {
			Group.AutoMessage(self_id, group_id, user_id, "小黑屋已经清除", groupfunction)
		} else {
			Group.AutoMessage(self_id, group_id, user_id, "小黑屋里面没有人啦~", groupfunction)
		}
		break

	default:
		groupHandle_acfur_middle(self_id, group_id, user_id, message_id, message, raw_message, sender, groupmember, groupfunction)
		break
	}
}

const group_function_number = 15

var group_function_type = []string{"unknow", "ban_group", "url_detect", "ban_weixin", "ban_share", "ban_word", "setting",
	"sign", "轮盘", "威望查询", "威望排行", "长度限制", "自动回复", "atme", "道具"}

func groupHandle_acfur_middle(self_id, group_id, user_id, message_id int64, message, raw_message string, sender _Sender, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	function := make([]bool, group_function_number+1, group_function_number+1)
	new_text := make([]string, group_function_number+1, group_function_number+1)
	var wg sync.WaitGroup
	wg.Add(group_function_number)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		new_text[idx] = message
	}(0, &wg)
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
		ok := service.Serv_ban_weixin(message)
		new_text[idx] = message
		function[idx] = ok
	}(3, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_share(message)
		new_text[idx] = message
		function[idx] = ok
	}(4, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur屏蔽"})
		new_text[idx] = str
		function[idx] = ok
	}(5, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur设定"})
		new_text[idx] = str
		function[idx] = ok
	}(6, &wg)
	//签到(直接)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(message, []string{"签到"})
		new_text[idx] = str
		function[idx] = ok
	}(7, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"轮盘"})
		new_text[idx] = str
		function[idx] = ok
	}(8, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(message, []string{"积分查询", "查询积分", "威望查询", "查询威望", "钱包", "查询余额", "余额查询"})
		new_text[idx] = str
		function[idx] = ok
	}(9, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		_, ok := service.Serv_text_match_all(message, []string{"积分排行", "威望排行", "排行榜"})
		new_text[idx] = ""
		function[idx] = ok
	}(10, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		if int64(len(raw_message)) > groupfunction["word_limit"].(int64) {
			new_text[idx] = raw_message
			function[idx] = true
		}
	}(11, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_auto_reply(group_id, raw_message)
		new_text[idx] = str
		function[idx] = ok
	}(12, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		_, ok := service.Serv_text_match_any(message, []string{"[CQ:at,qq=" + Calc.Any2String(self_id) + "]"})
		new_text[idx] = ""
		function[idx] = ok
	}(13, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"道具"})
		new_text[idx] = str
		function[idx] = ok
	}(14, &wg)
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
	ret.MessageId = message_id
	ret.Self_id = self_id

	switch Type {

	case "道具":
		Group.App_group_daoju(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		break

	case "atme":
		api.Sendgroupmsg(self_id, group_id, app_default.Default_welcome, true)
		break

	case "sign":
		Group.App_group_sign(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		break

	case "轮盘":
		Group.App_group_lunpan(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
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
			Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_url)
		}
		break

	case "ban_group":
		if groupfunction["ban_group"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			Group.App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_group)
		}
		break

	case "ban_weixin":
		if groupfunction["ban_wx"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_weixin)
		}
		break

	case "ban_share":
		if groupfunction["ban_share"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				api.Retract_chan_instant <- ret
			}
			Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_share)
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
		Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction,
			app_default.Default_length_limit+"本群消息长度限制为："+Calc.Int642String(groupfunction["word_limit"].(int64)))
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
				Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, "请不要在"+Calc.Any2String(groupfunction["repeat_time"])+"秒内重复发送相同内容")
			} else if int64(num)+1 > groupfunction["repeat_count"].(int64) {
				api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+Calc.Any2String(groupfunction["repeat_time"])+"秒内请勿重复发送相同内容", auto_retract)
			}
		}

		//验证程序
		code, err := Redis.GetString("verify_" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
		if err != nil {

		} else {
			if code == message {
				GroupBanPermenentModel.Api_delete(group_id, user_id)
				Redis.Del("ban_" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
				str := ""
				if groupfunction["auto_welcome"] == 1 {
					str = "\n" + Calc.Any2String(groupfunction["welcome_word"])
				}
				api.Retract_chan_instant <- ret
				api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+"验证成功"+str, true)
			}
		}

		if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {
			api.Retract_chan_instant <- ret
			api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+"请先输入上述四位数字"+Calc.Any2String(code), true)
		} else if Redis.CheckExists("ban_" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id)) {
			api.Retract_chan_instant <- ret
			api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+"请先输入上述四位数字"+Calc.Any2String(code), true)
		}
		break
	}
}
