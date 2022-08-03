package api

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func New() *Server {
	var server = &Server{}
	server.setupRouter()
	return server

}

func (s *Server) setupRouter() {
	//gin.Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	//TODO add routes
	s.router = router

}

//Start calls the ginEngine.Run methods,
//Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
