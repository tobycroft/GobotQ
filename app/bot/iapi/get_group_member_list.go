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

type GroupMemberListRet struct {
	Data    []GroupMemberList `json:"data"`
	Retcode int64             `json:"retcode"`
	Status  string            `json:"status"`
}

type GroupMemberList struct {
	UserId          int64  `json:"user_id"`
	GroupId         int64  `json:"group_id"`
	UserName        string `json:"user_name"`
	Sex             string `json:"sex"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	Nickname        string `json:"nickname"`
	UserDisplayname string `json:"user_displayname"`
	Distance        int64  `json:"distance"`
	Honor           []int  `json:"honor"`
	JoinTime        int64  `json:"join_time"`
	LastActiveTime  int64  `json:"last_active_time"`
	LastSentTime    int64  `json:"last_sent_time"`
	UniqueName      string `json:"unique_name"`
	Area            string `json:"area"`
	Level           int64  `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	CardChangeable  bool   `json:"card_changeable"`
}

func (api Post) Getgroupmemberlist(self_id, group_id any) ([]GroupMemberList, error) {
	post := map[string]any{
		"group_id": group_id,
		"no_cache": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_group_member_list", nil, post, nil, nil).RetString()
	if err != nil {
		return nil, err
	}
	var gms GroupMemberListRet

	err = sonic.UnmarshalString(data, &gms)
	if err != nil {
		return nil, err
	}
	return gms.Data, nil
}
func (api Ws) Getgroupmemberlist(self_id, group_id any) ([]GroupMemberList, error) {
	post := map[string]any{
		"group_id": group_id,
		"no_cache": true,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return nil, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_group_member_list",
		Params: post,
		Echo: echo{
			Action: "get_group_member_list",
			SelfId: Calc.Any2Int64(self_id),
			Extra:  group_id,
		},
	})
	if err != nil {
		return nil, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return nil, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return nil, err
}
