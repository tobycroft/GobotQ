package MessageBuilder

import "github.com/tobycroft/Calc"

type reply struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) Reply(MessageId int64) IMessageBuilder {
	self.Message = append(self.Message, iMessage[reply]{
		Type: "reply",
		Data: reply{
			Id: MessageId,
		},
	})
	self.RawMessage.WriteString("[CQ:reply,id=" + Calc.Any2String(MessageId) + "]")
	return self
}
