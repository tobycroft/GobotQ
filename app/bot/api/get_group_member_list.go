package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

type GroupMemberListRet struct {
	Data    []GroupMemberList `json:"data"`
	Retcode int               `json:"retcode"`
	Status  string            `json:"status"`
}

type GroupMemberList struct {
	Age             int    `json:"age"`
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

func Getgroupmemberlist(self_id, group_id interface{}) ([]GroupMemberList, error) {
	post := map[string]interface{}{
		"group_id": group_id,
		"no_cache": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/get_group_member_list", nil, post, nil, nil)
	if err != nil {
		return nil, err
	}
	var gms GroupMemberListRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gms)
	if err != nil {
		return nil, err
	}
	return gms.Data, nil
}
