package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type GroupInfoRet struct {
	Data    GroupInfo `json:"data"`
	Retcode int64     `json:"retcode"`
	Status  string    `json:"status"`
}

type GroupInfo struct {
	GroupId        int64   `json:"group_id"`
	GroupName      string  `json:"group_name"`
	GroupRemark    string  `json:"group_remark"`
	GroupUin       int64   `json:"group_uin"`
	Admins         []int64 `json:"admins"`
	ClassText      string  `json:"class_text"`
	IsFrozen       bool    `json:"is_frozen"`
	MaxMember      int64   `json:"max_member"`
	MemberNum      int64   `json:"member_num"`
	MemberCount    int64   `json:"member_count"`
	MaxMemberCount int64   `json:"max_member_count"`
}

func (api Post) GetGroupInfo(self_id, group_id any) (GroupInfo, error) {
	post := map[string]any{
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

	err = sonic.UnmarshalString(data, &ret1)
	if err != nil {
		return GroupInfo{}, err
	}
	if ret1.Retcode == 0 {
		return ret1.Data, nil
	} else {
		return GroupInfo{}, errors.New(ret1.Status)
	}

}
func (api Ws) GetGroupInfo(self_id, group_id any) (GroupInfo, error) {
	post := map[string]any{
		"group_id": group_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GroupInfo{}, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_group_info",
		Params: post,
		Echo: echo{
			Action: "get_group_info",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return GroupInfo{}, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return GroupInfo{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return GroupInfo{}, err
}
