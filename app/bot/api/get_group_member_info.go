package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

type GroupMemberInfoRet struct {
	Data    GroupMemberList `json:"data"`
	Retcode int             `json:"retcode"`
	Status  string          `json:"status"`
}

func GetGroupMemberInfo(self_id, group_id, user_id interface{}) (GroupMemberList, error) {
	post := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GroupMemberList{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/get_group_member_info", nil, post, nil, nil)
	if err != nil {
		return GroupMemberList{}, err
	}
	var gms GroupMemberInfoRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gms)
	if err != nil {
		return GroupMemberList{}, err
	}
	if gms.Retcode == 0 {
		return gms.Data, nil
	} else {
		return GroupMemberList{}, nil
	}
}
