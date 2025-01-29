package main

import (
	"mundo/src/products/infraestructure/dependencies"
	"mundo/src/products/infraestructure/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	dependencies.Init()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()

}