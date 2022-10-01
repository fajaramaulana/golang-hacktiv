package main

import (
	"assignment2/config"
	"assignment2/httpserver/controllers"
	"assignment2/httpserver/repositories/gorm"
	"assignment2/httpserver/route"
	"assignment2/httpserver/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectMysqlGORM()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	itemRepo := gorm.NewItemRepository(db)
	orderRepo := gorm.NewOrderRepository(db)

	orderService := services.NewOrderService(orderRepo, itemRepo)

	orderController := controllers.NewOrderController(orderService)

	app := route.NewRouter(router, orderController)
	app.SetupRouter(":8081")
}
