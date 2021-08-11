package service

import (
	"regexp"
)

func Serv_text_match_any(text string, Case []string) (string, bool) {
	for _, str := range Case {
		reg := regexp.MustCompile(str)
		active := reg.MatchString(text)
		if active {
			new_text := reg.ReplaceAllString(text, "")
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

func Serv_text_match_define(text *string, Case []string) bool {
	for _, str := range Case {
		reg := regexp.MustCompile("(?i)^" + str)
		active := reg.MatchString(*text)
		if active {
			*text = reg.ReplaceAllString(*text, "")
			return true
		}
	}
	return false
}
