package MessageBuilder

type record struct {
	File string `json:"file"`
}

func (self IMessageBuilder) Record(File string) IMessageBuilder {
	self.message = append(self.message, iMessage[record]{
		Type: "record",
		Data: record{
			File: File,
		},
	})
	self.raw_message.WriteString("[CQ:record,file=" + File + "]")
	return self
}
