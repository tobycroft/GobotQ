package apipost

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

func (api ApiPost) DeleteFriend(self_id, friend_id any) (bool, error) {
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
