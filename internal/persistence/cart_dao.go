package persistence

import (
	db "basket-api/internal/dpsql"
	"basket-api/internal/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type CartDAO interface {
	GetCartByCustomerID(ctx context.Context, customerID int) (model.Cart, error)
}

type CartDAOPostgres struct {
	dbPool *pgxpool.Pool
}

var _ CartDAO = (*CartDAOPostgres)(nil)

func NewCartDAOPostgres(dbPool *pgxpool.Pool) CartDAOPostgres {
	return CartDAOPostgres{dbPool: dbPool}
}

func (c CartDAOPostgres) GetCartByCustomerID(ctx context.Context, customerID int) (model.Cart, error) {

	var customerCart model.Cart
	resErr := c.dbPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		cart, err := db.New(conn).GetSavedCartByCustomerID(ctx, int64(customerID))
		if err != nil {
			return errors.Wrap(err, "unable to get cart of customer")
		}
		customerCart = createCartModelFromDpSQLModel(cart)
		return nil

	})
	return customerCart, resErr
}

func createCartModelFromDpSQLModel(dbCart db.GetSavedCartByCustomerIDRow) model.Cart {
	return model.Cart{
		ID:           dbCart.ID,
		TotalPrice:   dbCart.TotalPrice,
		Vat:          dbCart.Vat,
		Discount:     dbCart.Discount,
		CustomerID:   dbCart.CustomerID,
		CustomerName: dbCart.CustomerName.String,
		Status:       model.CartStatus(dbCart.Status),
		CreatedAt:    dbCart.CreatedAt.Time,
		UpdatedAt:    dbCart.UpdatedAt.Time,
	}
}
