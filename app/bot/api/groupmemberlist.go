package api

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type Gms struct {
	Ret  string            `json:"ret"`
	List []GroupMemberList `json:"List" bson:"List"`
}
type GroupMemberList struct {
	UIN              int    `json:"UIN"`
	Age              int    `json:"Age"`
	Sex              int    `json:"Sex"`
	NickName         string `json:"NickName" bson:"NickName"`
	Email            string `json:"Email" bson:"Email"`
	Card             string `json:"Card" bson:"Card"`
	Remark           string `json:"Remark" bson:"Remark"`
	SpecTitle        string `json:"SpecTitle" bson:"SpecTitle"`
	Phone            string `json:"Phone"`
	SpecTitleExpired int    `json:"SpecTitleExpired"`
	MuteTime         int    `json:"MuteTime"`
	AddGroupTime     int    `json:"AddGroupTime"`
	LastMsgTime      int    `json:"LastMsgTime"`
	GroupLevel       int    `json:"GroupLevel"`
}

func Getgroupmemberlist(bot, group interface{}) (interface{}, error) {
	post := map[string]interface{}{
		"logonqq": bot,
		"group":   group,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getgroupmemberlist", nil, post, nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	var gms Gms
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gms)
	if err != nil {
		return nil, err
	}
	return gms.List, nil
}
