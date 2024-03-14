package MessageBuilder

type Music struct {
	Type string `json:"type"`
	Data struct {
		Type   string `json:"type"`
		Url    string `json:"url"`
		Audio  string `json:"audio"`
		Title  string `json:"title"`
		Singer string `json:"singer"`
		Image  string `json:"image"`
	} `json:"data"`
}
