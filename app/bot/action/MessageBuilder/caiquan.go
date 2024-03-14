package MessageBuilder

import "github.com/tobycroft/Calc"

type caiquan struct {
	Id int64 `json:"id"`
}

func (self IMessageBuilder) Caiquan(Id int64) IMessageBuilder {
	self.Message = append(self.Message, iMessage[caiquan]{
		Type: "new_rps",
		Data: caiquan{
			Id: Id,
		},
	})
	self.RawMessage.WriteString("[CQ:new_rps,qq=" + Calc.Any2String(Id) + "]")
	return self
}
