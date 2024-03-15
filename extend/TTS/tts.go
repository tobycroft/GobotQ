package TTS

import (
	"errors"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"time"
)

func (self *Audio) Huihui(text string) (Audio, error) {
	post := Net.Post{}.New().SetTimeOut(100*time.Second).PostFormData("http://tts.aerofsx.com/request_tts.php", nil, map[string]string{
		"service": "StreamElements",
		"voice":   "Huihui",
		"text":    text,
	}, map[string]string{}, map[string]string{})
	audio := Audio{}
	err := post.RetJson(&audio)
	if err != nil {
		return audio, err
	}
	if audio.Success == false {
		return audio, errors.New(Calc.Any2String(audio.ErrorMsg))
	}
	return audio, err
}

func (self Audio) New() *Audio {
	return &self
}

type Audio struct {
	Success         bool        `json:"success"`
	AudioUrl        string      `json:"audio_url"`
	Info            string      `json:"info"`
	ErrorMsg        interface{} `json:"error_msg"`
	ServiceResponse interface{} `json:"service_response"`
	Meta            struct {
		Service       string `json:"service"`
		VoiceId       string `json:"voice_id"`
		VoiceName     string `json:"voice_name"`
		Text          string `json:"text"`
		PlaylistIndex int    `json:"playlist_index"`
	} `json:"meta"`
}
