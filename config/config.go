package config

import "os"

var Config *Configuration
var OSPREFIX = "ZT"
type Configuration struct {
	BaseUrl string
	Language string
}

func LoadConfig() {
	lang := getEnv("LANGUAGE", "en")
	baseUrl := getEnv("BASEURL", "http://transltr:5000")
	Config = &Configuration{BaseUrl:baseUrl,Language:lang}
}



func getEnv(key string, defaultValue string) string {
	fullKey := OSPREFIX + "_" + key
	val := os.Getenv(OSPREFIX + "_" + key)
	if val == "" {
		if defaultValue != "" {
			return defaultValue
		}
		panic(fullKey + " is not set")
	}
	return val

}