package MessageBuilder

type Url struct {
	Type string `json:"type"`
	Data struct {
		Url     string `json:"url"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Image   string `json:"image"`
		File    string `json:"file"`
	} `json:"data"`
}
