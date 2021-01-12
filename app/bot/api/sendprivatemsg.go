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
    "ret": "{\"retcode\":0,\"retmsg\":\"\",\"time\":\"1609562723\"}",
    "random": 1609583468,
    "req": 21315
}
*/

var Private_send_chan = make(chan PrivateSendStruct, 20)

type PrivateMsg struct {
	Ret    string `json:"ret"`
	Random int    `json:"random"`
	Req    int    `json:"req"`
}

type PrivateMsgRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Sendprivatemsg(fromqq, toqq interface{}, text string, AutoRetract bool) {
	var pss PrivateSendStruct
	pss.Fromqq = fromqq
	pss.Toqq = toqq
	pss.Text = text
	pss.AutoRetract = AutoRetract

	select {
	case Private_send_chan <- pss:

	case <-time.After(5 * time.Second):
		return
	}
}

type PrivateSendStruct struct {
	Fromqq      interface{}
	Toqq        interface{}
	Text        string
	AutoRetract bool
}

func Send_private() {
	for pss := range Private_send_chan {
		sendprivatemsg(pss)
	}
}

func sendprivatemsg(pss PrivateSendStruct) (PrivateMsg, PrivateMsgRet, error) {
	post := map[string]interface{}{
		"fromqq": pss.Fromqq,
		"toqq":   pss.Toqq,
		"text":   url.QueryEscape(pss.Text),
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendprivatemsg", nil, post, nil, nil)
	if err != nil {
		return PrivateMsg{}, PrivateMsgRet{}, err
	}
	var pm PrivateMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &pm)
	if err != nil {
		return PrivateMsg{}, PrivateMsgRet{}, err
	}
	var ret PrivateMsgRet
	err = jsr.UnmarshalFromString(pm.Ret, &ret)
	if err != nil {
		return pm, PrivateMsgRet{}, err
	}
	if ret.Retcode != 0 {
		return pm, ret, errors.New(ret.Retmsg)
	}
	return pm, ret, nil
}
