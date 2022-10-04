package router

import (
	"github.com/Novometrix/web-server-template/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRoutes(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
