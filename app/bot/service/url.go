package service

import "regexp"

func Serv_url_detect(text string) bool {
	c1, _ := regexp.MatchString("[http].*.[com|cn|ltd|xyz|net|xyz|top|tech|org|gov|edu|ink|int|pub|cc|info|io]", text)
	return c1
}
