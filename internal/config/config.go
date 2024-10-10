package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Params Config

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		wd = ""
	}

	//default values
	viper.SetDefault("port", "8888")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(wd + "/configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config: %w", err))
	}

	errors := viper.Unmarshal(&Params)
	if errors != nil {
		log.Panicln(errors)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
