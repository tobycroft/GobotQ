package MessageBuilder

import (
	"strings"
)

type IMessageBuilder struct {
	Message    []any
	RawMessage strings.Builder
}

type iMessage[T im] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type im interface {
	at | basketball | caiquan | face | bubbleFace | image | music | poke | pokeDoubleTap | record | reply | text | touzi | url | video
}

func (self IMessageBuilder) New() IMessageBuilder {
	self.Message = make([]any, 0)
	self.RawMessage = strings.Builder{}
	return self
}

func (self IMessageBuilder) BuildRawMessage() string {
	return self.RawMessage.String()
}
