package MessageBuilder

import "github.com/tobycroft/Calc"

type touzi struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) Touzi(Id int64) IMessageBuilder {
	self.message = append(self.message, iMessage[touzi]{
		Type: "new_dice",
		Data: touzi{
			Id,
		},
	})
	self.raw_message.WriteString("[CQ:new_dice,id=" + Calc.Any2String(Id) + "]")
	return self
}
