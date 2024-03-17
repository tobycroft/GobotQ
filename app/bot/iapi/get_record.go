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

type GetRecord struct {
	Data    GetRecordData `json:"data"`
	Retcode int64         `json:"retcode"`
	Status  string        `json:"status"`
}

type GetRecordData struct {
	File string `json:"file"`
	Url  string `json:"url"`
}

func (api Post) GetRecord(self_id int64, file string, out_format string) (GetRecordData, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GetRecordData{}, errors.New("botinfo_notfound")
	}
	//Net.WsServer_WriteChannel <- Net.WsData{}
	data, err := Net.Post{}.New().PostUrlXEncode(botinfo["url"].(string)+"/get_record", map[string]interface{}{
		"file": file,
	}, map[string]any{
		"out_format": out_format,
	}, nil, nil).RetString()
	if err != nil {
		return GetRecordData{}, err
	}
	var ret GetRecord
	err = sonic.UnmarshalString(data, &ret)
	if err != nil {
		return GetRecordData{}, err
	}
	if ret.Retcode == 0 {
		return ret.Data, nil
	} else {
		return GetRecordData{}, nil
	}
}

func (api Ws) GetRecord(self_id int64, file string, out_format string) (GetRecordData, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return GetRecordData{}, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_record",
		Params: map[string]any{
			"file":       file,
			"out_format": out_format,
		},
		Echo: echo{
			Action: "get_record",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return GetRecordData{}, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return GetRecordData{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return GetRecordData{}, nil
}
