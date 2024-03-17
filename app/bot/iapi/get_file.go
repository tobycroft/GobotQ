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

type GetFile struct {
	Data    GetFileData `json:"data"`
	Retcode int64       `json:"retcode"`
	Status  string      `json:"status"`
}

type GetFileData struct {
	File         string `json:"file"`
	Base64String string `json:"base64String"`
	Md5          string `json:"md5"`
}

func (api Post) GetFile(self_id int64, file string, Type string) (GetFileData, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GetFileData{}, errors.New("botinfo_notfound")
	}
	//Net.WsServer_WriteChannel <- Net.WsData{}
	data, err := Net.Post{}.New().PostUrlXEncode(botinfo["url"].(string)+"/get_file", map[string]interface{}{
		"file": file,
	}, map[string]any{
		"file_type": "base64",
	}, nil, nil).RetString()
	if err != nil {
		return GetFileData{}, err
	}
	var ret GetFile
	err = sonic.UnmarshalString(data, &ret)
	if err != nil {
		return GetFileData{}, err
	}
	if ret.Retcode == 0 {
		return ret.Data, nil
	} else {
		return GetFileData{}, nil
	}
}

func (api Ws) GetFile(self_id int64, file string, OrginalFile string) (GetFileData, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GetFileData{}, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_file",
		Params: map[string]any{
			"file":      file,
			"file_type": "base64",
		},
		Echo: echo{
			Action: "get_file",
			SelfId: Calc.Any2Int64(self_id),
			Extra:  OrginalFile,
		},
	})
	if err != nil {
		return GetFileData{}, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return GetFileData{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return GetFileData{}, nil
}
