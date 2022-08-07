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

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	basketApiGroup := router.Group("/api")
	{
		route.AddProductRoutes(basketApiGroup, persistence.NewProductDAOPostgres(dbPool))
		route.AddCartRoutes(basketApiGroup,
			service.NewCartServiceImp(
				persistence.NewCartDAOPostgres(dbPool),
				persistence.NewCartItemDAOPostgres(dbPool),
				persistence.NewProductDAOPostgres(dbPool)))
		route.AddOrderRoutes(basketApiGroup, service.NewOrderServiceImp(service.NewCartServiceImp(persistence.NewCartDAOPostgres(dbPool),
			persistence.NewCartItemDAOPostgres(dbPool),
			persistence.NewProductDAOPostgres(dbPool)), persistence.NewOrderDAOPostgres(dbPool)))
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
