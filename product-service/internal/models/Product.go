package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Price       int    `gorm:"check:price >= 0"`
	Description string `json:"description"`
	Slug        string `gorm:"size:255"`
	OriginUrl   string `gorm:"size:255;unique"`
}
