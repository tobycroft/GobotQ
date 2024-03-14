package MessageBuilder

import "github.com/tobycroft/Calc"

type face struct {
	Id  int64 `json:"id"`
	Big bool  `json:"big"`
}

type bubbleFace struct {
	Id    int64 `json:"id"`
	Count int64 `json:"count"`
}

func (self IMessageBuilder) Face(Id int64) IMessageBuilder {
	self.New()
	self.message = append(self.message, iMessage[face]{
		Type: "face",
		Data: face{
			Id: Id,
		},
	})
	self.rawMessage.WriteString("[CQ:face,id=" + Calc.Any2String(Id) + "]")
	return self
}

func (self IMessageBuilder) BubbleFace(Id, Count int64) IMessageBuilder {
	self.New()
	self.message = append(self.message, iMessage[bubbleFace]{
		Type: "bubble_face",
		Data: bubbleFace{
			Id:    Id,
			Count: Count,
		},
	})
	self.rawMessage.WriteString("[CQ:bubble_face,id=" + Calc.Any2String(Id) + ",count=" + Calc.Any2String(Count) + "]")
	return self
}
