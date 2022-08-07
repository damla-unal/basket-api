package service

import (
	"basket-api/internal/model"
	"basket-api/internal/model/request"
	"basket-api/mocks"
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCartServiceImp_GetCustomerCart(t *testing.T) {
	t.Run("When dao methods return successfully, it should return customers' cart with items",
		func(t *testing.T) {
			service, cartDAO, cartItemDAO, _, ctx := setupCartStubs()

			customerCart := model.Cart{CustomerID: 1}
			cartDAO.On("GetCartByCustomerID", mock.Anything, mock.AnythingOfType("int")).
				Return(customerCart, nil)

			cartItemDAO.On("GetCartItemsByCartID", mock.Anything, mock.AnythingOfType("int")).
				Return(getCartItems(), nil)

			customerCart.Items = getCartItems()
			actual, err := service.GetCustomerCart(ctx, 1)
			require.NoError(t, err)
			require.EqualValues(t, customerCart, actual)
		})

	t.Run("When cartDao.GetCartByCustomerID method return an error, it should return error",
		func(t *testing.T) {
			service, cartDAO, cartItemDAO, _, ctx := setupCartStubs()

			customerCart := model.Cart{}
			cartDAO.On("GetCartByCustomerID", mock.Anything, mock.AnythingOfType("int")).
				Return(customerCart, errors.New("error from dao"))

			_, err := service.GetCustomerCart(ctx, 1)
			require.Error(t, err, "error from dao")
			cartDAO.AssertNumberOfCalls(t, "GetCartByCustomerID", 1)
			cartItemDAO.AssertNumberOfCalls(t, "GetCartItemsByCartID", 0)
		})
}

func TestCartServiceImp_AddItemToCart(t *testing.T) {
	t.Run("Successfully add item to cart",
		func(t *testing.T) {
			service, cartDAO, cartItemDAO, productDAO, ctx := setupCartStubs()

			customerCart := model.Cart{
				CustomerID: 1,
				Items:      getCartItems(),
				TotalPrice: 8540,
			}
			cartDAO.On("GetCartByCustomerID", mock.Anything, mock.AnythingOfType("int")).
				Return(customerCart, nil)

			cartItemDAO.On("GetCartItemsByCartID", mock.Anything, mock.AnythingOfType("int")).
				Return(getCartItems(), nil)

			productDAO.On("GetProductByID", mock.Anything, mock.AnythingOfType("int")).
				Return(getProduct(), nil)

			cartItemDAO.On("UpsertCartItem", mock.Anything, mock.AnythingOfType("model.CartItem"), mock.AnythingOfType("int")).
				Return(nil)

			err := service.AddItemToCart(ctx, request.CartItemRequest{
				CustomerID: 1,
				ProductID:  1,
			})
			require.NoError(t, err)
			cartDAO.AssertNumberOfCalls(t, "GetCartByCustomerID", 1)
			productDAO.AssertNumberOfCalls(t, "GetProductByID", 1)
			cartItemDAO.AssertNumberOfCalls(t, "UpsertCartItem", 1)
		})
}

func TestCartServiceImp_DeleteItemFromCart(t *testing.T) {
	t.Run("Successfully update cart item(decrease the item quantity and price and update discount), when cart item quantity is greater than 1",
		func(t *testing.T) {
			service, cartDAO, cartItemDAO, _, ctx := setupCartStubs()

			cartItemDAO.On("GetCartItemByID", mock.Anything, mock.AnythingOfType("int")).
				Return(model.CartItem{
					ID:           2,
					Quantity:     4,
					CartID:       1,
					Discount:     8,
					Price:        7840,
					ProductID:    1,
					ProductTitle: "Apple 20W USB-C Power Adapter",
					ProductVat:   8,
					QTYPrice:     2000,
				}, nil)

			cartDAO.On("GetCartByID", mock.Anything, mock.AnythingOfType("int")).
				Return(model.Cart{ID: 1}, nil)

			cartItemDAO.On("GetCartItemsByCartID", mock.Anything, mock.AnythingOfType("int")).
				Return(getCartItems(), nil)

			cartItemDAO.On("UpdateCartItem", mock.Anything, mock.AnythingOfType("model.CartItem"), mock.AnythingOfType("int")).
				Return(nil)

			err := service.DeleteItemFromCart(ctx, 2)
			require.NoError(t, err)
			cartItemDAO.AssertNumberOfCalls(t, "GetCartItemByID", 1)
			cartDAO.AssertNumberOfCalls(t, "GetCartByID", 1)
			cartItemDAO.AssertNumberOfCalls(t, "GetCartItemsByCartID", 1)
			cartItemDAO.AssertNumberOfCalls(t, "UpdateCartItem", 1)
			cartItemDAO.AssertNumberOfCalls(t, "DeleteCartItem", 0)
		})

	t.Run("Successfully delete cart item and update cart total price, when cart item quantity is equal to 1",
		func(t *testing.T) {
			service, cartDAO, cartItemDAO, _, ctx := setupCartStubs()

			cartItemDAO.On("GetCartItemByID", mock.Anything, mock.AnythingOfType("int")).
				Return(model.CartItem{
					ID:           3,
					Quantity:     1,
					CartID:       1,
					Discount:     0,
					Price:        100,
					ProductID:    5,
					ProductTitle: "Gonesh Purrrfect Pet Cat Incense Test2",
					ProductVat:   1,
					QTYPrice:     100,
				}, nil)

			cartDAO.On("GetCartByID", mock.Anything, mock.AnythingOfType("int")).
				Return(model.Cart{ID: 1, TotalPrice: 8540}, nil)

			cartItemDAO.On("GetCartItemsByCartID", mock.Anything, mock.AnythingOfType("int")).
				Return(getCartItems(), nil)

			cartItemDAO.On("DeleteCartItem", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
				Return(nil)

			err := service.DeleteItemFromCart(ctx, 3)
			require.NoError(t, err)
			cartItemDAO.AssertNumberOfCalls(t, "GetCartItemByID", 1)
			cartDAO.AssertNumberOfCalls(t, "GetCartByID", 1)
			cartItemDAO.AssertNumberOfCalls(t, "GetCartItemsByCartID", 1)
			cartItemDAO.AssertNumberOfCalls(t, "UpdateCartItem", 0)
			cartItemDAO.AssertNumberOfCalls(t, "DeleteCartItem", 1)
		})
}

func setupCartStubs() (CartService, *mocks.CartDAO, *mocks.CartItemDAO, *mocks.ProductDAO, context.Context) {
	cartDAOMock := &mocks.CartDAO{}
	cartItemDAOMock := &mocks.CartItemDAO{}
	productDAOMock := &mocks.ProductDAO{}
	service := NewCartServiceImp(cartDAOMock, cartItemDAOMock, productDAOMock)
	return service, cartDAOMock, cartItemDAOMock, productDAOMock, context.Background()

}

func getCartItems() []model.CartItem {
	return []model.CartItem{
		{
			ID:           1,
			Quantity:     2,
			CartID:       1,
			Discount:     0,
			Price:        600,
			ProductID:    3,
			ProductTitle: "Gonesh Purrrfect Pet Cat Incense 30ct",
			ProductVat:   1,
			QTYPrice:     300,
		},
		{
			ID:           2,
			Quantity:     4,
			CartID:       1,
			Discount:     8,
			Price:        7840,
			ProductID:    1,
			ProductTitle: "Apple 20W USB-C Power Adapter",
			ProductVat:   8,
			QTYPrice:     2000,
		},
		{
			ID:           3,
			Quantity:     1,
			CartID:       1,
			Discount:     0,
			Price:        100,
			ProductID:    5,
			ProductTitle: "Gonesh Purrrfect Pet Cat Incense Test2",
			ProductVat:   1,
			QTYPrice:     100,
		},
	}
}

func getProduct() model.Product {
	return model.Product{
		ID:    1,
		Title: "Apple 20W USB-C Power Adapter",
		Price: 2000,
		Vat:   8,
	}

}
