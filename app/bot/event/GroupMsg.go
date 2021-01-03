package event

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupMsgModel"
)

type GM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN       int    `json:"UIN"`
		Card      string `json:"Card"`
		SpecTitle string `json:"SpecTitle"`
		Pos       struct {
			Lo int `json:"Lo"`
			La int `json:"La"`
		} `json:"Pos"`
	} `json:"FromQQ"`
	LogonQQ   int `json:"LogonQQ"`
	TimeStamp struct {
		Recv int `json:"Recv"`
		Send int `json:"Send"`
	} `json:"TimeStamp"`
	FromGroup struct {
		GIN  int    `json:"GIN"`
		Name string `json:"name"`
	} `json:"FromGroup"`
	Msg struct {
		Req       int    `json:"Req"`
		Random    int    `json:"Random"`
		SubType   int    `json:"SubType"`
		AppID     int    `json:"AppID"`
		Text      string `json:"Text"`
		TextReply string `json:"Text_Reply"`
		BubbleID  int    `json:"BubbleID"`
	} `json:"Msg"`
	File struct {
		ID   string `json:"ID"`
		MD5  string `json:"MD5"`
		Name string `json:"Name"`
		Size int64  `json:"Size"`
	} `json:"File"`
}

func GroupMsg(gm GM) {
	GroupMsgModel.Api_insert(gm.LogonQQ, gm.FromQQ.UIN, gm.FromGroup.GIN, gm.Msg.Text, gm.Msg.Req, gm.Msg.Random, gm.File.ID, gm.File.MD5,
		gm.File.Name, gm.File.Size)
	is_self := false

	text := gm.Msg.Text
	bot := gm.LogonQQ
	uid := gm.FromQQ.UIN
	gid := gm.FromGroup.GIN
	retract := gm.Msg.Random

	if gm.LogonQQ == gm.FromQQ.UIN {
		is_self = true
	}

	if !is_self {
		GroupHandle(bot, gid, uid, text, gm.Msg.Req, retract)
	}

}

func GroupHandle(bot, gid, uid int, text string, req int, random int) {
	//active, _ := regexp.MatchString("(?i)^acfur", text)

	api.Sendgroupmsg(bot, gid, "Hi我是V!")

}
