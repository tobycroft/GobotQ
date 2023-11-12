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

type FriendListRet struct {
	Data    []FriendList `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
}

type FriendList struct {
	UserId          int64  `json:"user_id"`
	UserName        string `json:"user_name"`
	UserDisplayname string `json:"user_displayname"`
	UserRemark      string `json:"user_remark"`
	Age             int64  `json:"age"`
	Gender          int64  `json:"gender"`
	GroupId         int64  `json:"group_id"`
	Platform        string `json:"platform"`
	TermType        int64  `json:"term_type"`
}

func (api Post) Getfriendlist(self_id any) ([]FriendList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_friend_list", nil, nil, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gfl FriendListRet

	err = sonic.UnmarshalString(data, &gfl)
	if err != nil {
		return nil, err
	}
	return gfl.Data, nil
}
func (api Ws) Getfriendlist(self_id any) ([]FriendList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_friend_list",
		Params: map[string]any{},
		Echo: echo{
			Action: "get_friend_list",
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
