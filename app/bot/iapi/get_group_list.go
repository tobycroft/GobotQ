package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type GroupListRet struct {
	Data    []GroupList `json:"data"`
	Retcode int64       `json:"retcode"`
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

func (api Api) Getgrouplist(self_id any) ([]GroupList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_list", nil, nil, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gls GroupListRet

	err = sonic.UnmarshalString(data, &gls)
	if err != nil {
		return nil, err
	}
	return gls.Data, nil
}
func (api Ws) Getgrouplist(self_id any) ([]GroupList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_list", nil, nil, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gls GroupListRet

	err = sonic.UnmarshalString(data, &gls)
	if err != nil {
		return nil, err
	}
	return gls.Data, nil
}
