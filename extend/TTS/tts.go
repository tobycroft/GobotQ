package TTS

import (
	Net "github.com/tobycroft/TuuzNet"
	"time"
)

func (self Audio) Huihui(text string) (Audio, error) {
	post := Net.Post{}.SetTimeOut(100*time.Second).PostUrlXEncode("http://tts.aerofsx.com/request_tts.php", nil, map[string]interface{}{
		"service": "StreamElements",
		"voice":   "Huihui",
		"text":    text,
	}, map[string]string{}, map[string]string{})
	audio := Audio{}
	err := post.RetJson(&audio)
	return audio, err
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
