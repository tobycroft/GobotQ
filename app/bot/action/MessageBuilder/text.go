package MessageBuilder

type text struct {
	Text string `json:"text"`
}

func (self *IMessageBuilder) Text(Message string) *IMessageBuilder {
	self.message = append(self.message, iMessage[text]{
		Type: "text",
		Data: text{
			Text: Message,
		},
	})
	self.rawMessage.WriteString(Message)
	return self
}
