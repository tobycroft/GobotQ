package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

type UserInfoRet struct {
	Data    UserInfo `json:"data"`
	Retcode int      `json:"retcode"`
	Status  string   `json:"status"`
}

type UserInfo struct {
	Age       int    `json:"age"`
	Level     int    `json:"level"`
	LoginDays int    `json:"login_days"`
	Nickname  string `json:"nickname"`
	Qid       string `json:"qid"`
	Sex       string `json:"sex"`
	UserID    int    `json:"user_id"`
}

func GetStrangerInfo(self_id, user_id interface{}, no_cache bool) (UserInfo, error) {
	post := map[string]interface{}{
		"user_id":  user_id,
		"no_cache": no_cache,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return UserInfo{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_stranger_info", nil, post, nil, nil).RetString()
	if err != nil {
		return UserInfo{}, err
	}
	var ret1 UserInfoRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return UserInfo{}, err
	}
	return ret1.Data, nil
}
