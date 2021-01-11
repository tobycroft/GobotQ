package event

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Redis"
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

	Redis.SetRaw("GroupMsg:"+gid_string+":"+uid_string, Calc.Md5(text), 3)

	if !is_self {
		GroupHandle(bot, gid, uid, text, gm.Msg.Req, retract)
	} else {

	}
}

func GroupHandle(bot, gid, uid int, text string, req int, random int) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	groupmember := GroupMemberModel.Api_find(bot, gid, uid)
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
		api.Sendgroupmsg(*bot, *gid, app_default.Default_private_help, false)
		break

	case "设定":
		if !admin && !owner {
			not_admin(bot, gid, uid)
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
				not_admin(bot, gid, uid)
				return
			}
		}
		Group.App_refreshmember(bot, gid, uid)
		break

	case "刷新群信息":
		if !admin && !owner {
			not_admin(bot, gid, uid)
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

	default:
		groupHandle_acfur_middle(bot, gid, uid, &text, req, random, groupmember, groupfunction)
		break
	}
}

func not_admin(bot *int, gid *int, uid *int) {
	api.Sendgroupmsg(*bot, *gid, "你不是本群的管理员，无法使用本功能"+service.Serv_at(*uid), true)
}

const group_function_number = 2

var group_function_type = []string{"unknow", "sign", "setting"}

func groupHandle_acfur_middle(bot *int, gid *int, uid *int, text *string, req *int, random *int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	function := make([]bool, group_function_number+1, group_function_number+1)
	new_text := make([]string, group_function_number+1, group_function_number+1)
	var wg sync.WaitGroup
	wg.Add(group_function_number)
	//签到(直接)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(*text, []string{"签到"})
		new_text[idx] = str
		function[idx] = ok
	}(1, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(*text, []string{"设定"})
		new_text[idx] = str
		function[idx] = ok
	}(2, &wg)
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
	switch Type {

	case "sign":
		Group.App_group_sign(*bot, *gid, *uid, *req, *random, groupmember, groupfunction)
		break

	case "setting":
		Group.App_group_function_set(*bot, *gid, *uid, text, *req, *random, groupmember, groupfunction)
		break

	default:
		api.Sendgroupmsg(*bot, *uid, "Hi我是Acfur！如果需要帮助请发送acfurhelp", false)
		break
	}
}
