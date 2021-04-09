package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

/*
{
    "ret": "{\"retcode\":0,\"retmsg\":\"修改群名片成功\",\"time\":\"1609565150\"}"
}
*/

type GroupCard struct {
	Ret string `json:"ret"`
}

type GroupCardRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Setgroupcard(fromqq, togroup, toqq, card interface{}) (GroupCard, GroupCardRet, error) {
	post := map[string]interface{}{
		"fromqq":  fromqq,
		"togroup": togroup,
		"toqq":    toqq,
		"card":    card,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/setgroupcard", nil, post, nil, nil)
	if err != nil {
		return GroupCard{}, GroupCardRet{}, err
	}
	var gc GroupCard
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gc)
	if err != nil {
		return GroupCard{}, GroupCardRet{}, err
	}
	var ret GroupCardRet
	err = jsr.UnmarshalFromString(gc.Ret, &ret)
	if err != nil {
		return gc, GroupCardRet{}, err
	}
	if ret.Retcode != 0 {
		return gc, ret, errors.New(ret.Retmsg)
	}
	return gc, ret, nil
}
