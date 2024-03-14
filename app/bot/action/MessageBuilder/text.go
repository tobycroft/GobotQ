package MessageBuilder

type text struct {
	Text string `json:"text"`
}

func (self IMessageBuilder) Text(Message string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[text]{
		Type: "text",
		Data: text{
			Text: Message,
		},
	})
	self.RawMessage.WriteString(Message)
	return self
}
