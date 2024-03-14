package MessageBuilder

import (
	"github.com/tobycroft/Calc"
)

type pokeDoubleTap struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) PokeDoubleTap(qq int64) IMessageBuilder {
	self.Message = append(self.Message, iMessage[pokeDoubleTap]{
		Type: "touch",
		Data: pokeDoubleTap{
			Id: qq,
		},
	})
	self.RawMessage.WriteString("[CQ:touch,id=" + Calc.Any2String(qq) + "]")
	return self
}
