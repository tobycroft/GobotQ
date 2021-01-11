package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/tuuz/Calc"
	"strings"
)

func App_group_function_get_all(bot *int, gid *int, uid *int, text *string) {
	settings := group_function_attach(*gid)
	str := "您的群设定为：\r\n"
	for _, v := range settings {
		str += Calc.Any2String(v["name"]) + ":"
		value := Calc.Any2String(v["value"])
		switch v["type"] {
		case "bool":
			if value == "1" {
				str += "是"
			} else {
				str += "否"
			}
			break

		case "int":
			str += value
			break

		case "string":
			str += value
			break

		default:
			str += value
			break
		}

		str += "\r\n"
	}
	api.Sendgroupmsg(*bot, *gid, str, true)
}

func App_group_function_set(bot, gid, uid interface{}, text string, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	i1 := strings.Index(text, ":")
	i2 := strings.Index(text, "：")
	if i1 == i2 {
		api.Sendgroupmsg(bot, gid, "如果需要设定，请使用acfur设定设定内容：设定结果，例如\r\n\"acfur设定入群欢迎：开", true)
		return
	}
	strs := []string{}
	if i1 < i2 {
		strs = strings.Split(text, ":")
	} else {
		strs = strings.Split(text, "：")
	}
	name := ""
	set := ""
	for k, v := range strs {
		if k == 0 {
			name = v
		} else {
			set += v
		}
	}
	detail := GroupFunctionDetailModel.Api_find_byName(name)
	if len(detail) > 0 {
		var value interface{}
		switch detail["type"].(string) {
		case "bool":
			if set == "开" || set == "是" || set == "on" || set == "1" || set == "true" {
				value = "开"
			} else {
				value = "关"
			}
			break

		case "int":
			if len(set) > 0 {
				i, err := Calc.Any2Int_2(set)
				if err != nil {
					api.Sendgroupmsg(bot, gid, name+"只能设定为数字整数,请调整为数字整数", true)
					return
				} else {
					value = i
				}
			} else {
				api.Sendgroupmsg(bot, gid, name+"的设定有误，例子：acfur设定"+name+":"+"数字", true)
				return
			}
			break

		case "string":
			if len(set) > 0 {
				value = set
			} else {
				api.Sendgroupmsg(bot, gid, name+"的设定有误，例子：acfur设定"+name+":"+"你要设定的文字", true)
			}
			break

		default:
			api.Sendgroupmsg(bot, gid, name+"需要有设定项，你可以使用acfur设定"+name+":"+"设定值，对该功能进行设定", true)
			return
		}
		if GroupFunctionModel.Api_update(gid, name, value) {
			api.Sendgroupmsg(bot, gid, name+"设定成功为"+":"+set, true)
		} else {
			api.Sendgroupmsg(bot, gid, name+"设定失败", true)
		}
	} else {
		api.Sendgroupmsg(bot, gid, name+"没有找到对应的设定项目，\r\n如果需要设定，请使用acfur设定设定内容：设定结果，例如\r\n\"acfur设定入群欢迎：开", true)
	}
}

func group_function_attach(gid interface{}) map[string]map[string]interface{} {
	group_setting := GroupFunctionModel.Api_find(gid)
	if len(group_setting) < 1 {
		GroupFunctionModel.Api_insert(gid)
		return group_function_attach(gid)
	}
	function := GroupFunctionDetailModel.Api_select_kv()
	arr := make(map[string]map[string]interface{})
	for k, v := range group_setting {
		if function[k] != nil {
			function[k]["value"] = v
			arr[k] = function[k]
		}
	}
	return arr
}
