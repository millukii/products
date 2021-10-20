package router

import (
	"api-products/internal/products/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(baseUrl string, router *gin.Engine, productsController handler.ProductsController) *gin.Engine {
	router.GET("/health", productsController.Healthcheck)

	g := router.Group(baseUrl)

	g.Use(gin.Logger())

	g.Use(gin.Recovery())
	g.GET("/", productsController.GetAll)
	g.GET("/:sku", productsController.Get)
	g.DELETE("/:sku", productsController.Delete)
	g.PATCH("/:sku", productsController.Update)
	g.POST("/", productsController.Create)

	return router
}