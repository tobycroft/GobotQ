package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type GroupListRet struct {
	Data    []GroupList `json:"data"`
	Retcode int         `json:"retcode"`
	Status  string      `json:"status"`
}

type GroupList struct {
	GroupCreateTime int    `json:"group_create_time"`
	GroupID         int    `json:"group_id"`
	GroupLevel      int    `json:"group_level"`
	GroupMemo       string `json:"group_memo"`
	GroupName       string `json:"group_name"`
	MaxMemberCount  int    `json:"max_member_count"`
	MemberCount     int    `json:"member_count"`
}

func Getgrouplist(bot interface{}) ([]GroupList, error) {
	post := map[string]interface{}{
		"logonqq": bot,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getgrouplist", nil, post, nil, nil)
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
