package service

import "main.go/tuuz/Calc"

func Serv_at(qq interface{}) string {
	return "[@" + Calc.Any2String(qq) + "]"
}

func Serv_at_all() string {
	return "[@all]"
}
