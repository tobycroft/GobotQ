package service

import (
	"regexp"
)

func Serv_text_match(text string, Case []string) (string, bool) {
	for _, str := range Case {
		reg := regexp.MustCompile("(?i)^" + str)
		active := reg.MatchString(text)
		if active {
			new_text := reg.ReplaceAllString(text, "")
			return new_text, true
		}
	}
	return "", false
}
