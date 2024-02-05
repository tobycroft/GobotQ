package index

type EventStruct struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        string `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}
