package MessageBuilder

import "github.com/tobycroft/Calc"

type reply struct {
	Id int64 `json:"id"`
}

func (self *IMessageBuilder) Reply(MessageId int64) *IMessageBuilder {
	self.message = append(self.message, iMessage[reply]{
		Type: "reply",
		Data: reply{
			Id: MessageId,
		},
	})
	self.rawMessage.WriteString("[CQ:reply,id=" + Calc.Any2String(MessageId) + "]")
	return self
}
