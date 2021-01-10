package cron

import (
	"main.go/app/bot/api"
	"main.go/app/bot/event"
	"main.go/config/app_conf"
	"time"
)

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

var Retract_chan_private_instant = make(chan Retract_private, 20)
var Retract_chan_group_instant = make(chan Retract_group, 20)

func Retract() {
	go retract_private()
	go retract_group()
	go retract_private_instant()
	retract_group_instant()
}

func retract_private() {
	for r := range event.Retract_chan_private {
		go func() {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case event.Retract_chan_private_instant <- r:

			case <-time.After(5 * time.Second):
				return
			}
		}()
	}
}

func retract_group() {
	for r := range event.Retract_chan_group {
		go func() {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case event.Retract_chan_group_instant <- r:

			case <-time.After(5 * time.Second):
				return
			}
		}()
	}
}

func retract_private_instant() {
	for r := range event.Retract_chan_private_instant {
		api.Deleteprivatemsg(r.Fromqq, r.Toqq, r.Random, r.Req, r.Time)
	}
}

func retract_group_instant() {
	for r := range event.Retract_chan_group_instant {
		api.Deletegroupmsg(r.Fromqq, r.Group, r.Random, r.Req)
	}
}

func retract_private2() {
	for r := range Retract_chan_private {
		go func() {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case Retract_chan_private_instant <- r:

			case <-time.After(5 * time.Second):
				return
			}
		}()
	}
}

func retract_group2() {
	for r := range Retract_chan_group {
		go func() {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case Retract_chan_group_instant <- r:

			case <-time.After(5 * time.Second):
				return
			}
		}()
	}
}

func retract_private_instant2() {
	for r := range Retract_chan_private_instant {
		api.Deleteprivatemsg(r.Fromqq, r.Toqq, r.Random, r.Req, r.Time)
	}
}

func retract_group_instant2() {
	for r := range Retract_chan_group_instant {
		api.Deletegroupmsg(r.Fromqq, r.Group, r.Random, r.Req)
	}
}
