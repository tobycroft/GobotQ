package service

import "regexp"

func Serv_url_detect(text string) (string, bool) {
	str := "[http].*.[com|cn|ltd|xyz|net|xyz|top|tech|org|gov|edu|ink|int|pub|cc|info|io]"
	reg := regexp.MustCompile(str)
	active := reg.MatchString(text)
	if active {
		new_text := reg.ReplaceAllString(text, "")
		return new_text, true
	}
}
