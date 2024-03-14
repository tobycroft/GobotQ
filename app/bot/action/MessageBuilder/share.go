package MessageBuilder

type share struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	File    string `json:"file"`
}

func (self IMessageBuilder) Share(Url, Title, Content, Image, File string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[share]{
		Type: "share",
		Data: share{
			Url,
			Title,
			Content,
			Image,
			File,
		},
	})
	self.RawMessage.WriteString("[CQ:share,url=" + Url + ",title=" + Title + ",content=" + Content + ",image=" + Image + ",file=" + File + "]")
	return self
}
