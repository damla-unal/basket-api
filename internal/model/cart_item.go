package model

type CartItem struct {
	ID           int    `json:"id"`
	Quantity     int    `json:"quantity"`
	CartID       int    `json:"cart_id"`
	Discount     int    `json:"discount"`
	Price        int    `json:"price"`
	ProductID    int    `json:"product_id"`
	ProductTitle string `json:"product_title"`
	ProductVat   int    `json:"product_vat"`
	QTYPrice     int    `json:"qty_price"`
}
