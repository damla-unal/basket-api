package route

import (
	"basket-api/internal/model/request"
	"basket-api/internal/model/response"
	"basket-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddOrderRoutes(r *gin.RouterGroup, orderService service.OrderService) *gin.RouterGroup {
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("", createOrder(orderService))
	}
	return orderRoutes
}

func createOrder(orderService service.OrderService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		var orderRequest request.OrderRequest
		if err := context.ShouldBindJSON(&orderRequest); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := orderService.CreateOrder(ctx, orderRequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, response.SuccessfulResponse{Result: true})

	}
}
