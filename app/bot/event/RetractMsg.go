package event

type Struct_retract struct {
	Self_id   interface{}
	MessageId interface{}
}

var Retract_chan_private = make(chan Struct_retract, 20)
var Retract_chan_group = make(chan Struct_retract, 20)

var Retract_chan_private_instant = make(chan Struct_retract, 20)
var Retract_chan_group_instant = make(chan Struct_retract, 20)
