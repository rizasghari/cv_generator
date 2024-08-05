package configs

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Viper *viper.Viper
}

func GetConfig() *Config {
	once.Do(func() {
		config = newConfig()
	})
	return config
}

func newConfig() *Config {
	v := viper.New()
	v.SetConfigName("configs")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AddConfigPath("$HOME")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("Error while parsing configuration file", err)
	}

	return &Config{
		Viper: v,
	}
}
