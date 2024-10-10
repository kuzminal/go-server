package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Params Config
	Logger zerolog.Logger
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogFile  string `yaml:"logFile"`
	LogLevel int8   `yaml:"logLevel"`
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Info().Err(err)
		wd = ""
	}

	//default values
	viper.SetDefault("port", "8888")
	viper.SetDefault("logFile", "./log.log")
	viper.SetDefault("logLevel", 1)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(wd + "/configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config: %w", err))
	}

	errors := viper.Unmarshal(&Params)
	if errors != nil {
		log.Fatal().Err(errors)
	}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   Params.LogFile,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	Logger = zerolog.New(lumberjackLogger).With().Timestamp().
		Logger().Level(zerolog.Level(Params.LogLevel))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
