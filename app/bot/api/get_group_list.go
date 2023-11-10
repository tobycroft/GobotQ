package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

type GroupListRet struct {
	Data    []GroupList `json:"data"`
	Retcode int         `json:"retcode"`
	Status  string      `json:"status"`
}

type GroupList struct {
	GroupCreateTime int64  `json:"group_create_time"`
	GroupID         int64  `json:"group_id"`
	GroupLevel      int64  `json:"group_level"`
	GroupMemo       string `json:"group_memo"`
	GroupName       string `json:"group_name"`
	MaxMemberCount  int64  `json:"max_member_count"`
	MemberCount     int64  `json:"member_count"`
}

func Getgrouplist(self_id interface{}) ([]GroupList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/get_group_list", nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var gls GroupListRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gls)
	if err != nil {
		return nil, err
	}
	return gls.Data, nil
}
