package MessageBuilder

import (
	"strings"
)

type IMessageBuilder struct {
	message    []any
	rawMessage strings.Builder
}

type iMessage[T im] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type im interface {
	at | basketball | caiquan | face | bubbleFace | image | music | poke | pokeDoubleTap | record | reply | text | touzi | share | video
}

func (self IMessageBuilder) New() {
	if len(self.message) < 1 {
		self.rawMessage = strings.Builder{}
		self.message = make([]any, 0)
	}
}

func (self IMessageBuilder) BuildRawMessage() string {
	return self.rawMessage.String()
}
func (self IMessageBuilder) BuildMessage() []any {
	return self.message
}
