package private

type EventStruct[T PrivateMessageStruct] struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        T      `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}
type PrivateMessageStruct struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int64  `json:"message_id"`
	TargetId    int64  `json:"target_id"`
	PeerId      int64  `json:"peer_id"`
	UserId      int64  `json:"user_id"`
	Message     []struct {
		Data struct {
			Text string `json:"text"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"message"`
	RawMessage string        `json:"raw_message"`
	Font       int64         `json:"font"`
	Sender     PrivateSender `json:"sender"`
}
type PrivateSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Role     string `json:"role"`
	Title    string `json:"title"`
	Level    string `json:"level"`
}
