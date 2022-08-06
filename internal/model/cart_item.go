package model

type CartItem struct {
	Quantity     int64  `json:"quantity"`
	CartID       int64  `json:"cart_id"`
	Discount     int64  `json:"discount"`
	ProductID    int64  `json:"product_id"`
	ProductTitle string `json:"product_title"`
}
