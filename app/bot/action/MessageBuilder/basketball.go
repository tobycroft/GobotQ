package MessageBuilder

type Basketball struct {
	Type string `json:"type"`
	Data struct {
		Id int64 `json:"id"`
	} `json:"data"`
}