package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type UserInfoRet struct {
	Data    UserInfo `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
}

type UserInfo struct {
	Age       int64  `json:"age"`
	Level     int64  `json:"level"`
	LoginDays int64  `json:"login_days"`
	Nickname  string `json:"nickname"`
	Qid       string `json:"qid"`
	Sex       string `json:"sex"`
	UserID    int64  `json:"user_id"`
}

func (api Post) GetStrangerInfo(self_id, user_id any, no_cache bool) (UserInfo, error) {
	post := map[string]any{
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

	err = sonic.UnmarshalString(data, &ret1)
	if err != nil {
		return UserInfo{}, err
	}
	return ret1.Data, nil
}
func (api Ws) GetStrangerInfo(self_id, user_id any, no_cache bool) (UserInfo, error) {
	post := map[string]any{
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

	err = sonic.UnmarshalString(data, &ret1)
	if err != nil {
		return UserInfo{}, err
	}
	return ret1.Data, nil
}
