package service

import (
	"main.go/tuuz/Calc"
	"strings"
)

func Serv_at(qq interface{}) string {
	return "[CQ:at,qq=" + Calc.Any2String(qq) + "]"
}

func Serv_at_all() string {
	return "[CQ:at,qq=all]"
}

func Serv_is_at_me(self_id interface{}, message string) bool {
	return strings.Contains(message, Serv_at(self_id))
}
