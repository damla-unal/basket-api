package service

import (
	"basket-api/internal/model"
	"basket-api/internal/model/request"
	"basket-api/internal/persistence"
	"basket-api/internal/util"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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

	//check if the cart is empty or not, empty cart cannot be converted an order
	if len(cart.Items) == 0 {
		return errors.New("cannot create an order: customers' cart is empty")
	}

	orders, err := o.orderDAO.GetOrdersByCustomerID(ctx, request.CustomerID)
	if err != nil {
		return err
	}

	// check if the number of order is greater than 3 and cart price is greater and equal to given amount
	if len(orders) > 3 && cart.TotalPrice >= viper.GetInt("threshold_for_discount.amount") {
		isSuitableForDiscount := checkDiscountConditionsForEveryFourthOrder(&orders)
		if isSuitableForDiscount {
			cart.TotalPrice = o.calculateDiscountForEveryFourthOrder(cart.Items)
		}
	}

	//create an order
	err = o.orderDAO.CreateOrder(ctx, cart)
	if err != nil {
		return err
	}
	return nil
}

func (o OrderServiceImp) calculateDiscountForEveryFourthOrder(items []model.CartItem) int {
	vatDiscountMap := util.GetVatToDiscountMap()
	var newOrderPrice int
	for _, item := range items {
		if item.ProductVat == 1 {
			newOrderPrice += item.Price
			continue
		}
		if item.Discount < vatDiscountMap[item.ProductVat] {
			item.Discount = vatDiscountMap[item.ProductVat]
			priceWithoutDiscount := item.Quantity * item.QTYPrice
			newDiscountPrice := priceWithoutDiscount - (priceWithoutDiscount * item.Discount / 100)
			newOrderPrice += newDiscountPrice
		}
	}
	return newOrderPrice
}

func checkDiscountConditionsForEveryFourthOrder(orders *[]model.Order) bool {
	thresholdForDiscount := viper.GetInt("threshold_for_discount.amount")
	counter := 0
	for _, order := range *orders {
		if order.TotalPrice >= thresholdForDiscount {
			counter++
		}
	}
	if counter%3 == 0 {
		return true
	} else {
		return false
	}
}
