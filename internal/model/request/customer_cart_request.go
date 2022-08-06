package request

type CartItemRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
	ProductID  int `json:"product_id" binding:"required"`
}
