package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"log"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"reflect"
	"time"
)

//type Struct_Retract struct {
//	SelfId    any
//	MessageId any
//}

type RetractMessage struct {
	SelfId    int64         `json:"selfId"`
	MessageId int64         `json:"messageId"`
	Time      time.Duration `json:"time"`
}

type DefaultRetStruct struct {
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
	Retcode int64  `json:"retcode"`
	Status  string `json:"status"`
	Wording string `json:"wording"`
}
type RetractWsRetStruct struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Data    struct {
	} `json:"data"`
}

func (api Post) DeleteMsg(self_id, message_id int64) (bool, error) {
	post := map[string]any{
		"message_id": message_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.New().PostUrlXEncode(botinfo["url"].(string)+"/delete_msg", nil, post, nil, nil).RetString()
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
		Log.Crrs(errors.New(dls.Wording), "message:"+Calc.Any2String(message_id))
		return false, errors.New(dls.Wording)
	}
}
func (api Ws) DeleteMsg(self_id, message_id int64) (bool, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	post := map[string]any{
		"message_id": message_id,
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "delete_msg",
		Params: post,
		Echo: echo{
			Action: "delete_msg",
			SelfId: Calc.Any2Int64(self_id),
			Extra:  message_id,
		},
	})
	if err != nil {
		return false, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		log.Println(tuuz.FUNCTION_ALL(), "ClientNotFound", self_id, reflect.TypeOf(self_id))
		return false, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}

	return true, nil
}
