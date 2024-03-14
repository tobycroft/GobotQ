package MessageBuilder

type Video struct {
	Type string `json:"type"`
	Data struct {
		File string `json:"file"`
	} `json:"data"`
}
