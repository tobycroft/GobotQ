package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

type Struct_Retract struct {
	Self_id   interface{}
	MessageId interface{}
}

var Retract_chan = make(chan Struct_Retract, 100)
var Retract_chan_instant = make(chan Struct_Retract, 100)

type DefaultRetStruct struct {
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Retcode int         `json:"retcode"`
	Status  string      `json:"status"`
	Wording string      `json:"wording"`
}

func DeleteMsg(self_id, message_id interface{}) (bool, error) {
	post := map[string]interface{}{
		"message_id": message_id,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/delete_msg", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var dls DefaultRetStruct
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &dls)
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
