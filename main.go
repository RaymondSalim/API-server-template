package main

import (
	"context"
	"errors"
	"fmt"
	util "github.com/Novometrix/util/middleware"
	"github.com/RaymondSalim/API-server-template/config"
	"github.com/RaymondSalim/API-server-template/server/consumers"
	"github.com/RaymondSalim/API-server-template/server/controller"
	"github.com/RaymondSalim/API-server-template/server/db"
	"github.com/RaymondSalim/API-server-template/server/producers"
	"github.com/RaymondSalim/API-server-template/server/repository"
	"github.com/RaymondSalim/API-server-template/server/router"
	"github.com/RaymondSalim/API-server-template/server/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title       Web Server Template
// @version     0.0.1
// @description This is a template API server

// @contact.name  Raymond Salim
// @contact.url   https://raymonds.dev/#contact
// @contact.email raymond@raymonds.dev

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /

func main() {
	cfg := config.GetAppConfig()
	fmt.Printf("%+v", cfg)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	database, err := db.Init(&cfg)
	if err != nil {
		log.Panic("failed to initialize database with error: ", err)
	}

	prd := producers.InitProducers(&cfg)
	repositories := repository.InitRepository(database)
	services := service.InitService(repositories, prd)
	controllers := controller.InitController(services)
	csm := consumers.InitConsumers(&cfg, services)

	ginRouter := gin.New()

	ginRouter.Use(util.ResponseWrapperMiddleware)
	router.Init(ginRouter, controllers, &cfg)

	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: ginRouter,
	}

	// Shutdown Publishers and Consumers
	srv.RegisterOnShutdown(func() {
		prd.Stop()
		for _, c := range csm {
			c.Stop()
			<-c.StopChan
		}
	})

	go func() {
		log.Info("Starting Server on port " + cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Errorf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown")
	}

	log.Print("server exiting")
}
