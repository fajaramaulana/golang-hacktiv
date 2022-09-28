package main

import (
	"sesi8-latihan/db"
	"sesi8-latihan/server"
	"sesi8-latihan/server/controllers"
)

// @title Orders API
// @description Sample API Spec for Orders
// @version v1.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:8081
// @contact.name Fajar
// @contact.email fajar@gmail.com
func main() {
	db := db.ConnectGorm()
	peopleController := controllers.NewPersonController(db)
	router := server.NewRouter(peopleController)
	router.Start(":8081")
}
