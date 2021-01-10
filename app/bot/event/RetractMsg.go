package event

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
