package service

import "regexp"

func Serv_ban_weixin(text string) bool {
	c1, _ := regexp.MatchString("wx+|微|vx+|vv+", text)
	if c1 {
		c2, _ := regexp.MatchString("信|\\+", text)
		if c2 {
			c3, _ := regexp.MatchString("[a-zA-Z\\d_]{5,}", text)
			return c3
		}
	}
	return false
}
