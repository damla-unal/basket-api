package model

// CartItem TODO separate it as cartItem model and cartItemWithProductDetail model
type CartItem struct {
	ID           int    `json:"id"`
	Quantity     int    `json:"quantity"`
	CartID       int    `json:"cart_id"`
	Discount     int    `json:"discount"`
	Price        int    `json:"price"`
	ProductID    int    `json:"product_id"`
	ProductTitle string `json:"product_title"`
}
