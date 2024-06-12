package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	OriginUrl   string `json:"origin_url"`
}
