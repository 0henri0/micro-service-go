package services

import (
	"nath-backend/internal/models"
	"nath-backend/internal/repositories"
)

type ProductService interface {
	Find(id int) (models.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

// Find implements ProductService.
func (p *productService) Find(id int) (models.Product, error) {
	var product, err = p.productRepo.Find(id)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}
