package MessageBuilder

import "github.com/tobycroft/Calc"

type touzi struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) Touzi(Id int64) IMessageBuilder {
	self.Message = append(self.Message, iMessage[touzi]{
		Type: "new_dice",
		Data: touzi{
			Id,
		},
	})
	self.RawMessage.WriteString("[CQ:new_dice,id=" + Calc.Any2String(Id) + "]")
	return self
}
