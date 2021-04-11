package service

import (
	"strings"
)

func Serv_ban_share(text string) bool {
	c1 := strings.Contains(text, "[CQ:json,")
	if c1 {
		return c1
	}
	return strings.Contains(text, "<?xml")
}
