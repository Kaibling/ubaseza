package translation

import (
	"encoding/json"
	"fmt"

	"github.com/kaibling/zucktrans/config"
	"github.com/kaibling/zucktrans/utility"
)

var pluginName = "LibreTranslate"

type LibreTranslateService struct {
	defaultLang string
	baseUrl     string
	pluginName  string
}

func NewLibreTranslateService() *LibreTranslateService {
	defaultLang := config.Config.Language
	baseUrl := config.Config.BaseUrl
	return &LibreTranslateService{defaultLang: defaultLang, baseUrl: baseUrl, pluginName: pluginName}
}

func (s *LibreTranslateService) Translation(sentence string) (string, error) {

	langCodec, err := getLanguage(sentence, s.baseUrl)
	if err != nil {
		return "", err
	}
	translation, err := translate(sentence, langCodec, s.defaultLang, s.baseUrl)
	if err != nil {
		return "", err
	}
	return translation, nil
}
func (s *LibreTranslateService) GetPluginName() string {
	return s.pluginName
}
func (s *LibreTranslateService) GetLanguage(sentence string) (string, error) {
	return getLanguage(sentence, s.baseUrl)
}

func translate(sentence string, sourceLang string, targetLang string, baseUrl string) (string, error) {
	url := baseUrl + "/translate"
	requestData := map[string]string{
		"q":      sentence,
		"source": sourceLang,
		"target": targetLang,
		"format": "text",
	}
	bRequestData, _ := json.Marshal(requestData)
	response := utility.Request(url, []byte(bRequestData))
	var dat map[string]interface{}
	err := json.Unmarshal(response, &dat)
	if err != nil {
		return "", fmt.Errorf("request data could not be parsed %s", string(response))
	}
	lang := dat["translatedText"].(string)
	return lang, nil
}

func getLanguage(sentence string, baseUrl string) (string, error) {
	url := baseUrl + "/detect"
	requestData := map[string]string{
		"q": sentence,
	}
	bRequestData, _ := json.Marshal(requestData)

	response := utility.Request(url, []byte(bRequestData))

	var dat []map[string]interface{}
	err := json.Unmarshal(response, &dat)
	if err != nil {
		return "", fmt.Errorf("request data could not be parsed %s", string(response))
	}
	if len(dat) == 0 {
		return "", fmt.Errorf("request has no length %s", dat)
	}
	lang := dat[0]["language"].(string)
	return lang, nil
}
