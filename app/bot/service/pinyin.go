package service

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Net"
)

type py struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   string `json:"data"`
}

func Serv_pinyin(chinese interface{}) (string, error) {
	data, err := Net.PostRaw("http://www.box3.cn/developtoolbox/pinyin.php", Calc.Any2String(chinese))
	if err != nil {
		return "", err
	}
	var ret py
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return "", err
	}
	return ret.Data, nil
}
