package notice

type EventStruct[T Notice | groupApply | groupAdmin | groupBan | groupIncrease | groupDecrease | groupLiftBan] struct {
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

type groupApply struct {
	OperatorId int64  `json:"operator_id"`
	Flag       string `json:"flag"`
}

type groupAdmin struct {
	TargetId int64 `json:"target_id"`
}

type groupIncrease struct {
	OperatorId int64 `json:"operator_id"`
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	TargetId   int64 `json:"target_id"`
}

type groupDecrease struct {
	OperatorId int64 `json:"operator_id"`
	GroupId    int   `json:"group_id"`
	UserId     int   `json:"user_id"`
	TargetId   int   `json:"target_id"`
}

type groupBan struct {
	OperatorId int64 `json:"operator_id"`
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	Duration   int64 `json:"duration"`
	TargetId   int64 `json:"target_id"`
}

type groupLiftBan struct {
	OperatorId int64 `json:"operator_id"`
	UserId     int64 `json:"user_id"`
	SenderId   int64 `json:"sender_id"`
	TargetId   int64 `json:"target_id"`
}
