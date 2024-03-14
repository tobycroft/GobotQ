package MessageBuilder

import "github.com/tobycroft/Calc"

type at struct {
	Qq string `json:"qq"`
}

func (self IMessageBuilder) At(qq string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[at]{
		Type: "at",
		Data: at{
			Qq: qq,
		},
	})
	self.RawMessage.WriteString("[CQ:at,qq=" + Calc.Any2String(qq) + "]")
	return self
}
