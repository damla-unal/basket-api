package service

import (
	"basket-api/internal/model"
	"basket-api/internal/model/request"
	"basket-api/internal/persistence"
	"context"
)

type CartService interface {
	GetCustomerCart(ctx context.Context, customerID int) (model.Cart, error)
	GetCartByID(ctx context.Context, id int) (model.Cart, error)
	AddItemToCart(ctx context.Context, request request.CartItemRequest) error
	DeleteItemFromCart(ctx context.Context, itemID int) error
}

type CartServiceImp struct {
	cartDAO     persistence.CartDAO
	cartItemDAO persistence.CartItemDAO
	productDAO  persistence.ProductDAO
}

var _ CartService = (*CartServiceImp)(nil)

func NewCartServiceImp(
	cartDAO persistence.CartDAO,
	cartItemDAO persistence.CartItemDAOPostgres,
	productDAO persistence.ProductDAO,
) CartServiceImp {
	return CartServiceImp{
		cartDAO:     cartDAO,
		cartItemDAO: cartItemDAO,
		productDAO:  productDAO,
	}
}

func (c CartServiceImp) GetCustomerCart(ctx context.Context, customerID int) (model.Cart, error) {
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

func (c CartServiceImp) GetCartByID(ctx context.Context, id int) (model.Cart, error) {
	cart, err := c.cartDAO.GetCartByID(ctx, id)
	if err != nil {
		return model.Cart{}, err
	}
	cartItems, err := c.cartItemDAO.GetCartItemsByCartID(ctx, id)
	if err != nil {
		return model.Cart{}, err
	}
	cart.Items = cartItems
	return cart, nil

}

func (c CartServiceImp) AddItemToCart(ctx context.Context, request request.CartItemRequest) error {
	customerCart, err := c.GetCustomerCart(ctx, request.CustomerID)
	if err != nil {
		return err
	}

	product, err := c.productDAO.GetProductByID(ctx, request.ProductID)
	if err != nil {
		return err
	}
	discountRate := calculateDiscountForTheSameProducts(request.ProductID, customerCart, "add")
	afterDiscount := calculatePriceAfterDiscount(discountRate, int(product.Price))

	updatedCartTotalPrice := customerCart.TotalPrice + afterDiscount
	cartItemToUpsert := model.CartItem{
		Quantity:  1,
		CartID:    customerCart.ID,
		Discount:  discountRate,
		Price:     afterDiscount,
		ProductID: request.ProductID,
	}
	err = c.cartItemDAO.UpsertCartItem(ctx, cartItemToUpsert, updatedCartTotalPrice)
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

	cart, err := c.GetCartByID(ctx, foundCartItem.CartID)
	if err != nil {
		return err
	}

	// if the quantity greater than one, reduce by one (decreased quantity)
	if foundCartItem.Quantity > 1 {
		updatedDiscount := calculateDiscountForTheSameProducts(foundCartItem.ProductID, cart, "remove")
		removedPrice := calculatePriceAfterDiscount(foundCartItem.Discount, foundCartItem.QTYPrice)
		afterDeleteItemPrice := foundCartItem.Price - removedPrice
		afterDeleteCartPrice := cart.TotalPrice - removedPrice

		updatedQuantity := foundCartItem.Quantity - 1
		updatedCartItem := model.CartItem{
			ID:       itemID,
			Quantity: updatedQuantity,
			CartID:   foundCartItem.CartID,
			Discount: updatedDiscount,
			Price:    afterDeleteItemPrice,
		}
		err := c.cartItemDAO.UpdateCartItem(
			ctx,
			updatedCartItem,
			afterDeleteCartPrice,
		)
		if err != nil {
			return err
		}
	} else { // if the quantity is equal to 1, remove this cart item
		afterDeleteCartPrice := cart.TotalPrice - foundCartItem.QTYPrice
		err = c.cartItemDAO.DeleteCartItem(ctx, itemID, foundCartItem.CartID, afterDeleteCartPrice)
		if err != nil {
			return err
		}
	}
	return nil
}

//checkDiscountForTheSameProducts checks if there are more than 3 items of the same product,
// then fourth and subsequent ones would have %8 off.
func calculateDiscountForTheSameProducts(productID int, cart model.Cart, operationType string) int {

	for _, item := range cart.Items {
		if item.ProductID == productID {
			switch operationType {
			case "add":
				if item.Quantity >= 3 {
					return 8
				}
			case "remove":
				if item.Quantity <= 4 {
					return 0
				} else {
					return 8
				}
			default:
				return 0
			}
		}
	}
	return 0
}

func calculatePriceAfterDiscount(discountRate int, price int) int {
	afterDiscount := price - (price * discountRate / 100)
	return afterDiscount
}
