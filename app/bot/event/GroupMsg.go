package event

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Redis"
	"math"
	"regexp"
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

var GroupMsgChan = make(chan GM, 99)

func GroupMsg(gm GM) {
	GroupMsgChan <- gm
	is_self := false

	text := gm.Msg.Text
	bot := gm.LogonQQ
	uid := gm.FromQQ.UIN
	gid := gm.FromGroup.GIN
	retract := gm.Msg.Random

	if gm.LogonQQ == gm.FromQQ.UIN {
		is_self = true
	}

	uid_string := Calc.Int2String(uid)
	gid_string := Calc.Int2String(gid)

	text_exists := Redis.CheckExists("GroupMsg:" + gid_string + ":" + uid_string)
	if text_exists {
		return
	}

	Redis.SetRaw("GroupMsg:"+gid_string+":"+uid_string, Calc.Md5(text), 1)

	if !is_self {
		GroupHandle(bot, gid, uid, text, gm.Msg.Req, retract)
	} else {

	}
}

func GroupHandle(bot, gid, uid int, text string, req int, random int) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	groupmember := GroupMemberModel.Api_find(gid, uid)
	groupfunction := GroupFunctionModel.Api_find(gid)
	if active {
		groupHandle_acfur(&bot, &gid, &uid, text, new_text, &req, &random, groupmember, groupfunction)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		groupHandle_acfur_middle(&bot, &gid, &uid, &text, &req, &random, groupmember, groupfunction)
	}
}

func groupHandle_acfur(bot *int, gid *int, uid *int, text, new_text string, req *int, random *int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	admin := false
	owner := false
	if len(groupmember) > 0 {
		if groupmember["type"].(string) == "admin" {
			admin = true
		}
		if groupmember["type"].(string) == "owner" {
			admin = true
			owner = true
		}
	}

	switch new_text {
	case "help":
		api.Sendgroupmsg(*bot, *gid, app_default.Default_group_help, false)
		break

	case "app":
		api.Sendgroupmsg(*bot, *gid, app_default.Default_app_download_url, false)
		break

	case "设定":
		if !admin && !owner {
			service.Not_admin(bot, gid, uid)
			return
		}
		Group.App_group_function_get_all(bot, gid, uid, &new_text)
		break

	case "刷新":
		api.Sendgroupmsg(*bot, *gid, "可以使用“刷新人数”以及“刷新群信息”来控制刷新", true)
		break

	case "刷新人数":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(bot, gid, uid)
				return
			}
		}
		Group.App_refreshmember(bot, gid, uid)
		break

	case "刷新群信息":
		if !admin && !owner {
			service.Not_admin(bot, gid, uid)
			return
		}
		Group.App_refresh_groupinfo(bot, gid)
		break

	case "测试撤回":
		var ret Retract_group
		ret.Group = *gid
		ret.Fromqq = *bot
		ret.Random = *random
		ret.Req = *req
		if !admin {
			return
		}
		Retract_chan_group_instant <- ret
		break

	case "测试拼音":
		py, err := service.Serv_pinyin(new_text)
		if err != nil {

		} else {
			api.Sendgroupmsg(*bot, *gid, py, false)
		}
		break

	case "测试自动撤回":
		api.Sendgroupmsg(*bot, *gid, "自动撤回测试中……预计"+Calc.Int2String(app_conf.Retract_time_second+3)+"秒后撤回", true)
		break

	case "屏蔽":
		api.Sendgroupmsg(*bot, *gid, app_default.Default_str_ban_word, true)
		break

	case "屏蔽词":
		Group.App_group_ban_word_list(*bot, *gid, *uid, new_text, 1, groupmember, groupfunction)
		break

	case "T出词":
		Group.App_group_ban_word_list(*bot, *gid, *uid, new_text, 2, groupmember, groupfunction)
		break

	case "撤回词":
		Group.App_group_ban_word_list(*bot, *gid, *uid, new_text, 3, groupmember, groupfunction)
		break

	default:
		groupHandle_acfur_middle(bot, gid, uid, &text, req, random, groupmember, groupfunction)
		break
	}
}

const group_function_number = 6

var group_function_type = []string{"unknow", "ban_group", "url_detect", "ban_weixin", "ban_word", "setting", "sign"}

