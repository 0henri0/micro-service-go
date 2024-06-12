package repositories

import (
	"database/sql"
	"fmt"
	"micro-service-go/internal/models"
	"strings"
)

type ProductRepository interface {
	SaveProducts([]models.Product) error
}

type productRepository struct {
	db *sql.DB
}

func (p productRepository) SaveProducts(products []models.Product) error {
	query := `INSERT INTO products (name, price, origin_url, slug) VALUES `
	var values []interface{}
	for _, product := range products {
		query += fmt.Sprintf("('%s', '%d', '%s', '%s'),", product.Name, product.Price, product.OriginUrl, product.OriginUrl)
		values = append(values, product.Name, product.Price, product.OriginUrl)
	}
	query = query[:len(query)-1]
	query = strings.Trim(query, ",")
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
