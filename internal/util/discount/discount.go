package discount

import (
	"basket-api/internal/model"
	"basket-api/internal/util"
	"errors"
	"github.com/spf13/viper"
)

//CalculatePriceAfterDiscount returns the price after discount
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

//CheckDiscountConditionsForEveryFourthOrder checks the order history whether there are enough orders
//that exceed the given amount (: threshold_for_discount.amount)
// Every fourth orders must exceed the given amount for the discount to be applied.
// Discount will be applied to each product(except %1 vat products) in the every fourth order.
func CheckDiscountConditionsForEveryFourthOrder(orders *[]model.Order) bool {
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

//CalculateDiscountForEveryFourthOrder applies a discount to the products in the cart according to the VAT rates.
//And after all the discounts, it returns the total discounted price of the cart.
func CalculateDiscountForEveryFourthOrder(items []model.CartItem) int {
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
