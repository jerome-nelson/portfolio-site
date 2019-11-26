package config

import (
	"net/url"
	"fmt"
	"github.com/joho/godotenv"
	"strconv"
)

/*
Configuration type is used
*/
type Configuration struct {
	Server struct {
		Port int
	}
	DB struct {
		URL  *url.URL
		Name string
	}
}

/*
GetConfig used to grab config
*/
func GetConfig() Configuration {
	err := godotenv.Load()


	var configMap map[string]string
	configMap, err = godotenv.Read()

	if err != nil {
		fmt.Println("Cannot get config: ", err)
		return Configuration{}
	}

	config := Configuration{}
	 if config.Server.Port, err = strconv.Atoi(configMap["PORT"]); err != nil {
		fmt.Println("No config Server Port: ", err)
	 }

	if config.DB.URL, err = url.Parse(configMap["DB"]); err !=nil {
		fmt.Println("No config DB Url: ", err)
	}
	config.DB.Name = configMap["DB_NAME"];

	if err != nil {
		return Configuration{}
	}

	return config
}
