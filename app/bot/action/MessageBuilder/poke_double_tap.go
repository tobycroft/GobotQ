package MessageBuilder

import (
	"github.com/tobycroft/Calc"
)

type pokeDoubleTap struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) PokeDoubleTap(qq int64) IMessageBuilder {
	self.New()
	self.message = append(self.message, iMessage[pokeDoubleTap]{
		Type: "touch",
		Data: pokeDoubleTap{
			Id: qq,
		},
	})
	self.rawMessage.WriteString("[CQ:touch,id=" + Calc.Any2String(qq) + "]")
	return self
}
