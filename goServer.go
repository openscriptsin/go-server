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

// New returns a new GinServer instance
// it take n number of parameter
// fist is *dig.Container
// second is registerController func(*dig.Container),which is a function which help to load api endpoints/route
// then it allow n number of parameters of type gin.HandlerFunc, which are user define middlewares
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

	// server.Use(gin.Recovery())
	// server.Use(gin.Logger())
	// or
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
