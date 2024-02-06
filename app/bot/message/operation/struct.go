package operation

import "main.go/app/bot/iapi"

type EventStruct[T OperationEvent | iapi.LoginInfoRet | iapi.FriendListRet | iapi.GroupListRet | iapi.GroupMemberListRet | iapi.MessageRet | iapi.RetractWsRetStruct] struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        T      `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}

type OperationEvent struct {
	Echo struct {
		Action string `json:"action"`
		SelfId int64  `json:"self_id"`
		Extra  any    `json:"extra"`
	} `json:"echo"`
}
