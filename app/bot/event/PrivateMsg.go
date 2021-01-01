package event

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/FriendListModel"
	"main.go/app/bot/model/PrivateMsgModel"
)

type PM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN      int    `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"FromQQ"`
	LogonQQ   int `json:"LogonQQ"`
	TimeStamp struct {
		Recv int `json:"Recv"`
		Send int `json:"Send"`
	} `json:"TimeStamp"`
	FromGroup struct {
		GIN int `json:"GIN"`
	} `json:"FromGroup"`
	Msg struct {
		Req         int    `json:"Req"`
		Seq         int64  `json:"Seq"`
		Type        int    `json:"Type"`
		SubType     int    `json:"SubType"`
		SubTempType int    `json:"SubTempType"`
		Text        string `json:"Text"`
		BubbleID    int    `json:"BubbleID"`
	} `json:"Msg"`
	Hb struct {
		Type int `json:"Type"`
	} `json:"Hb"`
	File struct {
		ID   string `json:"ID"`
		MD5  string `json:"MD5"`
		Name string `json:"Name"`
		Size int    `json:"Size"`
	} `json:"File"`
}

func PrivateMsg(pm PM) {
	PrivateMsgModel.Api_insert(pm.LogonQQ, pm.FromQQ.UIN, pm.Msg.Text, pm.Msg.Req, pm.Msg.Seq, pm.Msg.Type, pm.Msg.SubType, pm.File.ID,
		pm.File.MD5, pm.File.Name, pm.File.Size)

	bots := BotModel.Api_select()
	for _, bot := range bots {
		fl, err := api.Getfriendlist(bot["bot"])
		if err != nil {

		} else {
			FriendListModel.Api_delete(bot["bot"])
			var fss FriendListModel.FriendLists
			for _, fll := range fl {
				var fs FriendListModel.FriendList
				fs.Bot = bot["bot"]
				fs.Uid = fll.UIN
				fs.Nickname = fll.NickName
				fs.Email = fll.Email
				fs.Remark = fll.Remark
				fss.Fl = append(fss.Fl, fs)
			}
			FriendListModel.Api_insert_more(fss)
		}
	}

}
