package persistence

import (
	db "basket-api/internal/dpsql"
	"basket-api/internal/model"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type CartItemDAO interface {
	GetCartItemsByCartID(ctx context.Context, cartID int) ([]model.CartItem, error)
	UpsertCartItem(ctx context.Context, cartID int, productID int, price int64) error
}

type CartItemDAOPostgres struct {
	dbPool *pgxpool.Pool
}

var _ CartItemDAO = (*CartItemDAOPostgres)(nil)

func NewCartItemDAOPostgres(dbPool *pgxpool.Pool) CartItemDAOPostgres {
	return CartItemDAOPostgres{dbPool: dbPool}
}

func (c CartItemDAOPostgres) GetCartItemsByCartID(ctx context.Context, cartID int) ([]model.CartItem, error) {

	cartItems := make([]model.CartItem, 0)
	resErr := c.dbPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		dbCartItems, err := db.New(conn).GetCartItemsByCartID(ctx, int64(cartID))
		if err != nil {
			return errors.Wrap(err, "unable to get items of cart")
		}
		for _, dbCartItem := range dbCartItems {
			cartItems = append(cartItems, createCartItemModelFromDpSQLModel(dbCartItem))
		}
		return nil

	})
	return cartItems, resErr
}

func (c CartItemDAOPostgres) UpsertCartItem(ctx context.Context, cartID int, productID int, price int64) error {

	resErr := c.dbPool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		queries := db.New(tx)

		upsertParams := db.UpsertCartItemParams{
			Quantity:  1,
			CartID:    int64(cartID),
			ProductID: int64(productID),
		}
		err := queries.UpsertCartItem(ctx, upsertParams)
		if err != nil {
			return errors.Wrap(err, "unable to upsert cart item")
		}
		updateCartParams := db.UpdateCartParams{
			Price:    price,
			Vat:      0,
			Discount: 0,
			Status:   "saved",
			ID:       int64(cartID),
		}
		err = queries.UpdateCart(ctx, updateCartParams)
		if err != nil {
			return errors.Wrap(err, "unable to update cart")
		}
		return nil
	})
	return resErr
}

func createCartItemModelFromDpSQLModel(dbCartItem db.GetCartItemsByCartIDRow) model.CartItem {
	return model.CartItem{
		Quantity:     dbCartItem.Quantity,
		CartID:       dbCartItem.CartID,
		Discount:     dbCartItem.Discount,
		ProductID:    dbCartItem.ProductID,
		ProductTitle: dbCartItem.ProductTitle.String,
	}

}
