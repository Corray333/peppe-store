package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const (
	LOCAL_LEVEL = "local"
	DEBUG_LEVEL = "debug"
	PROD_LEVEL  = "prod"
)

type Config struct {
	Env        string `yaml:"env"`
	HTTPServer `yaml:"http_server"`
}
type HTTPServer struct {
	Timeout     int `yaml:"timeout"`
	IdleTimeout int `yaml:"idle_timeout"`
}

var config Config

func GetEnv() string {
	return config.Env
}
func GetTimeout() int {
	return config.Timeout
}
func GetIdleTimeout() int {
	return config.IdleTimeout
}

func ConfigInit() {
	err := godotenv.Load("../.env")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	f, err := os.Open("../configs/config.yaml")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	fmt.Println(config)
	SetupLogger(config.Env)
}

func SetupLogger(env string) {
	var log *slog.Logger

	switch env {
	case LOCAL_LEVEL:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case DEBUG_LEVEL:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case PROD_LEVEL:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic("Invalid environment")
	}
	slog.SetDefault(log)
}
