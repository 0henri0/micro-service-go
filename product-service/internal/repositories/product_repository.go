package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"micro-service-go/internal/models"
)

type ProductRepository interface {
	SaveProducts([]models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func (p productRepository) SaveProducts(products []models.Product) error {
	if len(products) == 0 {
		return nil
	}

	result := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "origin_url"}},
		DoNothing: true,
	}).Create(&products)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Insert completed")
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
