package config

import (
	"github.com/Met4tron/go-rest-api/pkg/logger"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Uri  string `mapstructure:"uri"`
	Name string `mapstructure:"name"`
}

type Configuration struct {
	Server ServerConfig   `mapstructure:"server"`
	Db     DatabaseConfig `mapstructure:"db"`
}

var Config *Configuration

func LoadVariables() {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cmd")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()

	if err != nil {
		logger.Fatal("Error to load env vars from file ", err)
	}

	err = viper.Unmarshal(&configuration)

	if err != nil {
		logger.Fatal("Error to parse env file ", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
