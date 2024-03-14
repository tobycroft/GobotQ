package service

import (
	"regexp"
)

//func Serv_at(qq any) string {
//	if Calc.Any2String(qq) != "0" {
//		return "[CQ:at,qq=" + Calc.Any2String(qq) + "] "
//	}
//	return ""
//}

func Serv_at_all() string {
	return "[CQ:at,qq=all]"
}

func Serv_at_who(contains_at_message string) (string, string) {
	reg, err := regexp.Compile("\\[CQ\\:at\\,qq\\=[0-9]+\\]")
	if err != nil {
		return "", ""
	} else {
		at_str := reg.FindString(contains_at_message)
		if at_str != "" {
			reg, err := regexp.Compile("[0-9]+")
			if err != nil {
				return "", ""
			} else {
				return at_str, reg.FindString(at_str)
			}
		}
		return "", ""
	}
}

func Serv_get_qq(get_qq string) string {
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		return ""
	} else {
		return reg.FindString(get_qq)
	}
}

func Serv_get_num(get_num string) string {
	return Serv_get_qq(get_num)
}
