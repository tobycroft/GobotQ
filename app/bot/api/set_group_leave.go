package api

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

func SetGroupLeave(self_id, group_id interface{}) (bool, error) {
	post := map[string]interface{}{
		"group_id":   group_id,
		"is_dismiss": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post(botinfo["url"].(string)+"/set_group_leave", nil, post, nil, nil)
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
		fmt.Println(dls)
		Log.Crrs(errors.New(dls.Wording), tuuz.FUNCTION_ALL())
		return false, errors.New(dls.Wording)
	}
}
