package MessageBuilder

import "github.com/tobycroft/Calc"

type basketball struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) Basketball(Id int64) IMessageBuilder {
	self.message = append(self.message, iMessage[basketball]{
		Type: "basketball",
		Data: basketball{Id: Id},
	})
	self.raw_message.WriteString("[CQ:basketball,id=" + Calc.Any2String(Id) + "]")
	return self
}
