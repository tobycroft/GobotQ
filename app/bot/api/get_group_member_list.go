package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Net"
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
	GroupID         int    `json:"group_id"`
	JoinTime        int    `json:"join_time"`
	LastSentTime    int    `json:"last_sent_time"`
	Level           string `json:"level"`
	Nickname        string `json:"nickname"`
	Role            string `json:"role"`
	Sex             string `json:"sex"`
	Title           string `json:"title"`
	TitleExpireTime int    `json:"title_expire_time"`
	Unfriendly      bool   `json:"unfriendly"`
	UserID          int    `json:"user_id"`
}

func Getgroupmemberlist(self_id, group_id interface{}) ([]GroupMemberList, error) {
	post := map[string]interface{}{
		"group_id": group_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
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
