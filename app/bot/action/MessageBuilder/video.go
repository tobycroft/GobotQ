package MessageBuilder

type video struct {
	File string `json:"file"`
}

func (self IMessageBuilder) Video(File string) IMessageBuilder {
	self.New()
	self.message = append(self.message, iMessage[video]{
		Type: "video",
		Data: video{
			File: File,
		},
	})
	self.rawMessage.WriteString("[CQ:video,file=" + File + "]")
	return self
}
