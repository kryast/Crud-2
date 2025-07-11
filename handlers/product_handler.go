package handlers

import "github.com/kryast/Crud-2.git/services"

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{s}
}
