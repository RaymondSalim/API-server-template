package main

import (
	util "github.com/Novometrix/util/middleware"
	"github.com/Novometrix/web-server-template/config"
	"github.com/Novometrix/web-server-template/server/controller"
	"github.com/Novometrix/web-server-template/server/db"
	"github.com/Novometrix/web-server-template/server/repository"
	"github.com/Novometrix/web-server-template/server/router"
	"github.com/Novometrix/web-server-template/server/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title       Web Server Template
// @version     0.0.1
// @description This is a template for Novometrix's web server

// @contact.name  Raymond Salim
// @contact.url   https://raymonds.dev/#contact
// @contact.email raymond+novometrix@raymonds.dev

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /

func main() {
	cfg := config.GetAppConfig()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	database, err := db.Init(&cfg)
	if err != nil {
		log.Panic("failed to initialize database with error: ", err)
	}

	repositories := repository.InitRepository(database)
	services := service.InitService(repositories)
	controllers := controller.InitController(services)

	ginRouter := gin.New()

	ginRouter.Use(util.ResponseWrapperMiddleware)
	router.Init(ginRouter, controllers, cfg)

	log.Info("Starting Server on port " + cfg.Server.Port)
	ginRouter.Run(":" + cfg.Server.Port)
}
