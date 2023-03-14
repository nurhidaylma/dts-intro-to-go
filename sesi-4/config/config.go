package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var configEnv map[string]string

func init() {
	content, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Println("config.json file not found")
	} else {
		err = json.Unmarshal(content, &configEnv)
		if err != nil {
			log.Println("invalid config.json file")
		}
	}
}

func GetValue(key string) string {
	value := GetEnv(key) // check if os environemnt variable with the key exists
	if value == "" {
		value, ok := configEnv[key] // if value is empty, check config.json file
		if !ok {
			return ""
		}
		return value
	}
	return value
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
