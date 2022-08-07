package service

import (
	"basket-api/internal/model"
	"basket-api/internal/model/request"
	"basket-api/mocks"
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderServiceImp_CreateOrder(t *testing.T) {
	t.Run("Successfully create order",
		func(t *testing.T) {
			service, cartService, orderDAO, ctx := setupOrderStubs()

			customerCart := model.Cart{
				CustomerID: 1,
				Items:      getCartItems(),
				TotalPrice: 8540,
			}
			cartService.On("GetCustomerCart", mock.Anything, mock.AnythingOfType("int")).
				Return(customerCart, nil)

			orderDAO.On("GetOrdersByCustomerID", mock.Anything, mock.AnythingOfType("int")).
				Return(getOrders(), nil)

			orderDAO.On("CreateOrder", mock.Anything, mock.AnythingOfType("model.Cart")).
				Return(nil)

			err := service.CreateOrder(ctx, request.OrderRequest{
				CustomerID: 1,
			})
			require.NoError(t, err)
			cartService.AssertNumberOfCalls(t, "GetCustomerCart", 1)
			orderDAO.AssertNumberOfCalls(t, "GetOrdersByCustomerID", 1)
			orderDAO.AssertNumberOfCalls(t, "CreateOrder", 1)
		})

	t.Run("Failed create order, when customers' cart is empty",
		func(t *testing.T) {
			service, cartService, orderDAO, ctx := setupOrderStubs()

			customerCart := model.Cart{
				CustomerID: 1,
			}
			cartService.On("GetCustomerCart", mock.Anything, mock.AnythingOfType("int")).
				Return(customerCart, nil)

			err := service.CreateOrder(ctx, request.OrderRequest{
				CustomerID: 1,
			})
			require.Contains(t, err.Error(), "cannot create an order: customers' cart is empty")
			cartService.AssertNumberOfCalls(t, "GetCustomerCart", 1)
			orderDAO.AssertNumberOfCalls(t, "GetOrdersByCustomerID", 0)
			orderDAO.AssertNumberOfCalls(t, "CreateOrder", 0)
		})
}

func setupOrderStubs() (OrderService, *mocks.CartService, *mocks.OrderDAO, context.Context) {
	cartServiceMock := &mocks.CartService{}
	orderDAOMock := &mocks.OrderDAO{}
	service := NewOrderServiceImp(cartServiceMock, orderDAOMock)
	return service, cartServiceMock, orderDAOMock, context.Background()

}

func getOrders() []model.Order {
	return []model.Order{
		{
			CustomerID: 1,
			TotalPrice: 12000,
		},
		{
			CustomerID: 1,
			TotalPrice: 10000,
		},
		{
			CustomerID: 1,
			TotalPrice: 145,
		},
		{
			CustomerID: 1,
			TotalPrice: 14500,
		},
	}

}
