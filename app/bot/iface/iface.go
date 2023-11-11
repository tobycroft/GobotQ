package iface

type Iface interface {
	DeleteFriend(self_id, friend_id any) (bool, error)
	DeleteMsg(self_id, message_id any) (bool, error)
	Getfriendlist(self_id any) ([]FriendList, error)
}
