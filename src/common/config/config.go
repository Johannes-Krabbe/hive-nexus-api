package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    Port  string `mapstructure:"PORT"`
    DBUrl string `mapstructure:"DB_URL"`
    ClientUrl string `mapstructure:"CLIENT_URL"`
}

func LoadConfig() (c Config, err error) {
	log.Println("Loading config")
	viper.SetConfigFile("./.env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()

	if err != nil {
        return
    }

    err = viper.Unmarshal(&c)

	log.Println("Successfully loaded config")
    return
}