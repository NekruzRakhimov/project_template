package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"project_template/models"
)

var AppSettings models.Settings

func PutAdditionalSettings() {
	AppSettings.AppParams.LogDebug = "./logs/debug.log"
	AppSettings.AppParams.LogInfo = "./logs/info.log"
	AppSettings.AppParams.LogWarning = "./logs/warning.log"
	AppSettings.AppParams.LogError = "./logs/error.log"

	AppSettings.AppParams.LogCompress = true
	AppSettings.AppParams.LogMaxSize = 10
	AppSettings.AppParams.LogMaxAge = 100
	AppSettings.AppParams.LogMaxBackups = 100
	AppSettings.AppParams.AppVersion = "1.0"
}

func ReadSettings() {
	fmt.Println("Starting reading settings file")
	configFile, err := os.Open("./settings-prod.json")
	if err != nil {
		log.Fatal("Couldn't open config file. Error is: ", err.Error())
	}

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}
	}(configFile)

	fmt.Println("Starting decoding settings file")
	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		log.Fatal("Couldn't decode settings json file. Error is: ", err.Error())
	}

	log.Println(AppSettings)
	return
}
