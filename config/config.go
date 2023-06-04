package config

import (
	"flag"
	"fmt"
	"github.com/RaymondSalim/API-server-template/server/constants"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

// mapstructure field tag is not required, as viper can automatically choose a suitable name, but BindEnvs() only works on fields that has mapstructure declared

type Server struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	Port        string `mapstructure:"PORT"`
	Version     string `mapstructure:"VERSION"`
}

type Database struct {
	Type     string  `mapstructure:"TYPE"`
	Host     string  `mapstructure:"HOST"`
	Port     string  `mapstructure:"PORT"`
	Database string  `mapstructure:"DATABASE"`
	Schema   string  `mapstructure:"SCHEMA"`
	SSLMode  string  `mapstructure:"SSLMODE"`
	Username *string `mapstructure:"USERNAME"`
	Password *string `mapstructure:"PASSWORD"`
	Timezone string  `mapstructure:"TIMEZONE"`
}

type NSQ struct {
	NSQDUrl       string `mapstructure:"DAEMON_URL"`
	NSQLookupdURL string `mapstructure:"LOOKUP_DAEMON_URL"`
}

type AppConfig struct {
	ConfigFileName string
	Environment    string

	Server   `mapstructure:"SERVER"`
	Database `mapstructure:"DATABASE"`
	NSQ      `mapstructure:"NSQ"`
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
	v.AddConfigPath("./config")

	v.AutomaticEnv()
	BindEnvs(v, c)

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

func BindEnvs(viperInstance *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(viperInstance, v.Interface(), append(parts, tv)...)
		default:
			viperInstance.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
