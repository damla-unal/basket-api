package service

import (
	"basket-api/internal/model/request"
	"basket-api/internal/persistence"
	"context"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request request.OrderRequest) error
}

type OrderServiceImp struct {
	cartDAO  persistence.CartDAO
	orderDAO persistence.OrderDAO
}

var _ OrderService = (*OrderServiceImp)(nil)

func NewOrderServiceImp(
	cartDAO persistence.CartDAO,
	orderDAO persistence.OrderDAO,
) OrderServiceImp {
	return OrderServiceImp{
		cartDAO:  cartDAO,
		orderDAO: orderDAO,
	}
}

func (o OrderServiceImp) CreateOrder(ctx context.Context, request request.OrderRequest) error {
	//get customers' cart info to create an order from cart items
	cart, err := o.cartDAO.GetCartByCustomerID(ctx, request.CustomerID)
	if err != nil {
		return err
	}
	//create an order
	err = o.orderDAO.CreateOrder(ctx, cart)
	if err != nil {
		return err
	}
	return nil
}
