package db

import (
	"fmt"
	"github.com/RaymondSalim/API-server-template/config"
	"github.com/RaymondSalim/API-server-template/server/constants"
	"github.com/RaymondSalim/API-server-template/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

/*
	The db package only initializes the database connection and sets the configuration for the number of open connections and the connection lifetime.
*/

func Init(cfg *config.AppConfig) (db *gorm.DB, err error) {
	gormConfig := generateGormConfig(cfg)
	log.Debugf("gormConfig: %+v", gormConfig)

	log.Debugf("connecting to db of type: %s", cfg.Database.Type)
	if strings.ToLower(cfg.Database.Type) == "postgresql" {
		dsn := constructDataSourceName(cfg)

		log.Debugf("data source name: %v", dsn)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	} else if strings.ToLower(cfg.Database.Type) == "sqlite" {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
		_ = db.AutoMigrate(&models.Counter{}, &models.Foo{})
	} else {
		log.Panic("wrong value of database.type specified in configuration file")
	}

	return
}

func generateGormConfig(cfg *config.AppConfig) *gorm.Config {
	logLevel := logger.Error

	if cfg.Environment != constants.EnvironmentProduction {
		logLevel = logger.Info
	}

	if strings.ToLower(cfg.Database.Type) == "postgresql" {
		tablePrefix := ""
		if cfg.Database.Schema != "" {
			tablePrefix = fmt.Sprintf("%v.", cfg.Database.Schema)
		}

		return &gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: tablePrefix,
			},
			Logger: logger.Default.LogMode(logLevel),
		}
	}

	return &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logLevel),
	}
}

func constructDataSourceName(cfg *config.AppConfig) string {
	var sb strings.Builder

	if cfg.Database.Host != "" {
		sb.WriteString(fmt.Sprintf("host=%s ", cfg.Database.Host))
	}
	if cfg.Database.Port != "" {
		sb.WriteString(fmt.Sprintf("port=%s ", cfg.Database.Port))
	}
	if cfg.Database.Username != nil && *cfg.Database.Username != "" {
		sb.WriteString(fmt.Sprintf("user=%s ", *cfg.Database.Username))
	}
	if cfg.Database.Password != nil && *cfg.Database.Password != "" {
		sb.WriteString(fmt.Sprintf("password=%s ", *cfg.Database.Password))
	}
	if cfg.Database.Database != "" {
		sb.WriteString(fmt.Sprintf("dbname=%s ", cfg.Database.Database))
	}
	if cfg.Database.Host != "" {
		sb.WriteString(fmt.Sprintf("host=%s ", cfg.Database.Host))
	}
	if cfg.Database.SSLMode != "" {
		sb.WriteString(fmt.Sprintf("sslmode=%s ", cfg.Database.SSLMode))
	}
	if cfg.Database.Timezone != "" {
		sb.WriteString(fmt.Sprintf("TimeZone=%s ", cfg.Database.Timezone))
	}

	return sb.String()
}
