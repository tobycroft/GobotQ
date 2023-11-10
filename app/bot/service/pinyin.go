package service

import (
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
)

type py struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   string `json:"data"`
}

func Serv_pinyin(chinese interface{}) (string, error) {
	data, err := Net.Post{}.PostRaw("http://www.box3.cn/developtoolbox/pinyin.php", Calc.Any2String(chinese)).RetString()
	if err != nil {
		return "", err
	}
	var ret py

	err = sonic.UnmarshalString(data, &ret)
	if err != nil {
		return "", err
	}
	return ret.Data, nil
}
