package transltr

import (
	"fmt"
)

// type TranslationPluginService interface {
// 	Configure()
// 	Connect()
// 	Join()

// }

type TranslationPluginService interface {
	Translation(string) (string, error)
	GetPluginName() string
	GetLanguage(string) (string, error)
}
type Transltr struct {
	plugins map[string]TranslationPluginService
}

func NewTransltr() *Transltr {
	return &Transltr{plugins: make(map[string]TranslationPluginService)}
}
func (s *Transltr) AddPlugin(plugin TranslationPluginService) {
	s.plugins[plugin.GetPluginName()] = plugin
}

func (s *Transltr) Translate(sentence string, pluginName string) (string, error) {
	if _, ok := s.plugins[pluginName]; !ok {
		return "", fmt.Errorf("plugin '%s' not found ", pluginName)
	}
	return s.plugins[pluginName].Translation(sentence)
}

func (s *Transltr) GetLanguage(sentence string, pluginName string) (string, error) {
	if _, ok := s.plugins[pluginName]; !ok {
		return "", fmt.Errorf("plugin '%s' not found ", pluginName)
	}
	return s.plugins[pluginName].GetLanguage(sentence)
}
