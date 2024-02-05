package meta_event

type EventStruct struct {
	SelfId          int64           `json:"self_id"`
	MessageType     string          `json:"message_type"`
	PostType        string          `json:"post_type"`
	MetaEventStruct MetaEventStruct `json:"json"`
	RemoteAddr      string          `json:"remote_addr"`
}
type MetaEventStruct struct {
	Time          int64  `json:"time"`
	SelfId        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
	Status        struct {
		Self struct {
			Platform string `json:"platform"`
			UserId   int64  `json:"user_id"`
		} `json:"self"`
		Online   bool   `json:"online"`
		Good     bool   `json:"good"`
		QqStatus string `json:"qq.status"`
	} `json:"status"`
	Interval int64 `json:"interval"`
}
