package service

import (
	"regexp"
)

func Serv_ban_group(text string) bool {
	c1, _ := regexp.MatchString("qun+|群|裙|君羊+", text)
	if c1 {
		c2, _ := regexp.MatchString("[0-8]\\d{8}", text)
		if c2 {
			return true
		}
	}
	return false
}
