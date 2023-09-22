package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var AppConfig Config

func init() {
	_, config, _, _ := runtime.Caller(0)
	viper.AddConfigPath(path.Dir(config))
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("unmarshal error: %w \n", err))
	}

}

type Config struct {
	Webserver struct {
		Port string `json:"port"`
		Mode string `json:"mode"`
	} `json:"webserver"`
	LogLevel string `json:"log_level"`
}
