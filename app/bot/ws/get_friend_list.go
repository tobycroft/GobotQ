package api

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type FriendListRet struct {
	Data    []FriendList `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
}

type FriendList struct {
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	UserID   int64  `json:"user_id"`
}

func (ws Ws) Getfriendlist(self_id any) ([]FriendList, error) {
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_friend_list", nil, nil, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gfl FriendListRet

	err = sonic.UnmarshalString(data, &gfl)
	if err != nil {
		return nil, err
	}
	return gfl.Data, nil
}
