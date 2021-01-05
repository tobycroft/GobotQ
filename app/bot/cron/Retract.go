package cron

import "main.go/app/bot/api"

type Retract_private struct {
	Fromqq interface{}
	Toqq   interface{}
	Random interface{}
	Req    interface{}
	Time   interface{}
}

type Retract_group struct {
	Fromqq interface{}
	Group  interface{}
	Random interface{}
	Req    interface{}
}

var Retract_chan_private = make(chan Retract_private, 20)
var Retract_chan_group = make(chan Retract_group, 20)

func Retract() {
	go retract_private()
	retract_group()
}

func retract_private() {
	for r := range Retract_chan_private {
		api.Deleteprivatemsg(r.Fromqq, r.Toqq, r.Random, r.Req, r.Time)
	}
}

func retract_group() {
	for r := range Retract_chan_group {
		api.Deletegroupmsg(r.Fromqq, r.Group, r.Random, r.Req)
	}
}
