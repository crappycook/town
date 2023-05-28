package bootstrap

import (
	"bys/pkg/ulog"
	"flag"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
)

func LoadConfig() *Config {
	var configPath string
	flag.StringVar(&configPath, "config", "./bootstrap/config_sample.yml", "server config path")
	flag.Parse()

	cfg := loadConfigFromFile(configPath)
	infoLog := &ulog.TeeRotateOption{
		LogWritter: &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s-info.log", cfg.Log.Dir, time.Now().Format("2006-01-02-15")),
			MaxSize:    1,
			MaxAge:     1,
			MaxBackups: 1,
			LocalTime:  true,
			Compress:   true,
		},
		LevelEnablerFunc: func(l ulog.Level) bool {
			return l <= ulog.InfoLevel
		},
	}
	errorLog := &ulog.TeeRotateOption{
		LogWritter: &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s-error.log", cfg.Log.Dir, time.Now().Format("2006-01-02-15")),
			MaxSize:    1,
			MaxAge:     1,
			MaxBackups: 1,
			LocalTime:  true,
			Compress:   true,
		},
		LevelEnablerFunc: func(l ulog.Level) bool {
			return l > ulog.InfoLevel
		},
	}
	opts := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}
	logger := ulog.NewTeeRotateLogger(
		[]*ulog.TeeRotateOption{infoLog, errorLog},
		opts...,
	)
	ulog.AttachLogger(logger)

	return cfg
}

func loadConfigFromFile(path string) *Config {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := new(Config)
	err = yaml.Unmarshal(content, config)
	if err != nil {
		panic(err)
	}
	return config
}
