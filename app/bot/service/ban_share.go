package service

import (
	"strings"
)

func Serv_ban_share(text string) bool {
	c1 := strings.Contains(text, "[CQ:json,")
	if c1 {
		return c1
	}
	if strings.Contains(text, "<?xml") {
		if strings.Contains(text, "serviceID=\"104\" templateID=\"1\"") {
			return false
		}
		return true
	}
}
