package db

import (
	"fmt"
	"github.com/RaymondSalim/API-server-template/config"
	"github.com/RaymondSalim/API-server-template/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

/*
	The db package only initializes the database connection and sets the configuration for the number of open connections and the connection lifetime.
*/

func Init(cfg *config.AppConfig) (db *gorm.DB, err error) {
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	log.Infof("connecting to db with type: %s", cfg.Database.Type)
	if strings.ToLower(cfg.Database.Type) == "postgresql" {
		dsn := constructDataSourceName(cfg)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	} else if strings.ToLower(cfg.Database.Type) == "sqlite" {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
		_ = db.AutoMigrate(&models.Counter{}, &models.Foo{})
	} else {
		log.Panic("wrong value of database.type specified in configuration file")
	}

	return
}

func constructDataSourceName(cfg *config.AppConfig) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("host=%s ", cfg.Database.Host))
	sb.WriteString(fmt.Sprintf("port=%s ", cfg.Database.Port))
	sb.WriteString(fmt.Sprintf("user=%s ", *cfg.Database.Username))
	sb.WriteString(fmt.Sprintf("password=%s ", *cfg.Database.Password))
	sb.WriteString(fmt.Sprintf("dbname=%s ", cfg.Database.Database))
	sb.WriteString(fmt.Sprintf("host=%s ", cfg.Database.Host))
	sb.WriteString(fmt.Sprintf("sslmode=%s ", cfg.Database.SSLMode))
	sb.WriteString(fmt.Sprintf("TimeZone=%s ", cfg.Database.Timezone))

	return sb.String()
}
