package iapi

import (
	"sync"
	"time"
)

// Api : 是IfaceApi的接口实例模式，使用Ws来启用websocket发送模式，接口化操作
var Api = IfaceApi(Ws{})

var ClientToConn = new(sync.Map)
var ConnToClient = new(sync.Map)

type Post struct{}
type Ws struct{}

type sendStruct struct {
	Action string         `json:"action"`
	Params map[string]any `json:"params"`
	Echo   echo           `json:"echo"`
}

type echo struct {
	Action string `json:"action"`
	SelfId int64  `json:"self_id"`
	Extra  any    `json:"extra"`
}

type IfaceApi interface {
	DeleteFriend(self_id int64, friend_id any) (bool, error)
	DeleteMsg(self_id, message_id int64) (bool, error)
	GetFriendList(self_id int64) ([]FriendList, error)
	GetGroupInfo(self_id, group_id int64) (GroupInfo, error)
	GetGroupList(self_id int64) ([]GroupList, error)
	GetGroupMemberInfo(self_id, group_id, user_id int64) (GroupMemberList, error)
	GetGroupMemberList(self_id, group_id int64) ([]GroupMemberList, error)
	GetLoginInfo(self_id int64) (LoginInfo, error)
	GetStrangerInfo(self_id, user_id int64, no_cache bool) (UserInfo, error)
	SendGroupMsg(Self_id, Group_id int64, Message string, AutoRetract bool)
	SendGroupMsgWithTime(Self_id, Group_id int64, Message string, AutoRetract bool, Time time.Duration)
	Send_group()
	SendPrivateMsg(Self_id, UserId, GroupId int64, Message string, AutoRetract bool)
	SendPrivateMsgWithTime(Self_id, UserId, GroupId int64, Message string, AutoRetract bool, Time time.Duration)
	Send_private()
	SetFriendAddRequest(self_id int64, flag any, approve bool, remark any) (bool, error)
	SetGroupAddRequestRet(self_id int64, flag, sub_type any, approve bool, reason string) (bool, error)
	SetGroupAdmin(self_id, group_id, user_id int64, enable bool) (bool, error)
	SetGroupBan(self_id, group_id, user_id int64, duration float64) (bool, error)
	SetGroupCard(self_id, group_id, user_id int64, card any) (bool, error)
	SetGroupKick(self_id, group_id, user_id int64, reject_add_request bool) (bool, error)
	SetGroupLeave(self_id, group_id int64) (bool, error)
	SetGroupWholeBan(self_id, group_id int64, enable bool) (bool, error)
}
