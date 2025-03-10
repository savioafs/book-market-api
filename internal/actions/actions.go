package actions

type UpdateBookStockAction string

var (
	BookRenew UpdateBookStockAction = "RENEW"
	BookSale  UpdateBookStockAction = "SALE"
)

func (a UpdateBookStockAction) IsValid() bool {
	switch a {
	case BookRenew, BookSale:
		return true
	default:
		return false
	}
}
