package main

import (
	"assignment3/server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8081"
	route := gin.Default()
	route.LoadHTMLGlob("web/templates/**/*")
	route.GET("/", controllers.Assignment3)
	route.Run(port)
}
