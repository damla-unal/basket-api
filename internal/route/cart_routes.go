package route

import (
	"basket-api/internal/model/request"
	"basket-api/internal/model/response"
	"basket-api/internal/service"
	"basket-api/internal/util/http_helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCartRoutes(r *gin.RouterGroup, cartService service.CartService) *gin.RouterGroup {
	cartRoutes := r.Group("/cart")
	{
		cartRoutes.GET("", showCustomerCart(cartService))
		cartRoutes.POST("/items", addItemToCart(cartService))
	}
	return cartRoutes
}

func showCustomerCart(cartService service.CartService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		customerID, err := http_helpers.GetPositiveIntegerQueryParameter(context, "customer-id")
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		customerCart, err := cartService.ShowCustomerCart(ctx, *customerID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, customerCart)
	}

}

func addItemToCart(cartService service.CartService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		var cartItemRequest request.CartItemRequest
		if err := context.ShouldBindJSON(&cartItemRequest); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := cartService.AddItemToCart(ctx, cartItemRequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		context.JSON(http.StatusOK, response.SuccessfulResponse{Result: true})

	}

}
