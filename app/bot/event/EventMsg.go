package event

type EM struct {
	Type   string `json:"Type"`
	FromQQ struct {
		UIN      int64  `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"FromQQ"`
	OperateQQ struct {
		UIN      int64  `json:"UIN"`
		NickName string `json:"NickName"`
	} `json:"OperateQQ"`
	LogonQQ   int64 `json:"LogonQQ"`
	FromGroup struct {
		GIN  int64  `json:"GIN"`
		Name string `json:"Name"`
	} `json:"FromGroup"`
	Msg struct {
		Seq       int    `json:"Seq"`
		TimeStamp int    `json:"TimeStamp"`
		Type      int    `json:"Type"`
		SubType   int    `json:"SubType"`
		Text      string `json:"Text"`
	} `json:"Msg"`
}

func EventMsg(em EM) {

}
