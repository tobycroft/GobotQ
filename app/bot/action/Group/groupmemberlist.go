package Private

type Gml struct {
	Ret  string           `json:"ret"`
	List GroupMemberLists `json:"List"`
}

type GroupMemberLists []struct {
	UIN              int    `json:"UIN"`
	Age              int    `json:"Age"`
	Sex              int    `json:"Sex"`
	NickName         string `json:"NickName"`
	Email            string `json:"Email"`
	Card             string `json:"Card"`
	Remark           string `json:"Remark"`
	SpecTitle        string `json:"SpecTitle"`
	Phone            string `json:"Phone"`
	SpecTitleExpired int    `json:"SpecTitleExpired"`
	MuteTime         int    `json:"MuteTime"`
	AddGroupTime     int    `json:"AddGroupTime"`
	LastMsgTime      int    `json:"LastMsgTime"`
	GroupLevel       int    `json:"GroupLevel"`
}
