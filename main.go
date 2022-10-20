package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jskaza/wedding-site/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./views/*.html")
	router.Static("./assets/css", "./assets/css")
	router.Static("./assets/img", "./assets/img")

	public := router.Group("/")
	// private := router.Group("/")
	routes.PublicRoutes(public)
	routes.PrivateRoutes(public)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
