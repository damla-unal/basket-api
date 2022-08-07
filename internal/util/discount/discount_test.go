package discount

import (
	"basket-api/internal/model"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculatePriceAfterDiscount(t *testing.T) {
	t.Run("Successfully calculated", func(t *testing.T) {
		discountRate := 8
		price := 1000
		actual, err := CalculatePriceAfterDiscount(discountRate, price)
		require.NoError(t, err)
		require.Equal(t, 920, actual)
	})
	t.Run("With negative value, failed", func(t *testing.T) {
		discountRate := -1
		price := 1000
		_, err := CalculatePriceAfterDiscount(discountRate, price)
		require.Error(t, err)
	})

}

func TestCalculateDiscountForTheSameProducts(t *testing.T) {
	t.Run("Successfully calculated discount rate for adding operation: "+
		"if there are more than 3 items of the same product, newly would have %8 off.",
		func(t *testing.T) {
			cart := model.Cart{
				Items: getCartItems(),
			}
			productID := 1
			expectedDiscount := 8
			actualDiscount := CalculateDiscountForTheSameProducts(productID, cart, "add")
			require.Equal(t, expectedDiscount, actualDiscount)

		})
	t.Run("Successfully calculated discount rate for adding operation: "+
		"if there are less than 3 items of the same product, discount should not be applied",
		func(t *testing.T) {
			cart := model.Cart{
				Items: getCartItems(),
			}
			productID := 3
			expectedDiscount := 0
			actualDiscount := CalculateDiscountForTheSameProducts(productID, cart, "add")
			require.Equal(t, expectedDiscount, actualDiscount)

		})
	t.Run("Successfully calculated discount rate for removing operation: "+
		"if there are less than and equal to 4 items of the same product, discount is canceled in the deletion process.",
		func(t *testing.T) {
			cart := model.Cart{
				Items: getCartItems(),
			}
			productID := 1
			expectedDiscount := 0
			actualDiscount := CalculateDiscountForTheSameProducts(productID, cart, "remove")
			require.Equal(t, expectedDiscount, actualDiscount)

		})
	t.Run("Successfully calculated discount rate for removing operation: "+
		"if there are more than 4 items of the same product, discount is continued in deletion operation",
		func(t *testing.T) {
			cart := model.Cart{
				Items: getCartItems(),
			}
			productID := 5
			expectedDiscount := 8
			actualDiscount := CalculateDiscountForTheSameProducts(productID, cart, "remove")
			require.Equal(t, expectedDiscount, actualDiscount)

		})

}

func TestCheckDiscountConditionsForEveryFourthOrder(t *testing.T) {
	t.Run("It should return false if the number of orders greater than or equal to the given amount is not a multiple of 3 and 3.",
		func(t *testing.T) {
			viper.Set("threshold_for_discount.amount", 10000)
			orders := getOrders()
			actual := CheckDiscountConditionsForEveryFourthOrder(&orders)
			require.Equal(t, false, actual)
		})

	t.Run("It should return true if the number of orders greater than or equal to the given quantity is a multiple of 3 and 3.",
		func(t *testing.T) {
			viper.Set("threshold_for_discount.amount", 10000)
			orders := getOrders()
			orders = append(orders, model.Order{
				CustomerID: 1,
				TotalPrice: 14567,
			})
			actual := CheckDiscountConditionsForEveryFourthOrder(&orders)
			require.Equal(t, true, actual)
		})

}

func TestCalculateDiscountForEveryFourthOrder(t *testing.T) {
	t.Run("It should return false if the number of orders greater than or equal to the given amount is not a multiple of 3 and 3.",
		func(t *testing.T) {
			cartItems := getCartItems()
			expectedCartTotalPrice := 18600
			actual := CalculateDiscountForEveryFourthOrder(cartItems)
			require.Equal(t, expectedCartTotalPrice, actual)
		})
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
			Quantity:     6,
			CartID:       1,
			Discount:     8,
			Price:        12000,
			ProductID:    5,
			ProductTitle: "Test product",
			ProductVat:   8,
			QTYPrice:     2000,
		},
	}
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
	}

}
