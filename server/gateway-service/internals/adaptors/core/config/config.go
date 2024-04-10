package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
)

const DEFAULT_ENVIRONMENT = "development"

type config *models.Config

var env string
var cfg *models.Config

type Adaptor struct {
	environment string
	config      *models.Config
}

func Initialize() *Adaptor {
	cfg = &models.Config{
		Env:      make(map[string]any),
		DataChan: make(chan any),
		Logger:   log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Response: &models.JSONResponse{},
		Error:    &models.ErrorResponse{},
	}
	return &Adaptor{
		environment: "",
		config:      cfg,
	}
}

func (cfgAd *Adaptor) GetConfig() *models.Config {
	return cfgAd.config
}

func (cfgAd *Adaptor) LoadEnvironment() error {
	flag.StringVar(&env, "envflag", DEFAULT_ENVIRONMENT, "sets the default environment stage")
	flag.Parse()
	if env == "production" {
		for _, value := range os.Environ() {
			e := strings.Split(value, "=")
			k, v := e[0], e[1]
			cfgAd.config.Env[k] = v
		}
	} else {
		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "environment", fmt.Sprintf("%s.env", env))
		viper.SetConfigFile(filePath)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config file %v", err)
		}
		for key, value := range viper.AllSettings() {
			cfgAd.config.Env[key] = value
		}
	}
	return nil
}
