package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

type FriendListRet struct {
	Data    []FriendList `json:"data"`
	Retcode int          `json:"retcode"`
	Status  string       `json:"status"`
}

type FriendList struct {
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	UserID   int64  `json:"user_id"`
}

func Getfriendlist(self_id interface{}) ([]FriendList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/get_friend_list", nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var gfl FriendListRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gfl)
	if err != nil {
		return nil, err
	}
	return gfl.Data, nil
}