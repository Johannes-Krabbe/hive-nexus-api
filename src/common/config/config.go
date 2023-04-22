package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	DBUrl     string `mapstructure:"DB_URL"`
	ClientUrl string `mapstructure:"CLIENT_URL"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (c Config, err error) {
	log.Println("Loading config")
	// viper.SetConfigFile("./app.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	log.Println(c)
	log.Println("Successfully loaded config")
	return
}

func GetValueFromEnv(key string) string {
	value := viper.GetString(key)
	return value
}

/*

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

*/
/*
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
*/

/*
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
*/
