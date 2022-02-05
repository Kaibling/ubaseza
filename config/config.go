package config

import (
	"os"
	"strings"
)

var Config *Configuration
var OSPREFIX = "ZT"

type IrcConfig struct {
	Plugin      string
	Channels    []string
	Lang_plugin string
}
type LangConfig struct {
	url string
}

type Configuration struct {
	BaseUrl  string
	Language string
	IRC      *IrcConfig
	Lang     *LangConfig
}

func LoadConfig() {
	lang := getEnv("LANGUAGE", "en")
	baseUrl := getEnv("BASEURL", "http://transltr:5000")
	ircPlugin := getEnv("IRC_PLUGIN", "")
	ircChannelString := getEnv("IRC_CHANNELS", "")
	ircChannels := strings.Split(ircChannelString, ",")
	ircLangPlugin := getEnv("IRC_LANG_PLUGIN", "")
	ircConfig := &IrcConfig{Plugin: ircPlugin, Channels: ircChannels, Lang_plugin: ircLangPlugin}
	//channelString := getEnv("LANG_PLUGIN", "")

	Config = &Configuration{BaseUrl: baseUrl, Language: lang, IRC: ircConfig}
}

func getEnv(key string, defaultValue string) string {
	fullKey := OSPREFIX + "_" + key
	val := os.Getenv(OSPREFIX + "_" + key)
	if val == "" {
		if defaultValue != "" {
			return defaultValue
		}
		panic("Env: " + fullKey + " is not set")
	}
	return val

}
