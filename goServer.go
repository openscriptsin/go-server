package goServer

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type GinServer struct {
	GinApp *gin.Engine
}

// New returns a new GinServer instance
// it take two parameter
// *dig.Container
// registerController is a function which help to load api endpoints/route
func New(c *dig.Container, registerController func(*dig.Container)) *GinServer {
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

	// register controllers here because these are the entry point of app
	registerController(c)

	return &GinServer{GinApp: server}
}

// Start the gin server
// - port takes sting or if you want to use default port pass empty string
// Default values
// port - 8080
func (g *GinServer) Start(port string) {

	if port == "" {
		port = "8080"
	}
	g.GinApp.Run(":" + port)
}
