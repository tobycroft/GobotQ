package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

/*
{
    "ret": "true",
    "Info": {
        "Name": "云机器人小组",
        "Pos": "西洪小区",
        "Type": "IT/互联网",
        "Tag": "火线兔|头条程序员|B站程序员|狙神|哔哩哔哩助手|golang|个人狙|逆战火线兔|头条程序员|B站程序员|狙神|哔哩哔哩助手|golang|个人狙|逆战开发|技术|计算机|运维|云机器人|acfur机器人",
        "Introduction": "Acfur云机器人群"
    }
}
*/

type QueryGroupInfoRet struct {
	Ret  string         `json:"ret"`
	Info QueryGroupInfo `json:"Info"`
}

type QueryGroupInfo struct {
	Name         string `json:"Name"`
	Pos          string `json:"Pos"`
	Type         string `json:"Type"`
	Tag          string `json:"Tag"`
	Introduction string `json:"Introduction"`
}

func Querygroupinfo(logonqq, qq interface{}) (QueryGroupInfo, error) {
	post := map[string]interface{}{
		"logonqq": logonqq,
		"qq":      qq,
	}
	data, err := Net.Post(app_conf.Http_Api+"/querygroupinfo", nil, post, nil, nil)
	if err != nil {
		return QueryGroupInfo{}, err
	}
	var ret1 QueryGroupInfoRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return QueryGroupInfo{}, err
	}
	return ret1.Info, nil
}
