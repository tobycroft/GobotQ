package MessageBuilder

import (
	"github.com/tobycroft/Calc"
)

type at struct {
	Qq string `json:"qq"`
}

func (self *IMessageBuilder) At(qq any) *IMessageBuilder {
	self.message = append(self.message, iMessage[at]{
		Type: "at",
		Data: at{
			Qq: Calc.Any2String(qq),
		},
	})
	self.rawMessage.WriteString("[CQ:at,qq=" + Calc.Any2String(qq) + "]")
	self.Text(" ")
	return self
}
