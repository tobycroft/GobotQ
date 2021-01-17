package service

import (
	"regexp"
	"strings"
)

func Serv_url_detect(text string) bool {
	if strings.Contains(text, "http") {
		c2, _ := regexp.MatchString("\\.co+|\\.cn+|\\.ltd+|\\.xyz+|\\.net+|\\.xyz+|\\.top+|\\.tech+|\\.org+|\\.gov+|\\.edu+|\\.ink+|\\.int+|\\.pub+|\\.cc+|\\.info+|\\.io+", text)
		return c2
	}
	return false
}
