package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogLevel int    `yaml:"logLevel"`
}

type Path struct {
	FileName  string
	Extension string
	FilePath  string
}

func LoadConfig(confPath string) Config {
	//default values
	viper.SetDefault("port", "8888")
	viper.SetDefault("logLevel", 0)

	viper.AutomaticEnv()

	cp, err := parseConfigPath(confPath)
	if err == nil {
		viper.SetConfigName(cp.FileName)
		viper.SetConfigType(cp.Extension)
		viper.AddConfigPath(cp.FilePath)
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error reading config: %w", err))
		}
	}
	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Could not load configuration, err: %e", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return conf
}

func parseConfigPath(configPath string) (Path, error) {
	if len(configPath) == 0 {
		return Path{}, errors.New("empty config path")
	}
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return Path{}, errors.New("file does not exist")
	}
	cp := Path{}
	var filePath, fileName string
	filePath, fileName = path.Split(configPath)

	cp.Extension = strings.Replace(path.Ext(configPath), ".", "", -1)
	cp.FileName = strings.Replace(fileName, path.Ext(configPath), "", 1)
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return Path{}, err
	}
	cp.FilePath = filePath
	return cp, nil
}
