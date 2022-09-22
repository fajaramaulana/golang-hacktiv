package routes

import (
	"log"
	"sesi7/cmd/controllers"
	"sesi7/cmd/repositories"
	"sesi7/cmd/services"
	"sesi7/configs/connections"

	"github.com/gin-gonic/gin"
)

func Init() {
	App := gin.Default()
	db, err := connections.ConnectionGormPg()

	if err != nil {
		log.Fatal(err.Error())
	}

	employeeRepository := repositories.NewRepositoryEmployee(db)
	employeeService := services.EmployeeNewService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	v1 := App.Group("/api/v1")

	employeeController.EmployeeRoutes(v1)

	App.Run(":8081")
}
