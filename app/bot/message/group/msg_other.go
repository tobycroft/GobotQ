package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func group_message_other() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivateValid) {
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			pm := es.Json
		}
	}
}
