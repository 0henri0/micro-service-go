package repositories

import (
	"gorm.io/gorm"
	"nath-backend/internal/models"
)

type ProductRepository interface {
	Find(id int) (models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

// Find implements ProductRepository.
func (p *productRepository) Find(id int) (models.Product, error) {
	var product models.Product
	if err := p.db.First(&product, id).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
