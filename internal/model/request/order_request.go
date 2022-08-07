package request

type OrderRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
}
