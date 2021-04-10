package service

import (
	"regexp"
	"strings"
)

func Serv_url_detect(text string) bool {
	c2, _ := regexp.MatchString("\\.co+|\\.cn+|\\.ltd+|\\.xyz+|\\.net+|\\.top+|\\.tech+|\\.org+|\\.gov+|\\.edu+|\\.ink+|\\.int+|\\.pub+|\\.cc+|\\.info+|\\.io+", text)
	if c2 {
		if strings.Contains(text, "http") {
			return true
		} else if strings.Contains(text, "?") {
			if strings.Contains(text, "=") {
				return true
			}
		}
	}
	return false
}
