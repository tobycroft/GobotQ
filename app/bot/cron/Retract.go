package cron

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/event"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Log"
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
	go retract_group_instant()

	go retract_private2()
	go retract_group2()
	go retract_private_instant2()
	go retract_group_instant2()

	go retract_private3()
	go retract_group3()
	go retract_private_instant3()
	retract_group_instant3()
}

func retract_private() {
	for r := range event.Retract_chan_private {
		go func(retract event.Retract_private) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case event.Retract_chan_private_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_group() {
	for r := range event.Retract_chan_group {
		go func(retract event.Retract_group) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case event.Retract_chan_group_instant <- retract:

			case <-time.After(5 * time.Second):
				Log.Errs(errors.New("retract_group失败"), tuuz.FUNCTION_ALL())
				return
			}
		}(r)
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
		go func(retract Retract_private) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case Retract_chan_private_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_group2() {
	for r := range Retract_chan_group {
		go func(retract Retract_group) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case Retract_chan_group_instant <- retract:

			case <-time.After(5 * time.Second):
				Log.Errs(errors.New("retract_group2失败"), tuuz.FUNCTION_ALL())
				return
			}
		}(r)
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

func retract_private3() {
	for r := range api.Retract_chan_private {
		go func(retract api.Retract_private) {
			time.Sleep(app_conf.Retract_time_second * time.Second)
			select {
			case api.Retract_chan_private_instant <- retract:

			case <-time.After(5 * time.Second):
				return
			}
		}(r)
	}
}

func retract_group3() {
	for r := range api.Retract_chan_group {
		go func(retract api.Retract_group) {
			//fmt.Println("retract_group3:countdown", retract)
			time.Sleep(app_conf.Retract_time_second * time.Second)
			//fmt.Println("retract_group3:sendto_chan", retract)

			select {
			case api.Retract_chan_group_instant <- retract:

			case <-time.After(5 * time.Second):
				Log.Errs(errors.New("retract_group3失败"), tuuz.FUNCTION_ALL())
				return
			}
		}(r)
	}
}

func retract_private_instant3() {
	for r := range api.Retract_chan_private_instant {
		api.Deleteprivatemsg(r.Fromqq, r.Toqq, r.Random, r.Req, r.Time)
	}
}

func retract_group_instant3() {
	for r := range api.Retract_chan_group_instant {
		//fmt.Println("retract_group_instant3:", r)
		api.Deletegroupmsg(r.Fromqq, r.Group, r.Random, r.Req)
	}
}
