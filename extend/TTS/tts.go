package TTS

import (
	"errors"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"regexp"
	"time"
)

func (self *Audio) Huihui(text string) (Audio, error) {
	post := Net.Post{}.New().SetTimeOut(100*time.Second).PostFormData("http://tts.aerofsx.com/request_tts.php", nil, map[string]string{
		"service": "StreamElements",
		"voice":   "Huihui",
		"text":    cleanText(text),
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

func cleanText(text string) string {
	// 定义正则表达式模式，匹配需要保留的字符
	re := regexp.MustCompile(`[^\p{Han}\p{Latin}\p{N}，。？！,\s]`) // 匹配除中英文、数字、逗号、句号、问号、感叹号、空格外的字符
	cleanedText := re.ReplaceAllString(text, "")                // 使用 ReplaceAllString 方法替换匹配的字符为空字符串
	return cleanedText
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
