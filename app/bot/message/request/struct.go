package request

type EventStruct[T RequestMessage | requestFriend | requestGroup] struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        T      `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}

type RequestMessage struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	RequestType string `json:"request_type"`
}

type requestFriend struct {
	UserId  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
}

type requestGroup struct {
	SubType string `json:"sub_type"`
	GroupId int    `json:"group_id"`
	UserId  int    `json:"user_id"`
	Flag    string `json:"flag"`
}
