package services

import (
	"github.com/kryast/Crud-2.git/models"
	"github.com/kryast/Crud-2.git/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) ProductService {
	return &productService{r}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}
