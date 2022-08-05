package model

type Product struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Price int64  `json:"price"`
	Vat   int64  `json:"vat"`
}
