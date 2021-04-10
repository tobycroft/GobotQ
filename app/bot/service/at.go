package service

import "main.go/tuuz/Calc"

func Serv_at(qq interface{}) string {
	return "[CQ:at,qq=" + Calc.Any2String(qq) + "]"
}

func Serv_at_all() string {
	return "[CQ:at,qq=all]"
}
