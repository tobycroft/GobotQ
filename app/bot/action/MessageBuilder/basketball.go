package MessageBuilder

import "github.com/tobycroft/Calc"

type basketball struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) Basketball(Id int64) IMessageBuilder {
	self.Message = append(self.Message, iMessage[basketball]{
		Type: "basketball",
		Data: basketball{Id: Id},
	})
	self.RawMessage.WriteString("[CQ:basketball,id=" + Calc.Any2String(Id) + "]")
	return self
}
