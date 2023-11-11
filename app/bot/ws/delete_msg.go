package api

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Log"
)

type Struct_Retract struct {
	Self_id   any
	MessageId any
}

var Retract_chan = make(chan Struct_Retract, 100)
var Retract_instant = make(chan Struct_Retract, 100)

type DefaultRetStruct struct {
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
	Retcode int64  `json:"retcode"`
	Status  string `json:"status"`
	Wording string `json:"wording"`
}

func (ws Ws) DeleteMsg(self_id, message_id any) (bool, error) {
	post := map[string]any{
		"message_id": message_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/delete_msg", nil, post, nil, nil).RetString()
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
