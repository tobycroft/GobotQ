package service

import "main.go/tuuz/Calc"

func Serv_at(qq int) string {
	return "[@" + Calc.Int2String(qq) + "]"
}
