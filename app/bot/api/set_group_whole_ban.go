package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

type SetGroupWholeMuteRet struct {
	Ret string `json:"ret"`
}

func SetGroupWholeBan(self_id, group_id interface{}, enable bool) (bool, error) {
	post := map[string]interface{}{
		"self_id":  self_id,
		"group_id": group_id,
		"enable":   enable,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/set_group_whole_ban", nil, post, nil, nil)
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
		Log.Crrs(errors.New(dls.Wording), tuuz.FUNCTION_ALL())
		return false, errors.New(dls.Wording)
	}
}
