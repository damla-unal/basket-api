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
	GetCartItemByID(ctx context.Context, ID int) (model.CartItem, error)
	UpsertCartItem(ctx context.Context, cartID int, productID int, price int, productPrice int) error
	UpdateCartItem(ctx context.Context, ID int, quantity int, discount int, updatedItemPrice int, cartPrice int, cartID int) error
	DeleteCartItem(ctx context.Context, ID int, cartID int, cartPrice int) error
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
			cartItems = append(cartItems, createCartItemModelFromGetCartItemsByCartIDRow(dbCartItem))
		}
		return nil

	})
	return cartItems, resErr
}

func (c CartItemDAOPostgres) UpsertCartItem(ctx context.Context, cartID int, productID int, price int, productPrice int) error {

	resErr := c.dbPool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		queries := db.New(tx)

		upsertParams := db.UpsertCartItemParams{
			Quantity:  1,
			CartID:    int64(cartID),
			Price:     int64(productPrice),
			ProductID: int64(productID),
		}
		err := queries.UpsertCartItem(ctx, upsertParams)
		if err != nil {
			return errors.Wrap(err, "unable to upsert cart item")
		}
		updateCartParams := db.UpdateCartParams{
			Price:    int64(price),
			Vat:      0,
			Discount: 0,
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

func (c CartItemDAOPostgres) GetCartItemByID(ctx context.Context, ID int) (model.CartItem, error) {
	var foundCartItem model.CartItem
	resErr := c.dbPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		dbCartItem, err := db.New(conn).GetCartItemByID(ctx, int64(ID))
		if err != nil {
			return errors.Wrap(err, "unable to get cart item by id")
		}
		foundCartItem = createCartItemModelFromDbModel(dbCartItem)
		return nil
	})
	return foundCartItem, resErr
}

func (c CartItemDAOPostgres) UpdateCartItem(ctx context.Context, ID int, quantity int, discount int, updatedItemPrice int, cartPrice int, cartID int) error {
	resErr := c.dbPool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		queries := db.New(tx)
		err := queries.UpdateCartItem(ctx, db.UpdateCartItemParams{
			Quantity: int64(quantity),
			Discount: int64(discount),
			Price:    int64(updatedItemPrice),
			ID:       int64(ID),
		})
		if err != nil {
			return errors.Wrap(err, "unable to update cart item")
		}

		updateCartParams := db.UpdateCartParams{
			Price:    int64(cartPrice),
			Vat:      0,
			Discount: 0,
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

func (c CartItemDAOPostgres) DeleteCartItem(ctx context.Context, ID int, cartID int, cartPrice int) error {
	resErr := c.dbPool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		queries := db.New(tx)
		err := queries.DeleteCartItem(ctx, int64(ID))
		if err != nil {
			return errors.Wrap(err, "unable to delete cart item")
		}

		updateCartParams := db.UpdateCartParams{
			Price:    int64(cartPrice),
			Vat:      0,
			Discount: 0,
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

func createCartItemModelFromGetCartItemsByCartIDRow(dbCartItem db.GetCartItemsByCartIDRow) model.CartItem {
	return model.CartItem{
		ID:           int(dbCartItem.ID),
		Quantity:     int(dbCartItem.Quantity),
		CartID:       int(dbCartItem.CartID),
		Discount:     int(dbCartItem.Discount),
		Price:        int(dbCartItem.Price),
		ProductID:    int(dbCartItem.ProductID),
		ProductTitle: dbCartItem.ProductTitle.String,
	}
}

func createCartItemModelFromDbModel(item db.CartItem) model.CartItem {
	return model.CartItem{
		Quantity:  int(item.Quantity),
		CartID:    int(item.CartID),
		Discount:  int(item.Discount),
		Price:     int(item.Price),
		ProductID: int(item.ProductID),
	}
}
