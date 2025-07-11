package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-2.git/handlers"
)

func SetupRouter(productHandler *handlers.ProductHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/products", productHandler.CreateProduct)

	return r
}
