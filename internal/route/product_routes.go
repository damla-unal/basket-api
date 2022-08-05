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

func listProducts(productDAO persistence.ProductDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		products, err := productDAO.ListProducts(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		productsResponse := make([]response.ProductResponse, 0)
		for _, p := range products {
			productsResponse = append(productsResponse, createProductResponse(p))
		}

		c.JSON(http.StatusOK, response.ProductsResponse{Products: productsResponse})
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
