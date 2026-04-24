package config

import "github.com/spf13/viper"

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port int
}

var Conf Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}
