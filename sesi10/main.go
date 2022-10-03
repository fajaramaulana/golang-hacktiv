package main

import (
	"sesi10/server/controllers"
	"sesi10/server/router"
)

func main() {
	userController := controllers.NewUserController()
	router := router.NewRouterGin(userController)
	router.StartGin(":8081")
}
