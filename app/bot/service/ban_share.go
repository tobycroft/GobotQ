package service

import "regexp"

func Serv_ban_share(text string) bool {
	c1, _ := regexp.MatchString("com.tencent.miniapp", text)
	if c1 {
		return c1
	}
	return false
}
