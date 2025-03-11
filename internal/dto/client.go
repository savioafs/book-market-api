package dto

type ClientInputDTO struct {
	Name  string `json:"name"  `
	Email string `json:"email"`
	Phone string `json:"phone" `
}

type ClientOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"  `
	Email string `json:"email"`
	Phone string `json:"phone" `
}

type ClientOutputForSaleDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
