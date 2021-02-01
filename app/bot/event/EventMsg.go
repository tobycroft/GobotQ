package event

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
)

type EM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN      int    `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"FromQQ"`
	OperateQQ struct {
		UIN      int    `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"OperateQQ"`
	LogonQQ   int `json:"LogonQQ"`
	FromGroup struct {
		GIN  int    `json:"GIN"`
		Name string `json:"Name"`
	} `json:"FromGroup"`
	Msg struct {
		Seq       int    `json:"Seq"`
		TimeStamp int    `json:"TimeStamp"`
		Type      int    `json:"Type"`
		SubType   int    `json:"SubType"`
		Text      string `json:"Text"`
	} `json:"Msg"`
}

func EventMsg(em EM) {
	//operator := em.OperateQQ.UIN
	//text := em.Msg.Text
	bot := em.LogonQQ
	uid := em.FromQQ.UIN
	gid := em.FromGroup.GIN
	Type := em.Msg.Type

	groupfunction := GroupFunctionModel.Api_find(gid)
	if len(groupfunction) < 1 {
		GroupFunctionModel.Api_insert(gid)
		groupfunction = GroupFunctionModel.Api_find(gid)
	}
	switch Type {
	//取消管理
	case 9:
		if uid == bot {
			if GroupMemberModel.Api_update_type(gid, uid, "member") {
				api.Sendgroupmsg(bot, gid, "Acfur-Off，权限已回收，将在2小时内退群", true)
			} else {
				api.Sendgroupmsg(bot, gid, "Acfur-Off，权限已回收，数据故障", true)
			}

		} else {

		}
		break

	//设定管理
	case 10:
		if uid == bot {
			if GroupMemberModel.Api_update_type(gid, uid, "admin") {
				api.Sendgroupmsg(bot, gid, "Acfur-On，已获取权限，可使用acfurhelp查看功能", true)
			} else {
				api.Sendgroupmsg(bot, gid, "Acfur-On，已获取权限，数据故障，请使用acfur刷新人数来更新信息", true)
			}

		} else {

		}
		break

	}

}
