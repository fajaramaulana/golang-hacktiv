package routes

import (
	"sesi6/ginfw/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/cars", controllers.GetAllCars)
	r.GET("/cars/:id", controllers.GetCarById)
	r.POST("/cars", controllers.CreateCar)

	r.PUT("/cars/:id", controllers.UpdateCar)
	r.DELETE(("cars/:id"), controllers.DeleteCar)

	return r
}
