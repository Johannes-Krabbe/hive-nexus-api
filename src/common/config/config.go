package config

import (
	"log"
	//	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	DBUrl     string `mapstructure:"DB_URL"`
	ClientUrl string `mapstructure:"CLIENT_URL"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

/*
func LoadConfig() (c Config, err error) {
	log.Println("Loading config")
	if os.Getenv("ENVIRONMENT") == "local" {

		viper.SetConfigFile("./.env")

		viper.AutomaticEnv()

		err = viper.ReadInConfig()

		if err != nil {
			return
		}

	} else {
		viper.AutomaticEnv()
	}

	err = viper.Unmarshal(&c)
	log.Println("Successfully loaded config")
	return
}

func GetValueFromEnv(key string) string {
	value := viper.GetString(key)
	return value
}
*/

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

func GetValueFromEnv(key string) string {
	value := viper.GetString(key)
	return value
}
