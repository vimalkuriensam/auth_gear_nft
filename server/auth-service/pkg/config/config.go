package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const DEFAULT_ENVIRONMENT = "development"

var env string
var cfg *Config

type Config struct {
	Env      map[string]any
	DataChan chan any
	Logger   *log.Logger
	Response *JSONResponse
	Error    *ErrorResponse
}

type JSONResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Status    int       `json:"status"`
	Path      string    `json:"path"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type ReadValue struct {
	B []byte
	D interface{}
}

func Initialize() *Config {
	cfg = &Config{
		Env:      make(map[string]any),
		DataChan: make(chan any),
		Logger:   log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Response: &JSONResponse{},
		Error:    &ErrorResponse{},
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}

func (config *Config) LoadEnvironment() error {
	flag.StringVar(&env, "envflag", DEFAULT_ENVIRONMENT, "sets the default environment stage")
	flag.Parse()
	if env == "production" {
		for _, value := range os.Environ() {
			e := strings.Split(value, "=")
			k, v := e[0], e[1]
			config.Env[k] = v
		}
	} else {
		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "environment", fmt.Sprintf("%s.env", env))
		viper.SetConfigFile(filePath)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config file %v", err)
		}
		for key, value := range viper.AllSettings() {
			config.Env[key] = value
		}
	}
	return nil
}
