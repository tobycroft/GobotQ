package MessageBuilder

import (
	"strings"
)

type IMessageBuilder struct {
	message     []any
	raw_message strings.Builder
}

type iMessage[T im] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type im interface {
	at | basketball | caiquan | face | bubbleFace | image | music | poke | pokeDoubleTap | record | reply | text | touzi | share | video
}

func (self IMessageBuilder) New() IMessageBuilder {
	self.message = make([]any, 0)
	self.raw_message = strings.Builder{}
	return self
}

func (self IMessageBuilder) BuildRawMessage() string {
	return self.raw_message.String()
}
func (self IMessageBuilder) BuildMessage() []any {
	return self.message
}
