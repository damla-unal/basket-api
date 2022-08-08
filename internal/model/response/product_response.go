package response

type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
}

type ProductResponse struct {
	ID    int64  `example:"2" json:"id"`
	Title string `example:"Garlic Knots" json:"title"`
	Price int64  `example:"699" json:"price"`
	Vat   int64  `example:"8" json:"vat"`
}
