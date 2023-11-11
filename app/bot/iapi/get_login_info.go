package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
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

func (api Api) GetLoginInfo(self_id any) (LoginInfo, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return LoginInfo{}, errors.New("botinfo_notfound")
	}
	//Net.WsServer_WriteChannel <- Net.WsData{}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_login_info", nil, nil, nil, nil).RetString()
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
func (api Ws) GetLoginInfo(self_id any) (LoginInfo, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return LoginInfo{}, errors.New("botinfo_notfound")
	}
	//Net.WsServer_WriteChannel <- Net.WsData{}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_login_info", nil, nil, nil, nil).RetString()
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
