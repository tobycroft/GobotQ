package MessageBuilder

type poke struct {
	Type     int64 `json:"type"`
	Id       int64 `json:"id"`
	Strength int64 `json:"strength"`
}

func (self IMessageBuilder) Poke() IMessageBuilder {
	self.New()
	self.message = append(self.message, iMessage[poke]{
		Type: "poke",
		Data: poke{
			Type:     1,
			Id:       10000,
			Strength: 1,
		},
	})
	self.rawMessage.WriteString("[CQ:poke,type=1,id=10000]")
	return self
}
