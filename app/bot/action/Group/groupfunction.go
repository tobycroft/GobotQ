package Group

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"

	"strings"
)

func App_group_function_get_all(self_id, group_id, user_id int64, message string, groupfunction map[string]any) {
	settings := group_function_attach(group_id)
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
	AutoMessage(self_id, group_id, user_id, str, groupfunction)
}

func App_group_function_set(self_id, group_id, user_id int64, message string, groupfunction map[string]any) {
	i1 := strings.Index(message, ":")
	i2 := strings.Index(message, "：")
	if i1 == i2 {
		AutoMessage(self_id, group_id, user_id, "如果需要设定，请使用acfur设定设定内容：设定结果，例如\r\n\"acfur设定入群欢迎：开", groupfunction)
		return
	}
	strs := []string{}
	if (i1 < i2 || i2 == -1) && i1 != -1 {
		strs = strings.Split(message, ":")
	} else {
		strs = strings.Split(message, "：")
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
		var value any
		switch detail["type"].(string) {
		case "bool":
			if set == "开" || set == "是" || set == "on" || set == "1" || set == "true" {
				value = true
			} else {
				value = false
			}
			break

		case "int":
			if len(set) > 0 {
				i, err := Calc.Any2Int_2(set)
				if err != nil {
					AutoMessage(self_id, group_id, user_id, name+"只能设定为数字整数,请调整为数字整数", groupfunction)
					return
				} else {
					value = i
				}
			} else {
				AutoMessage(self_id, group_id, user_id, name+"的设定有误，例子：acfur设定"+name+":"+"数字", groupfunction)
				return
			}
			break

		case "string":
			if len(set) > 0 {
				value = set
			} else {
				AutoMessage(self_id, group_id, user_id, name+"的设定有误，例子：acfur设定"+name+":"+"你要设定的文字", groupfunction)
			}
			break

		default:
			AutoMessage(self_id, group_id, user_id, name+"需要有设定项，你可以使用acfur设定"+name+":"+"设定值，对该功能进行设定", groupfunction)
			return
		}
		if GroupFunctionModel.Api_update(group_id, detail["key"], value) {
			AutoMessage(self_id, group_id, user_id, name+"设定成功为"+":"+set, groupfunction)
		} else {
			AutoMessage(self_id, group_id, user_id, name+"设定失败", groupfunction)
		}
	} else {
		AutoMessage(self_id, group_id, user_id, name+"没有找到对应的设定项目，\r\n如果需要设定，请使用acfur设定设定内容：设定结果，例如\r\nacfur设定入群欢迎：开", groupfunction)
	}
}

func group_function_attach(group_id any) map[string]map[string]any {
	group_setting := GroupFunctionModel.Api_find(group_id)
	if len(group_setting) < 1 {
		GroupFunctionModel.Api_insert(group_id)
		return group_function_attach(group_id)
	}
	function := GroupFunctionDetailModel.Api_select_kv()
	arr := make(map[string]map[string]any)
	for k, v := range group_setting {
		if function[k] != nil {
			function[k]["value"] = v
			arr[k] = function[k]
		}
	}
	return arr
}
