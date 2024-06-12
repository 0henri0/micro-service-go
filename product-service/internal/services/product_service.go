package services

import (
	"micro-service-go/internal/models"
	"micro-service-go/internal/repositories"
)

type ProductService interface {
	SaveProducts([]models.Product) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func (p productService) SaveProducts(products []models.Product) error {
	err := p.productRepo.SaveProducts(products)
	if err != nil {
		return err
	}
	return nil
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}
