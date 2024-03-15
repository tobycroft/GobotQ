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

func (api Post) SetFriendAddRequest(self_id int64, flag any, approve bool, remark any) (bool, error) {
	post := map[string]any{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.New().PostUrlXEncode(botinfo["url"].(string)+"/set_friend_add_request", nil, post, nil, nil).RetString()
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

// SetFriendAddRequest not yet done
func (api Ws) SetFriendAddRequest(self_id int64, flag any, approve bool, remark any) (bool, error) {
	post := map[string]any{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "set_friend_add_request",
		Params: post,
		Echo: echo{
			Action: "set_friend_add_request",
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
