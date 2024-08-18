package service

import (
	"regexp"
	"strings"
)

func Serv_text_match_any(text string, Case []string) (string, bool) {
	for _, str := range Case {
		active := strings.Contains(text, str)
		if active {
			new_text := strings.ReplaceAll(text, str, "")
			return new_text, true
		}
	}
	return "", false
}

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

func Serv_text_match_all(text string, Case []string) (string, bool) {
	for _, str := range Case {
		reg := regexp.MustCompile("(?i)^" + str + "$")
		active := reg.MatchString(text)
		if active {
			new_text := reg.ReplaceAllString(text, "")
			return new_text, true
		}
	}
	return "", false
}
