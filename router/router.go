package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-2.git/handlers"
)

func SetupRouter(productHandler *handlers.ProductHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetProduct)

	return r
}
