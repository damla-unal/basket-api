package persistence

import (
	"basket-api/internal/dpsql"
	"basket-api/internal/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type ProductDAO interface {
	ListProducts(ctx context.Context) ([]model.Product, error)
}

type ProductDAOPostgres struct {
	dbPool *pgxpool.Pool
}

var _ ProductDAO = (*ProductDAOPostgres)(nil)

func NewProductDAOPostgres(dbPool *pgxpool.Pool) ProductDAOPostgres {
	return ProductDAOPostgres{dbPool: dbPool}
}

func (p ProductDAOPostgres) ListProducts(ctx context.Context) ([]model.Product, error) {
	products := make([]model.Product, 0)

	resErr := p.dbPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		dbProducts, err := db.New(conn).ListProducts(ctx)
		if err != nil {
			return errors.Wrap(err, "unable to list all products")
		}
		for _, dbProduct := range dbProducts {
			products = append(products, createProductModelFromDpSQLModel(dbProduct))
		}
		return nil

	})
	return products, resErr
}

func createProductModelFromDpSQLModel(dbProduct db.Product) model.Product {
	return model.Product{
		ID:    dbProduct.ID,
		Title: dbProduct.Title,
		Price: dbProduct.Price,
		Vat:   dbProduct.Vat,
	}
}
