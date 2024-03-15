package STT

import (
	"errors"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
)

func (self *Audio) SpeechToText(file_url string) (string, error) {
	post := Net.Post{}.New().PostUrlXEncode("http://127.0.0.1:84/v1/tts/stt/qq", map[string]interface{}{
		"token": SystemParamModel.Api_value("aigc"),
	}, map[string]interface{}{
		"url": file_url,
	}, map[string]string{}, map[string]string{})
	audio := Audio{}
	err := post.RetJson(&audio)
	if err != nil {
		return "", err
	}
	if audio.Code != 0 {
		return "", errors.New(Calc.Any2String(audio.Echo))
	}
	return audio.Echo, err
}

func (self Audio) New() *Audio {
	return &self
}

type Audio struct {
	Code int64  `json:"code"`
	Echo string `json:"echo"`
}
