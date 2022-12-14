package persistence

import (
	"basket-api/internal/dpsql"
	"basket-api/internal/model"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type OrderDAO interface {
	CreateOrder(ctx context.Context, cartToOrder model.Cart) error
	GetOrdersByCustomerID(ctx context.Context, customerID int) ([]model.Order, error)
}

type OrderDAOPostgres struct {
	dbPool *pgxpool.Pool
}

var _ OrderDAO = (*OrderDAOPostgres)(nil)

func NewOrderDAOPostgres(dbPool *pgxpool.Pool) OrderDAOPostgres {
	return OrderDAOPostgres{dbPool: dbPool}
}

func (o OrderDAOPostgres) CreateOrder(ctx context.Context, cartToOrder model.Cart) error {
	resErr := o.dbPool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		queries := db.New(tx)

		createOrderParams := db.CreateOrderParams{
			TotalPrice: int64(cartToOrder.TotalPrice),
			CustomerID: int64(cartToOrder.CustomerID),
		}
		err := queries.CreateOrder(ctx, createOrderParams)
		if err != nil {
			return errors.Wrap(err, "unable to create order")
		}

		//empty to cart, remove all cart-item rows for the customer
		err = queries.DeleteAllCartItem(ctx, int64(cartToOrder.ID))
		if err != nil {
			return errors.Wrap(err, "unable to delete all cart item")
		}

		//update customers' cart to initial state
		updateCartParams := db.UpdateCartParams{
			Price:    0,
			Vat:      0,
			Discount: 0,
			ID:       int64(cartToOrder.ID),
		}
		err = queries.UpdateCart(ctx, updateCartParams)
		if err != nil {
			return errors.Wrap(err, "unable to update cart")
		}

		return nil
	})
	return resErr
}

func (o OrderDAOPostgres) GetOrdersByCustomerID(ctx context.Context, customerID int) ([]model.Order, error) {
	orders := make([]model.Order, 0)
	resErr := o.dbPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		dbOrders, err := db.New(conn).GetOrdersByCustomerID(ctx, int64(customerID))
		if err != nil {
			return errors.Wrap(err, "unable to get orders of customer")
		}
		for _, dbOrder := range dbOrders {
			orders = append(orders, createOrderModelFromDbModel(dbOrder))
		}
		return nil

	})
	return orders, resErr
}

func createOrderModelFromDbModel(order db.Order) model.Order {
	return model.Order{
		CustomerID: int(order.CustomerID),
		TotalPrice: int(order.TotalPrice),
	}

}
