package STT

import (
	"errors"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
)

func (self *Audio) SpeechToText(file_url string) (string, error) {
	post := Net.Net{}.New().SetUrl("http://10.0.0.182:84/v1/tts/stt/qq").
		SetPostData(map[string]string{
			"url": file_url,
		}).
		SetHeader(map[string]string{
			"token": Calc.Any2String(SystemParamModel.Api_value("aigc")),
		}).
		PostFormData()
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

func (self *Audio) SpeechBase64ToText(b64 string) (string, error) {
	post := Net.Net{}.New().SetUrl("http://10.0.0.182:84/v1/tts/stt/b64").
		SetPostData(map[string]string{
			"base64": b64,
		}).
		SetHeader(map[string]string{
			"token": Calc.Any2String(SystemParamModel.Api_value("aigc")),
		}).
		PostFormData()
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
