package group

import "github.com/tobycroft/gorose-pro"

type EventStruct[T GroupMessageStruct | string] struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        T      `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}

type GroupMessageRedirect[T string | GroupMessageStruct] struct {
	Json          T           `json:"json"`
	GroupFunction gorose.Data `json:"group_function"`
	GroupMember   gorose.Data `json:"group_member"`
}

type GroupMessageStruct struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int64  `json:"message_id"`
	GroupId     int64  `json:"group_id"`
	PeerId      int64  `json:"peer_id"`
	UserId      int64  `json:"user_id"`
	Message     []struct {
		Data map[string]any `json:"data"`
		Type string         `json:"type"`
	} `json:"message"`
	RawMessage string      `json:"raw_message"`
	Font       int64       `json:"font"`
	Sender     GroupSender `json:"sender"`
}

type GroupSender struct {
	UserId   int    `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Role     string `json:"role"`
	Title    string `json:"title"`
	Level    string `json:"level"`
}

type RefreshGroupStruct struct {
	UserId  int64
	SelfId  int64
	GroupId int64
}
