package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/tuuz/Input"
	"main.go/tuuz/Net"
	"strings"
)

type GroupAdminList struct {
	Ret string `json:"ret"`
}

func Getgroupmgrlist(bot, group interface{}) (map[string]interface{}, error) {
	post := map[string]interface{}{
		"fromqq": bot,
		"group":  group,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/getgroupmgrlist", nil, post, nil, nil)
	if err != nil {
		return nil, err
	}
	data = Input.Fliter_Ascii(data)
	data = Input.Fliter_error_encode(data)
	var ret GroupAdminList
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return nil, err
	}
	strs := strings.Split(ret.Ret, "\r\n")
	arr := make(map[string]interface{})
	for _, v := range strs {
		arr[v] = true
	}
	return arr, nil
}
