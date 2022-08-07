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
	GetCartByID(ctx context.Context, ID int) (model.Cart, error)
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
		cart, err := db.New(conn).GetCartByCustomerID(ctx, int64(customerID))
		if err != nil {
			return errors.Wrap(err, "unable to get cart of customer")
		}
		customerCart = createCartModelFromDpSQLModel(cart)
		return nil

	})
	return customerCart, resErr
}

func (c CartDAOPostgres) GetCartByID(ctx context.Context, ID int) (model.Cart, error) {
	var foundCart model.Cart
	resErr := c.dbPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		cart, err := db.New(conn).GetCartByCustomerID(ctx, int64(ID))
		if err != nil {
			return errors.Wrap(err, "unable to get cart by ID")
		}
		foundCart = createCartModelFromDpSQLModel(cart)
		return nil

	})
	return foundCart, resErr
}

func createCartModelFromDpSQLModel(dbCart db.GetCartByCustomerIDRow) model.Cart {
	return model.Cart{
		ID:           int(dbCart.ID),
		TotalPrice:   int(dbCart.TotalPrice),
		Vat:          int(dbCart.Vat),
		Discount:     int(dbCart.Discount),
		CustomerID:   int(dbCart.CustomerID),
		CustomerName: dbCart.CustomerName.String,
		CreatedAt:    dbCart.CreatedAt.Time,
		UpdatedAt:    dbCart.UpdatedAt.Time,
	}
}
