package MessageBuilder

type music struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Audio  string `json:"audio"`
	Title  string `json:"title"`
	Singer string `json:"singer"`
	Image  string `json:"image"`
}

func (self IMessageBuilder) Music(Type string, Url string, Audio string, Title string, Singer string, Image string) IMessageBuilder {
	self.Message = append(self.Message, iMessage[music]{
		Type: Type,
		Data: music{
			Url:    Url,
			Audio:  Audio,
			Title:  Title,
			Singer: Singer,
			Image:  Image,
		},
	})
	self.RawMessage.WriteString("[CQ:music,type=" + Type + ",url=" + Url + ",audio=" + Audio + ",title=" + Title + ",singer=" + Singer + ",image=" + Image + "]")
	return self
}
