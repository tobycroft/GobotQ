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

type LoginInfoRet struct {
	Data    LoginInfo `json:"data"`
	Retcode int64     `json:"retcode"`
	Status  string    `json:"status"`
}

type LoginInfo struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
}

func (api Post) GetLoginInfo(self_id int64) (LoginInfo, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return LoginInfo{}, errors.New("botinfo_notfound")
	}
	//Net.WsServer_WriteChannel <- Net.WsData{}
	data, err := Net.Post{}.New().PostUrlXEncode(botinfo["url"].(string)+"/get_login_info", nil, nil, nil, nil).RetString()
	if err != nil {
		return LoginInfo{}, err
	}
	var ret LoginInfoRet
	err = sonic.UnmarshalString(data, &ret)
	if err != nil {
		return LoginInfo{}, err
	}
	if ret.Retcode == 0 {
		return ret.Data, nil
	} else {
		return LoginInfo{}, nil
	}
}

func (api Ws) GetLoginInfo(self_id int64) (LoginInfo, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return LoginInfo{}, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_login_info",
		Params: map[string]any{},
		Echo: echo{
			Action: "get_login_info",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return LoginInfo{}, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return LoginInfo{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return LoginInfo{}, nil
}
