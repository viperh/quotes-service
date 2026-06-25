package routes

import (
	"QuotesService/internal/api/controllers"

	_ "QuotesService/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DefineRoutes(g *gin.Engine, c *controllers.Controller) {
	g.GET("/", c.GetRoot)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
