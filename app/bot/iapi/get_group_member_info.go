package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type GroupMemberInfoRet struct {
	Data    GroupMemberList `json:"data"`
	Retcode int64           `json:"retcode"`
	Status  string          `json:"status"`
}

func (api Api) GetGroupMemberInfo(self_id, group_id, user_id any) (GroupMemberList, error) {
	post := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"no_cache": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GroupMemberList{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_member_info", nil, post, nil, nil).RetString()
	if err != nil {
		return GroupMemberList{}, err
	}
	var gms GroupMemberInfoRet

	err = sonic.UnmarshalString(data, &gms)
	if err != nil {
		return GroupMemberList{}, err
	}
	if gms.Retcode == 0 {
		return gms.Data, nil
	} else {
		return GroupMemberList{}, nil
	}
}
