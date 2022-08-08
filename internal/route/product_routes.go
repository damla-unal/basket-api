package route

import (
	"basket-api/internal/model"
	"basket-api/internal/model/response"
	"basket-api/internal/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddProductRoutes add product group to gin engine
func AddProductRoutes(r *gin.RouterGroup, productDAO persistence.ProductDAO) *gin.RouterGroup {
	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", listProducts(productDAO))
	}
	return productRoutes
}

//listProducts endpoint list all products as a list.
// GET localhost:8080/api/products
func listProducts(productDAO persistence.ProductDAO) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()

		products, err := productDAO.ListProducts(ctx)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.FailedResponse{Error: err.Error()})
			return
		}

		productsResponse := make([]response.ProductResponse, 0)
		for _, p := range products {
			productsResponse = append(productsResponse, createProductResponse(p))
		}

		context.JSON(http.StatusOK, response.ProductsResponse{Products: productsResponse})
	}
}

func createProductResponse(product model.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:    product.ID,
		Title: product.Title,
		Price: product.Price,
		Vat:   product.Vat,
	}
}
