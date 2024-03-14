package MessageBuilder

type video struct {
	File string `json:"file"`
}

func (self IMessageBuilder) Video(File string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[video]{
		Type: "video",
		Data: video{
			File: File,
		},
	})
	self.RawMessage.WriteString("[CQ:video,file=" + File + "]")
	return self
}
