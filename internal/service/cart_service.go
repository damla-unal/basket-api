package service

import (
	"basket-api/internal/model"
	"basket-api/internal/model/request"
	"basket-api/internal/persistence"
	"context"
	"github.com/go-playground/validator/v10"
)

type CartService interface {
	ShowCustomerCart(ctx context.Context, customerID int) (model.Cart, error)
	AddItemToCart(ctx context.Context, request request.CartItemRequest) error
	//ValidateCustomerCartRequest(request request.CustomerCartRequest) error
}

type CartServiceImp struct {
	cartDAO     persistence.CartDAO
	cartItemDAO persistence.CartItemDAO
	productDAO  persistence.ProductDAO
	Validator   *validator.Validate
}

var _ CartService = (*CartServiceImp)(nil)

func NewCartServiceImp(
	cartDAO persistence.CartDAO,
	cartItemDAO persistence.CartItemDAOPostgres,
	productDAO persistence.ProductDAO,
	validator *validator.Validate) CartServiceImp {
	return CartServiceImp{
		cartDAO:     cartDAO,
		cartItemDAO: cartItemDAO,
		productDAO:  productDAO,
		Validator:   validator,
	}
}

func (c CartServiceImp) ShowCustomerCart(ctx context.Context, customerID int) (model.Cart, error) {
	customerCart, err := c.cartDAO.GetCartByCustomerID(ctx, customerID)
	if err != nil {
		return model.Cart{}, err
	}
	cartItems, err := c.cartItemDAO.GetCartItemsByCartID(ctx, int(customerCart.ID))
	if err != nil {
		return model.Cart{}, err
	}
	customerCart.Items = cartItems
	return customerCart, nil
}

func (c CartServiceImp) AddItemToCart(ctx context.Context, request request.CartItemRequest) error {
	customerCart, err := c.cartDAO.GetCartByCustomerID(ctx, request.CustomerID)
	if err != nil {
		return err
	}

	product, err := c.productDAO.GetProductByID(ctx, request.ProductID)
	if err != nil {
		return err
	}
	updatedCartTotalPrice := customerCart.TotalPrice + product.Price
	err = c.cartItemDAO.UpsertCartItem(ctx, int(customerCart.ID), request.ProductID, updatedCartTotalPrice)
	if err != nil {
		return err
	}
	return nil
}

//func (c CartServiceImp) ValidateCustomerCartRequest(request request.CustomerCartRequest) error {
//	if err := c.Validator.Struct(request); err != nil {
//		return err
//	}
//	return nil
//}
