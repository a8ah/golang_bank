package utils

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ENV         string
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func (c Configuration) getEnv() string {
	if c.ENV == "" {
		return "local"
	}
	return c.ENV

}

func GetConfig(params ...string) (Configuration, error) {

	configuration := Configuration{}

	if err := fileExist("./config.json"); err != nil {
		log.Print("No File found. \n", err)
		return configuration, err
	}

	gonfig.GetConf("./config.json", &configuration)

	env := configuration.ENV
	log.Print("ENV:" + Configuration.getEnv(configuration) + " \n")

	if env != "" {
		fileName := fmt.Sprintf("./%s_config.json", env)
		if err := fileExist(fileName); err != nil {
			log.Print("No File found. \n", err)
			return configuration, err
		}
		gonfig.GetConf(fileName, &configuration)
	}

	return configuration, nil
}

func fileExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("No File found :" + path)
	}
	return nil
}

func Port() string {

	var port string

	configuration, err := GetConfig()
	if err != nil || configuration.PORT == "" {
		port = "3000"
	} else {
		port = configuration.PORT
	}

	return port
}
