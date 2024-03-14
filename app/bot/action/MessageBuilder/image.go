package MessageBuilder

type image struct {
	File string `json:"file"`
}

func (self IMessageBuilder) Image(File string) IMessageBuilder {
	self.New()
	self.message = append(self.message, iMessage[image]{
		Type: "image",
		Data: image{
			File: File,
		},
	})
	self.rawMessage.WriteString("[CQ:image,file=" + File + "]")
	return self
}
