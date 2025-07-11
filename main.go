package main

import (
	"log"

	"github.com/kryast/Crud-2.git/database"
	"github.com/kryast/Crud-2.git/handlers"
	"github.com/kryast/Crud-2.git/models"
	"github.com/kryast/Crud-2.git/repositories"
	"github.com/kryast/Crud-2.git/router"
	"github.com/kryast/Crud-2.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal koneksi DB:", err)
	}
	db.AutoMigrate(&models.Product{})

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	r := router.SetupRouter(productHandler)
	r.Run(":8080")
}
