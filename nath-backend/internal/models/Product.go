package models

type Product struct {
	ID                      uint       `gorm:"primaryKey"`
	Code                    string     `gorm:"unique"`
	Name                    string
	Slug                    string
	Image                   string
	Count                   int
	CountSell               int
	Price                   float64
	PriceYuan               float64
	ShippingCost            float64
	LinkProduct             string
	LinkResource            string
	Description             string
	IsEcommerce             bool
	IsDraft                 bool
}
