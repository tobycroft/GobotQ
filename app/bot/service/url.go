package service

import "regexp"

func Serv_url_detect(text string) bool {
	c1, _ := regexp.MatchString("[http]+", text)
	if c1 {
		c2, _ := regexp.MatchString("[com|cn|ltd|xyz|net|xyz|top|tech|org|gov|edu|ink|int|pub|cc|info|io]", text)
		return c2
	}
	return c1
}
