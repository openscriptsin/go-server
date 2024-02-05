package goServer

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type GinServer interface {
	Start(string)
	GetEngine() *gin.Engine
}

type ginServer struct {
	ginApp *gin.Engine
}

// New creates and returns a new GinServer instance.
// It requires the following parameters:
// 1. container (*dig.Container): The dependency injection container for managing dependencies.
// 2. registerController (func(*dig.Container)): A function responsible for loading API endpoints and routes by configuring the given container.
// 3. middlewareHandlers ...gin.HandlerFunc: A variable number of parameters of type gin.HandlerFunc, representing user-defined middlewares to be applied in the server's request processing pipeline.

// Example Usage:
//
//	container := dig.New()
//	server := New(container, RegisterController, LoggerMiddleware, AuthMiddleware)
//	// Use the 'server' instance to further configure and run the Gin server.
func New(
	c *dig.Container,
	registerController func(*dig.Container),
	middlewares ...gin.HandlerFunc,
) GinServer {
	// registering gin server instance to dig instance
	// since this gin server instance is getting used in all the controllers to define routes
	c.Provide(gin.New)
	var server *gin.Engine
	err := c.Invoke(func(s *gin.Engine) {
		server = s
	})

	if err != nil {
		log.Fatal(err)
	}
	server.Use(gin.Recovery(), gin.Logger())
	server.Use(middlewares...)

	// register controllers here because these are the entry point of app
	registerController(c)

	return &ginServer{ginApp: server}
}

// Start the gin server
// - port takes sting or if you want to use default port pass empty string
// Default values
// port - 8080
func (g *ginServer) Start(port string) {

	if port == "" {
		port = "8080"
	}
	g.ginApp.Run(":" + port)
}

// return the instance of *gin.Engine
func (g *ginServer) GetEngine() *gin.Engine {
	return g.ginApp
}
