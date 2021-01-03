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
        "UIN": 710209521,
        "NickName": "修水熊伟伟",
        "Remark": "恒邦保险熊",
        "Email": "",
        "OnlineState": "",
        "CLike": 0,
        "Sign": "恒邦保险熊",
        "QQLevel": 26,
        "Age": 33,
        "Country": "中国",
        "Province": "江西",
        "City": "九江",
        "SvcList": [],
        "Group": "",
        "TodayLiked": 0,
        "TodayCLike": 10
    }
}
*/

type QueryUserInfoRet struct {
	Ret  string        `json:"ret"`
	Info QueryUserInfo `json:"Info"`
}
type QueryUserInfo struct {
	UIN         int           `json:"UIN"`
	NickName    string        `json:"NickName"`
	Remark      string        `json:"Remark"`
	Email       string        `json:"Email"`
	OnlineState string        `json:"OnlineState"`
	CLike       int           `json:"CLike"`
	Sign        string        `json:"Sign"`
	QQLevel     int           `json:"QQLevel"`
	Age         int           `json:"Age"`
	Country     string        `json:"Country"`
	Province    string        `json:"Province"`
	City        string        `json:"City"`
	SvcList     []interface{} `json:"SvcList"`
	Group       string        `json:"Group"`
	TodayLiked  int           `json:"TodayLiked"`
	TodayCLike  int           `json:"TodayCLike"`
}

func Queryuserinfo(logonqq, qq interface{}) (QueryUserInfo, error) {
	post := map[string]interface{}{
		"logonqq": logonqq,
		"qq":      qq,
	}
	data, err := Net.Post(app_conf.Http_Api+"/queryuserinfo", nil, post, nil, nil)
	if err != nil {
		return Info{}, err
	}
	var ret1 QueryUserInfoRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return Info{}, err
	}
	return ret1.Info, nil
}
