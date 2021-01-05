package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
	"net/url"
	"time"
)

/*
{
    "ret": "{\"retcode\":0,\"retmsg\":\"\",\"time\":\"1609564779\"}"
}
*/
var Group_send_chan = make(chan GroupSendStruct, 20)

type GroupMsg struct {
	Ret string `json:"ret"`
}

type GroupMsgRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Sendgroupmsg(fromqq, togroup interface{}, text string) {
	var gss GroupSendStruct
	gss.Fromqq = fromqq
	gss.Togroup = togroup
	gss.Text = text

	select {
	case Group_send_chan <- gss:

	case <-time.After(5 * time.Second):
		return
	}
}

type GroupSendStruct struct {
	Fromqq  interface{}
	Togroup interface{}
	Text    string
}

func Send_group() {
	for gss := range Group_send_chan {
		sendgroupmsg(gss)
	}
}

func sendgroupmsg(gss GroupSendStruct) (GroupMsg, GroupMsgRet, error) {
	post := map[string]interface{}{
		"fromqq":  gss.Fromqq,
		"togroup": gss.Togroup,
		"text":    url.QueryEscape(gss.Text),
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendgroupmsg", nil, post, nil, nil)
	if err != nil {
		return GroupMsg{}, GroupMsgRet{}, err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return GroupMsg{}, GroupMsgRet{}, err
	}
	var ret GroupMsgRet
	err = jsr.UnmarshalFromString(gm.Ret, &ret)
	if err != nil {
		return gm, GroupMsgRet{}, err
	}
	if ret.Retcode != 0 {
		return gm, ret, errors.New(ret.Retmsg)
	}
	return gm, ret, nil
}
