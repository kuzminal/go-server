package config

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ConfigPath struct {
	FileName  string
	Extention string
	FilePath  string
}

func LoadConfig(confPath string) Config {
	//default values
	viper.SetDefault("port", "8888")

	viper.AutomaticEnv()

	cp, err := parseConigPath(confPath)
	if err == nil {
		viper.SetConfigName(cp.FileName)
		viper.SetConfigType(cp.Extention)
		viper.AddConfigPath(cp.FilePath)
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error reading config: %w", err))
		}
	}
	var conf Config
	errors := viper.Unmarshal(&conf)
	if errors != nil {
		log.Fatal(errors)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return conf
}

func parseConigPath(configPath string) (ConfigPath, error) {
	if len(configPath) == 0 {
		return ConfigPath{}, errors.New("empty config path")
	}
	cp := ConfigPath{}
	var filePath, fileName string
	filePath, fileName = path.Split(configPath)

	cp.Extention = strings.Replace(path.Ext(configPath), ".", "", -1)
	cp.FileName = strings.Replace(fileName, path.Ext(configPath), "", 1)
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return ConfigPath{}, err
	}
	cp.FilePath = filePath
	return cp, nil
}