func groupHandle_acfur_middle(bot *int, gid *int, uid *int, text *string, req *int, random *int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	function := make([]bool, group_function_number+1, group_function_number+1)
	new_text := make([]string, group_function_number+1, group_function_number+1)
	var wg sync.WaitGroup
	wg.Add(group_function_number)

	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_group(*text)
		new_text[idx] = *text
		function[idx] = ok
	}(1, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_url_detect(*text)
		new_text[idx] = *text
		function[idx] = ok
	}(2, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_weixin(*text)
		new_text[idx] = *text
		function[idx] = ok
	}(3, &wg)

	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(*text, []string{"acfur屏蔽"})
		new_text[idx] = str
		function[idx] = ok
	}(4, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(*text, []string{"acfur设定"})
		new_text[idx] = str
		function[idx] = ok
	}(5, &wg)
	//签到(直接)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(*text, []string{"签到"})
		new_text[idx] = str
		function[idx] = ok
	}(6, &wg)
	wg.Wait()
	function_route := 0
	for i := range function {
		if function[i] == true {
			function_route = i
			break
		}
	}
	groupHandle_acfur_other(group_function_type[function_route], bot, gid, uid, new_text[function_route], req, random, groupmember, groupfunction)
}

func groupHandle_acfur_other(Type string, bot *int, gid *int, uid *int, text string, req *int, random *int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	admin := false
	owner := false
	if len(groupmember) > 0 {
		if groupmember["type"].(string) == "admin" {
			admin = true
		}
		if groupmember["type"].(string) == "owner" {
			admin = true
			owner = true
		}
	}
	var ret Retract_group
	ret.Group = *gid
	ret.Fromqq = *bot
	ret.Random = *random
	ret.Req = *req
	switch Type {
	case "sign":
		Group.App_group_sign(*bot, *gid, *uid, *req, *random, groupmember, groupfunction)
		break

	case "setting":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(bot, gid, uid)
				return
			}
		}
		Group.App_group_function_set(*bot, *gid, *uid, text, *req, *random, groupmember, groupfunction)
		break

	case "ban_word":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(bot, gid, uid)
				return
			}
		}
		Group.App_group_ban_word_set(*bot, *gid, *uid, text, groupmember, groupfunction)
		break

	case "url_detect":
		if groupfunction["ban_url"].(int64) == 1 {
			Retract_chan_group_instant <- ret
			api.Sendgroupmsg(*bot, *gid, app_default.Default_ban_url, true)
			time := GroupBanModel.Api_count(*gid, *uid)
			GroupBanModel.Api_insert(*gid, *uid)
			api.Mutegroupmember(*bot, *gid, *uid, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "ban_group":
		if groupfunction["ban_group"].(int64) == 1 {
			Retract_chan_group_instant <- ret
			api.Sendgroupmsg(*bot, *gid, app_default.Default_ban_group, true)
			time := GroupBanModel.Api_count(*gid, *uid)
			GroupBanModel.Api_insert(*gid, *uid)
			api.Mutegroupmember(*bot, *gid, *uid, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "ban_weixin":
		if groupfunction["ban_wx"].(int64) == 1 {
			Retract_chan_group_instant <- ret
			api.Sendgroupmsg(*bot, *gid, app_default.Default_ban_weixin, true)
			time := GroupBanModel.Api_count(*gid, *uid)
			GroupBanModel.Api_insert(*gid, *uid)
			api.Mutegroupmember(*bot, *gid, *uid, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
		}
		break

	case "积分查询":
		gbl := GroupBalanceModel.Api_find(*gid, *uid)
		if len(gbl) > 0 {
			gt := GroupBalanceModel.Api_select_gt(*gid, gbl["balance"])
			lt := GroupBalanceModel.Api_select_lt(*gid, gbl["balance"])
			if len(gt) > 9 {

			}
		} else {

		}

		break

	case "积分排行":
		gbl := GroupBalanceModel.Api_select(*gid, 10)
		str := ""
		for i1, i2 := range gbl {
			user := GroupMemberModel.Api_find(*gid, i2["uid"].(int64))
			if len(user) > 0 {
				if len(user["card"]) > 2 {
					str += "第" + Calc.Int2String(i1+1) + "名：" + user["card"].(string) + "\r\n"
				} else {
					str += "第" + Calc.Int2String(i1+1) + "名：" + user["nickname"].(string) + "\r\n"
				}

			}
		}
		break

	default:
		api.Sendgroupmsg(*bot, *uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		break
	}
}
