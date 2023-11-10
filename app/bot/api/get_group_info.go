package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
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

type GroupInfoRet struct {
	Data    GroupInfo `json:"data"`
	Retcode int       `json:"retcode"`
	Status  string    `json:"status"`
}

type GroupInfo struct {
	GroupCreateTime int    `json:"group_create_time"`
	GroupID         int    `json:"group_id"`
	GroupLevel      int    `json:"group_level"`
	GroupMemo       string `json:"group_memo"`
	GroupName       string `json:"group_name"`
	MaxMemberCount  int    `json:"max_member_count"`
	MemberCount     int    `json:"member_count"`
}

func GetGroupInfo(self_id, group_id interface{}) (GroupInfo, error) {
	post := map[string]interface{}{
		"group_id": group_id,
		"no_cache": false,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GroupInfo{}, errors.New("botinfo_notfound")
	}

	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_info", nil, post, nil, nil).RetString()
	if err != nil {
		return GroupInfo{}, err
	}
	var ret1 GroupInfoRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return GroupInfo{}, err
	}
	if ret1.Retcode == 0 {
		return ret1.Data, nil
	} else {
		return GroupInfo{}, errors.New(ret1.Status)
	}

}
