package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-2.git/models"
	"github.com/kryast/Crud-2.git/services"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{s}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, _ := h.service.GetProducts()
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)
	if err := h.service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, product)
}
