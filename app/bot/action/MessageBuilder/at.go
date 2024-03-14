package MessageBuilder

type At struct {
	Type string `json:"type"`
	Data struct {
		Qq string `json:"qq"`
	} `json:"data"`
}
