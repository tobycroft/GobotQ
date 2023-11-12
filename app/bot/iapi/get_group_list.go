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

type GroupListRet struct {
	Data    []GroupList `json:"data"`
	Retcode int64       `json:"retcode"`
	Status  string      `json:"status"`
}

type GroupList struct {
	GroupId        int    `json:"group_id"`
	GroupName      string `json:"group_name"`
	GroupRemark    string `json:"group_remark"`
	GroupUin       int    `json:"group_uin"`
	Admins         []int  `json:"admins"`
	ClassText      string `json:"class_text"`
	IsFrozen       bool   `json:"is_frozen"`
	MaxMember      int    `json:"max_member"`
	MemberNum      int    `json:"member_num"`
	MemberCount    int    `json:"member_count"`
	MaxMemberCount int    `json:"max_member_count"`
}

func (api Post) Getgrouplist(self_id any) ([]GroupList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_list", nil, nil, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gls GroupListRet

	err = sonic.UnmarshalString(data, &gls)
	if err != nil {
		return nil, err
	}
	return gls.Data, nil
}
func (api Ws) Getgrouplist(self_id any) ([]GroupList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_group_list",
		Params: nil,
		Echo: echo{
			Action: "get_group_list",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return nil, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return nil, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return nil, err
}
