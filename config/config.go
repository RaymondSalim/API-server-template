package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

const configType = "toml"

type AppConfig struct {
	Environment string
	Server      struct {
		ServiceName string
		Port        string
		Version     string
	}

	Database struct {
		Type     string
		Host     string
		Port     string
		Database string
		Schema   string
		SSLMode  string
		Username *string
		Password *string
		Timezone string
	}
}

func GetAppConfig() AppConfig {
	var c AppConfig

	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType(configType)
	v.AddConfigPath(".")

	v.SetDefault("GOENV", Development)
	c.Environment = strings.ToLower(v.GetString("GOENV"))

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %w", err))
	}

	err = v.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("fatal error while unmarshaling config file: %w", err))
	}

	return c
}
