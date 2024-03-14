package MessageBuilder

type Poke struct {
	Type string `json:"type"`
	Data struct {
		Type     int64 `json:"type"`
		Id       int64 `json:"id"`
		Strength int64 `json:"strength"`
	} `json:"data"`
}
