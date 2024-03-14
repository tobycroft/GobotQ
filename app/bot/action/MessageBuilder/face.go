package MessageBuilder

type Face struct {
	Type string `json:"type"`
	Data struct {
		Id  int64 `json:"id"`
		Big bool  `json:"big"`
	} `json:"data"`
}

type BubbleFace struct {
	Type string `json:"type"`
	Data struct {
		Id    int64 `json:"id"`
		Count int64 `json:"count"`
	} `json:"data"`
}
