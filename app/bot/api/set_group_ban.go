package api

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"
	"main.go/tuuz"

	"main.go/tuuz/Log"
)

type MuteGroupmeMberRet struct {
	Ret string `json:"ret"`
}

func SetGroupBan(self_id, group_id, user_id interface{}, duration float64) (bool, error) {
	post := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"duration": duration,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(errors.New("bot:"+Calc.Any2String(self_id)), tuuz.FUNCTION_ALL())
		return false, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/set_group_ban", nil, post, nil, nil).RetString()
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
