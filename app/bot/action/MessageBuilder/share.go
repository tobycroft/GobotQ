package MessageBuilder

type share struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	File    string `json:"file"`
}

func (self IMessageBuilder) Share(Url, Title, Content, Image, File string) IMessageBuilder {
	self.message = append(self.message, iMessage[share]{
		Type: "share",
		Data: share{
			Url,
			Title,
			Content,
			Image,
			File,
		},
	})
	self.raw_message.WriteString("[CQ:share,url=" + Url + ",title=" + Title + ",content=" + Content + ",image=" + Image + ",file=" + File + "]")
	return self
}
