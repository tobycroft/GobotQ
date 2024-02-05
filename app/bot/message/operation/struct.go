package operation

import "net"

type EventStruct[T OperationEvent] struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        T      `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}
type OperationEvent struct {
	remoteaddr net.Addr
	json       string
	Echo       struct {
		Action string `json:"action"`
		SelfId int64  `json:"self_id"`
		Extra  any    `json:"extra"`
	} `json:"echo"`
}
