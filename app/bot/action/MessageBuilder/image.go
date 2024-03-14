package MessageBuilder

type image struct {
	File string `json:"file"`
}

func (self IMessageBuilder) Image(File string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[image]{
		Type: "image",
		Data: image{
			File: File,
		},
	})
	self.RawMessage.WriteString("[CQ:image,file=" + File + "]")
	return self
}
