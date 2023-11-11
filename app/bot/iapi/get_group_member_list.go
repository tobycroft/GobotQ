package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type GroupMemberListRet struct {
	Data    []GroupMemberList `json:"data"`
	Retcode int64             `json:"retcode"`
	Status  string            `json:"status"`
}

type GroupMemberList struct {
	Age             int64  `json:"age"`
	Area            string `json:"area"`
	Card            string `json:"card"`
	CardChangeable  bool   `json:"card_changeable"`
	GroupID         int64  `json:"group_id"`
	JoinTime        int64  `json:"join_time"`
	LastSentTime    int64  `json:"last_sent_time"`
	Level           string `json:"level"`
	Nickname        string `json:"nickname"`
	Role            string `json:"role"`
	Sex             string `json:"sex"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	Unfriendly      bool   `json:"unfriendly"`
	UserID          int64  `json:"user_id"`
}

func (api Post) Getgroupmemberlist(self_id, group_id any) ([]GroupMemberList, error) {
	post := map[string]any{
		"group_id": group_id,
		"no_cache": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_member_list", nil, post, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gms GroupMemberListRet

	err = sonic.UnmarshalString(data, &gms)
	if err != nil {
		return nil, err
	}
	return gms.Data, nil
}
func (api Ws) Getgroupmemberlist(self_id, group_id any) ([]GroupMemberList, error) {
	post := map[string]any{
		"group_id": group_id,
		"no_cache": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_member_list", nil, post, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gms GroupMemberListRet

	err = sonic.UnmarshalString(data, &gms)
	if err != nil {
		return nil, err
	}
	return gms.Data, nil
}
