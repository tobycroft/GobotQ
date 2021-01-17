package service

import (
	"strings"
)

func Serv_ban_share(text string) bool {
	c1 := strings.Contains(text, "com.tencent.")
	if c1 {
		return c1
	}
	if strings.Contains(text, "<?xml") {
		if strings.Contains(text, "web") {
			return true
		}
	}
	return false
}
