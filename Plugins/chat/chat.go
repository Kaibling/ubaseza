package chat

import (
	"encoding/json"
	"time"
)

type ChatPlugin interface {
	AddTranslator()
}

type ChatPluginService struct {
}

type ChatMessage struct {
	User       string
	Message    string
	Language   string
	Translated string
	Channel    string
	Time       time.Time
}

func (c ChatMessage) ToJson() string {
	jByteString, _ := json.Marshal(c)
	return string(jByteString)
}
