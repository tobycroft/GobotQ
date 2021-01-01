package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type Gms struct {
	Ret  string          `json:"ret"`
	List GroupMemberList `json:"List"`
}

type GroupMemberList []struct {
	UIN              int    `json:"UIN"`
	Age              int    `json:"Age"`
	Sex              int    `json:"Sex"`
	NickName         string `json:"NickName"`
	Email            string `json:"Email"`
	Card             string `json:"Card"`
	Remark           string `json:"Remark"`
	SpecTitle        string `json:"SpecTitle"`
	Phone            string `json:"Phone"`
	SpecTitleExpired int    `json:"SpecTitleExpired"`
	MuteTime         int    `json:"MuteTime"`
	AddGroupTime     int    `json:"AddGroupTime"`
	LastMsgTime      int    `json:"LastMsgTime"`
	GroupLevel       int    `json:"GroupLevel"`
}

func Getgroupmemberlist(bot, group interface{}) (GroupMemberList, error) {
	post := map[string]interface{}{
		"logonqq": bot,
		"group":   group,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getgroupmemberlist", nil, post, nil, nil)
	if err != nil {
		return nil, err
	}
	var gms Gms
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gms)
	if err != nil {
		return nil, err
	}
	return gms.List, nil
}
