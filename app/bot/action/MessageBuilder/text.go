package MessageBuilder

type Text struct {
	Type string `json:"type"`
	Data struct {
		Text int64 `json:"text"`
	} `json:"data"`
}
