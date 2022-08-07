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
	DeleteItemFromCart(ctx context.Context, itemID int) error
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
	cartItems, err := c.cartItemDAO.GetCartItemsByCartID(ctx, customerCart.ID)
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
	updatedCartTotalPrice := customerCart.TotalPrice + int(product.Price)
	err = c.cartItemDAO.UpsertCartItem(ctx, customerCart.ID, request.ProductID, updatedCartTotalPrice, int(product.Price))
	if err != nil {
		return err
	}
	return nil
}

func (c CartServiceImp) DeleteItemFromCart(ctx context.Context, itemID int) error {
	foundCartItem, err := c.cartItemDAO.GetCartItemByID(ctx, itemID)
	if err != nil {
		return err
	}

	cart, err := c.cartDAO.GetCartByID(ctx, foundCartItem.CartID)
	if err != nil {
		return err
	}

	priceOfQty := foundCartItem.Price / foundCartItem.Quantity
	updatedCartTotalPrice := cart.TotalPrice - priceOfQty

	// if the quantity greater than one, reduce by one (decreased quantity)
	if foundCartItem.Quantity > 1 {
		updatedQuantity := foundCartItem.Quantity - 1
		err := c.cartItemDAO.UpdateCartItem(
			ctx,
			itemID,
			updatedQuantity,
			foundCartItem.Discount,
			foundCartItem.Price-priceOfQty,
			updatedCartTotalPrice,
			foundCartItem.CartID,
		)
		if err != nil {
			return err
		}
	} else { // if the quantity is equal to 1, remove this cart item
		err = c.cartItemDAO.DeleteCartItem(ctx, itemID, foundCartItem.CartID, updatedCartTotalPrice)
		if err != nil {
			return err
		}
	}
	return nil
}

//func (c CartServiceImp) ValidateCustomerCartRequest(request request.CustomerCartRequest) error {
//	if err := c.Validator.Struct(request); err != nil {
//		return err
//	}
//	return nil
//}
