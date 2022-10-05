package router

import (
	"github.com/Novometrix/web-server-template/config"
	"github.com/Novometrix/web-server-template/server/controller"
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

func Init(engine *gin.Engine, controllers *controller.Controllers, cfg config.AppConfig) {
	engine.RedirectTrailingSlash = false

	// Foo Endpoints
	foo := engine.Group("/foo")
	{
		foo.POST("/create", controllers.FooController.AddFoo)
		foo.POST("/get", controllers.FooController.GetFoo)
		foo.POST("/delete", controllers.FooController.AddFoo)
	}

	// Health endpoint
	engine.GET("/health", controllers.HealthController.Status)

	if cfg.Environment != config.Production {
		InitSwaggerRoutes(engine)
	}
}
