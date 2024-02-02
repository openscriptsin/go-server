
# Go Server

## Overview
Go Server is a ready-to-use web server for building efficient and scalable web applications in Go. It is built on top of Gin for routing and Dig for seamless dependency injection. This project aims to provide a simple yet powerful foundation for your web development projects.

## Features
- **Fast and Lightweight:** Harness the speed of Gin and the efficiency of Dig for a performant server.
- **Easy to Get Started:** Start building your web application with minimal setup.
- **Dependency Injection:** Utilize Dig for managing dependencies and promoting clean, modular code.

## Installation
To integrate Go Server into your project, make sure you have Go installed. Run the following command to get the package:
```
go get -u github.com/openscriptsin/go-server
```
## Getting Started
ioc/kernel.go
```
package ioc

import (
	"fmt"
	"go-gin-api/src/controller"
	"go-gin-api/src/logger"

	"go.uber.org/dig"
)

var Controllers = []interface{}{
	controller.NewStatus,
}

var otherInjectable = []interface{}{
	logger.NewLogrus,
}

func NewKernal() *dig.Container {
	c := dig.New()

	for _, injectable := range otherInjectable {
		c.Provide(injectable)
	}

	return c
}

func RegisterControllers(c *dig.Container) {
	for _, controller := range Controllers {
		c.Invoke(controller)
	}
}

```
main.go
```
package main

import (
	"go-gin-api/src/ioc"

	goServer "github.com/openscriptsin/go-server"
)

func main() {

	kernal := ioc.NewKernal()
	ginServer := goServer.New(kernal, ioc.RegisterControllers)

	ginServer.Start(3002)

}
```

## Dependencies
- [Dig]
- [Gin]

## Example 
- [go-gin-poc]

## License
MIT

## Acknowledgments
- Hat tip to the contributors of Gin and Dig.
- Inspiration from various open-source projects.
- etc.

[//]: # (Links used in above documents)
[Dig]: <https://pkg.go.dev/go.uber.org/dig>
[Gin]: <https://pkg.go.dev/github.com/gin-gonic/gin>
[go-gin-poc]: <https://github.com/amiransari27/go-gin-poc>