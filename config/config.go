package config

import (
	"flag"
	"fmt"
	"github.com/RaymondSalim/API-server-template/server/constants"
	"github.com/spf13/viper"
	"strings"
)

type AppConfig struct {
	ConfigFileName string
	Environment    string
	Server         struct {
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

	NSQ struct {
		NSQDUrl       string `mapstructure:"DAEMON_URL"`
		NSQLookupdURL string `mapstructure:"LOOKUP_DAEMON_URL"`
	}
}

type LaunchOptions struct {
	Config string
}

func GetAppConfig() AppConfig {
	var c AppConfig

	launchOpt := GetLaunchOptions()

	v := viper.New()
	v.SetConfigType(constants.ConfigType)
	v.SetConfigName(launchOpt.Config)
	v.AddConfigPath(".")

	v.SetDefault("GOENV", constants.EnvironmentDevelopment)
	c.Environment = strings.ToLower(v.GetString("GOENV"))
	c.ConfigFileName = launchOpt.Config + "." + constants.ConfigType

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

func GetLaunchOptions() *LaunchOptions {
	const (
		defaultConfig = "server"
		configUsage   = "override default config toml file (server.toml) without extension"
	)
	var configName string

	flag.StringVar(&configName, "configname", defaultConfig, configUsage)
	flag.StringVar(&configName, "c", defaultConfig, configUsage)

	flag.Parse()

	return &LaunchOptions{
		Config: configName,
	}
}
