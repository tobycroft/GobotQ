package notice

type EventStruct[T Notice | GroupApply | GroupAdmin | GroupBan | GroupIncrease | GroupDecrease | GroupLiftBan | groupRecallMessage] struct {
	SelfId      int64  `json:"self_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	Json        T      `json:"json"`
	RemoteAddr  string `json:"remote_addr"`
}
type Notice struct {
	Time       int64  `json:"time"`
	SelfId     int64  `json:"self_id"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupId    int64  `json:"group_id"`
}

type GroupApply struct {
	OperatorId int64  `json:"operator_id"`
	Flag       string `json:"flag"`
}

type GroupAdmin struct {
	TargetId int64 `json:"target_id"`
}

type GroupIncrease struct {
	OperatorId int64 `json:"operator_id"`
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	TargetId   int64 `json:"target_id"`
}

type GroupDecrease struct {
	OperatorId int64 `json:"operator_id"`
	GroupId    int   `json:"group_id"`
	UserId     int   `json:"user_id"`
	TargetId   int   `json:"target_id"`
}

type GroupBan struct {
	OperatorId int64 `json:"operator_id"`
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	Duration   int64 `json:"duration"`
	TargetId   int64 `json:"target_id"`
}

type GroupLiftBan struct {
	OperatorId int64 `json:"operator_id"`
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	TargetId   int64 `json:"target_id"`
}
