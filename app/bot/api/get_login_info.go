package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

type LoginInfoRet struct {
	Data    LoginInfo `json:"data"`
	Retcode int       `json:"retcode"`
	Status  string    `json:"status"`
}

type LoginInfo struct {
	Nickname string `json:"nickname"`
	UserID   int    `json:"user_id"`
}

func GetLoginInfo(self_id interface{}) (LoginInfo, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return LoginInfo{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/get_login_info", nil, nil, nil, nil)
	if err != nil {
		return LoginInfo{}, err
	}
	var ret LoginInfoRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return LoginInfo{}, err
	}
	if ret.Retcode == 0 {
		return ret.Data, nil
	} else {
		return LoginInfo{}, nil
	}
}