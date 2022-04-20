package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
}

type ServerConfiguration struct {
	Port string
}

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// Setup initialize configuration
var (
	Config *Configuration
)

func Setup() {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	dbConfig := DatabaseConfiguration{
		Dbname:   fmt.Sprintf("%v", viper.Get("DB_DATABASE")),
		Username: fmt.Sprintf("%v", viper.Get("DB_USERNAME")),
		Password: fmt.Sprintf("%v", viper.Get("DB_PASSWORD")),
		Host:     fmt.Sprintf("%v", viper.Get("DB_HOST")),
		Port:     fmt.Sprintf("%v", viper.Get("DB_PORT")),
	}
	serverConfig := ServerConfiguration{
		Port: fmt.Sprintf("%v", viper.Get("SERVER_PORT")),
	}
	configuration = &Configuration{
		Server:   serverConfig,
		Database: dbConfig,
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
