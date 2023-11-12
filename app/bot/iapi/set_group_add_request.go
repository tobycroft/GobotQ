package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz"

	"main.go/tuuz/Log"
)

func (api Post) SetGroupAddRequestRet(self_id, flag, sub_type any, approve bool, reason string) (bool, error) {
	post := map[string]any{
		"flag":     flag,
		"sub_type": sub_type,
		"type":     sub_type,
		"approve":  approve,
		"reason":   reason,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/set_group_add_request", nil, post, nil, nil).RetString()
	if err != nil {
		return false, err
	}
	var dls DefaultRetStruct

	err = sonic.UnmarshalString(data, &dls)
	if err != nil {
		return false, err
	}
	if dls.Retcode == 0 {
		return true, nil
	} else {
		Log.Crrs(errors.New(dls.Wording), tuuz.FUNCTION_ALL())
		return false, errors.New(dls.Wording)
	}
}
func (api Ws) SetGroupAddRequestRet(self_id, flag, sub_type any, approve bool, reason string) (bool, error) {
	post := map[string]any{
		"flag":     flag,
		"sub_type": sub_type,
		"type":     sub_type,
		"approve":  approve,
		"reason":   reason,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "set_group_add_request",
		Params: post,
		Echo: echo{
			Action: "set_group_add_request",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return true, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return true, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return true, err
}
