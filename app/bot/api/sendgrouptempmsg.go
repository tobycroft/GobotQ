package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
	"net/url"
)

/*
{
    "ret": "{\"retcode\":0,\"retmsg\":\"\",\"time\":\"1609680793\"}",
    "random": 1609702875,
    "req": 22503
}
*/

type SendGroupTempMsg struct {
	Ret    string `json:"ret"`
	Random int    `json:"random"`
	Req    int    `json:"req"`
}

type SendGroupTempMsgRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Sendgrouptempmsg(fromqq, togroup, toqq interface{}, text string, AutoRetract bool) {
	var gst GroupSendTempStruct
	gst.Fromqq = fromqq
	gst.Togroup = togroup
	gst.Toqq = toqq
	gst.Text = text
	gst.AutoRetract = AutoRetract

	Group_send_temp_chan <- gst
}

var Group_send_temp_chan = make(chan GroupSendTempStruct, 20)

type GroupSendTempStruct struct {
	Fromqq      interface{}
	Togroup     interface{}
	Toqq        interface{}
	Text        string
	AutoRetract bool
}

func Send_GroupTempMsg() {
	for gst := range Group_send_temp_chan {
		pm, pmr, err := sendgrouptempmsg(gst)
		if err != nil {

		} else {
			if gst.AutoRetract {
				var r Retract_private
				r.Toqq = gst.Toqq
				r.Fromqq = gst.Fromqq
				r.Random = pm.Random
				r.Req = pm.Req
				r.Time = pmr.Time
				Retract_chan_private <- r
			}
		}
	}
}

func sendgrouptempmsg(gst GroupSendTempStruct) (SendGroupTempMsg, SendGroupTempMsgRet, error) {
	post := map[string]interface{}{
		"fromqq":  gst.Fromqq,
		"togroup": gst.Togroup,
		"toqq":    gst.Toqq,
		"text":    url.QueryEscape(gst.Text),
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendgrouptempmsg", nil, post, nil, nil)
	if err != nil {
		return SendGroupTempMsg{}, SendGroupTempMsgRet{}, err
	}
	var ret1 SendGroupTempMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return SendGroupTempMsg{}, SendGroupTempMsgRet{}, err
	}
	var ret2 SendGroupTempMsgRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, SendGroupTempMsgRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
