package service

import (
	"basket-api/internal/model/request"
	"basket-api/internal/persistence"
	"context"
	"github.com/pkg/errors"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request request.OrderRequest) error
}

type OrderServiceImp struct {
	cartService CartService
	orderDAO    persistence.OrderDAO
}

var _ OrderService = (*OrderServiceImp)(nil)

func NewOrderServiceImp(
	cartService CartService,
	orderDAO persistence.OrderDAO,
) OrderServiceImp {
	return OrderServiceImp{
		cartService: cartService,
		orderDAO:    orderDAO,
	}
}

func (o OrderServiceImp) CreateOrder(ctx context.Context, request request.OrderRequest) error {
	//get customers' cart info to create an order from cart items
	cart, err := o.cartService.GetCustomerCart(ctx, request.CustomerID)
	if err != nil {
		return err
	}

	if len(cart.Items) == 0 {
		return errors.New("cannot create an order: customers' cart is empty")
	}
	//create an order
	err = o.orderDAO.CreateOrder(ctx, cart)
	if err != nil {
		return err
	}
	return nil
}
