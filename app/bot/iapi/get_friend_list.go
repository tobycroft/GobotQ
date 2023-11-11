package iapi

type FriendListRet struct {
	Data    []FriendList `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
}

type FriendList struct {
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	UserID   int64  `json:"user_id"`
}
