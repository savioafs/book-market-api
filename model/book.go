package model

type Book struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Publisher     string  `json:"publisher"`
	ISBN          string  `json:"isbn"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
	Category      string  `json:"category"`
	PublishedYear int     `json:"published_year"`
	Description   string  `json:"description"`
}
