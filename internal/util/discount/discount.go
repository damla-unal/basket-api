package discount

import (
	"basket-api/internal/model"
	"errors"
)

func CalculatePriceAfterDiscount(discountRate int, price int) (int, error) {
	if discountRate < 0 {
		return 0, errors.New("discount rate should not be negative")

	}
	afterDiscount := price - (price * discountRate / 100)
	return afterDiscount, nil
}

//CalculateDiscountForTheSameProducts checks if there are more than 3 items of the same product,
// then fourth and subsequent ones would have %8 off.
func CalculateDiscountForTheSameProducts(productID int, cart model.Cart, operationType string) int {
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
