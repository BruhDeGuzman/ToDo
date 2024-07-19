package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        string `mapstructure:"PORT"`
	Host        string `mapstructure:"HOST"`
	Url         string `mapstructure:"URL"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

var EnvConfig *Config

func InitEnvConfig() {
	EnvConfig = LoadEnvVariables()
}

func LoadEnvVariables() (conf *Config) {
	viper.AddConfigPath("./env")
	viper.SetConfigName("dev.env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading env file ", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
