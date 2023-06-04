package router

import (
	"github.com/Novometrix/util/middleware"
	"github.com/RaymondSalim/API-server-template/config"
	"github.com/RaymondSalim/API-server-template/server/constants"
	"github.com/RaymondSalim/API-server-template/server/controller"
	"github.com/gin-gonic/gin"
)

/*
	The router ties all dependencies together.
	It creates a database instance and a gin router with custom middlewares for logging, cors, tracing, csrf, caching and oauth.
	Then the repositories, services and controllers are instantiated and are injected as dependencies where they are required.
	After that the router registers the paths with the functions on the controllers.
*/

/*
	fooRepository := repository.NewFooRepsitory(db)
	fooService := service.NewFooService(fooRepository)
	fooController := service.NewFooController(fooService)

	fooRouter.GET("/get-foo-path", fooController.GetFoo)
*/

func Init(engine *gin.Engine, controllers *controller.Controllers, cfg *config.AppConfig) {
	engine.RedirectTrailingSlash = false

	wrapped := engine.Group("")
	wrapped.Use(middleware.ResponseWrapperMiddleware())
	{
		// Foo Endpoints
		foo := wrapped.Group("/foo")
		{
			foo.POST("/create", controllers.FooController.AddFoo)
			foo.POST("/get", controllers.FooController.GetFoo)
			foo.POST("/delete", controllers.FooController.AddFoo)
		}

		// Counter Endpoints
		counter := wrapped.Group("/counter")
		{
			counter.GET("/get", controllers.CounterController.GetLastCounter)
			counter.POST("/add", controllers.CounterController.AddCounter)
			counter.POST("/reset", controllers.CounterController.ResetCounter)
		}
	}

	// Health endpoint
	engine.GET("/healthz", controllers.HealthController.Status)

	if cfg.Environment != constants.EnvironmentProduction {
		InitSwaggerRoutes(engine)
	}
}
