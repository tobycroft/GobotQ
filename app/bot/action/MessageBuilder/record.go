package MessageBuilder

type record struct {
	File string `json:"file"`
}

func (self IMessageBuilder) Record(File string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[record]{
		Type: "record",
		Data: record{
			File: File,
		},
	})
	self.RawMessage.WriteString("[CQ:record,file=" + File + "]")
	return self
}
