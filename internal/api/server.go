package api

import (
	"basket-api/internal/persistence"
	"basket-api/internal/route"
	"basket-api/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

type Server struct {
	router *gin.Engine
	DbPool *pgxpool.Pool
}

func New() (*Server, error) {
	var server = &Server{}
	err := server.setupRouter()
	if err != nil {
		return nil, err
	}
	return server, nil

}

func (s *Server) setupRouter() error {
	//gin.Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()

	// get the database connection URL
	postgresUrl := viper.GetString("postgres.url")
	// this returns connection pool
	dbPool, err := pgxpool.Connect(context.Background(), postgresUrl)
	if err != nil {
		return err
	}

	// check app health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	//creates a new router group named basketApiGroup.
	//add all the routes that have common middlewares or the same path prefix.
	basketApiGroup := router.Group("/api")
	cartDAO := persistence.NewCartDAOPostgres(dbPool)
	cartItemDAO := persistence.NewCartItemDAOPostgres(dbPool)
	productDAO := persistence.NewProductDAOPostgres(dbPool)
	orderDAO := persistence.NewOrderDAOPostgres(dbPool)
	cartService := service.NewCartServiceImp(cartDAO, cartItemDAO, productDAO)
	orderService := service.NewOrderServiceImp(cartService, orderDAO)
	{
		route.AddProductRoutes(basketApiGroup, productDAO)
		route.AddCartRoutes(basketApiGroup, cartService)
		route.AddOrderRoutes(basketApiGroup, orderService)
	}

	s.router = router
	s.DbPool = dbPool
	return nil

}

//Start calls the ginEngine.Run methods,
//Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
