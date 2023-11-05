package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigPort interface {
	LoadConfig() (ConfigPort, error)
	GetServerConfig() ServerConfig
	GetLoggerConfig() LoggerConfig
}

type ConfigAdapter struct {
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	ServerPort string
	ClientPort string
	ServerUrl  string
}

type LoggerConfig struct {
	OutputType string
	Level      string
}

func New() ConfigPort {
	return &ConfigAdapter{}
}

func (config ConfigAdapter) LoadConfig() (ConfigPort, error) {
	log.Println("loading configuration...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("config file not found...")
	}
	c := &ConfigAdapter{}

	// setting server configuration
	if viper.IsSet("SERVER_PORT") {
		c.Server.ServerPort = viper.GetString("SERVER_PORT")
	} else {
		log.Fatal("SERVER_PORT config not set")
	}

	if viper.IsSet("SERVER_URL") {
		c.Server.ServerUrl = viper.GetString("SERVER_URL")
	} else {
		log.Fatal("SERVER_URL config not set")
	}
	// setting server configuration
	if viper.IsSet("CLIENT_PORT") {
		c.Server.ClientPort = viper.GetString("CLIENT_PORT")
	} else {
		log.Fatal("CLIENT_PORT config not set")
	}
	// setting logger configuration
	if viper.IsSet("LOG_OUTPUT") {
		c.Logger.OutputType = viper.GetString("LOG_OUTPUT")
	} else {
		log.Fatal("LOG_OUTPUT config not set")
	}

	if viper.IsSet("LOG_LEVEL") {
		c.Logger.Level = viper.GetString("LOG_LEVEL")
	} else {
		log.Fatal("LOG_LEVEL config not set")
	}

	log.Println("configuration loaded successfully")
	return *c, nil
}

func (config ConfigAdapter) GetServerConfig() ServerConfig {
	return config.Server
}

func (config ConfigAdapter) GetLoggerConfig() LoggerConfig {
	return config.Logger
}
