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

func (api Post) DeleteFriend(self_id, friend_id any) (bool, error) {
	post := map[string]any{
		"friend_id": friend_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/delete_friend", nil, post, nil, nil).RetString()
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
		Log.Crrs(errors.New(dls.Wording), "message:"+Calc.Any2String(friend_id))
		return false, errors.New(dls.Wording)
	}
}
func (api Ws) DeleteFriend(self_id, user_id any) (bool, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	post := map[string]any{
		"user_id": user_id,
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "delete_friend",
		Params: post,
		Echo: echo{
			Action: "delete_friend",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return false, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return false, err
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return true, nil
}
