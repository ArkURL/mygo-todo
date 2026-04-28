package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host         string
	User         string
	Password     string
	Name         string
	Port         int
	SSLMode      string
	TimeZone     string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifeTime  time.Duration `mapstructure:"max_life_time"`
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

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.Port,
		c.SSLMode,
		c.TimeZone,
	)
}

func (c ServerConfig) PORT() string {
	return fmt.Sprintf(":%d", c.Port)
}
