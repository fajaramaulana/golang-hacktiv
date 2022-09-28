package server

import (
	"sesi8-latihan/server/controllers"

	docs "sesi8-latihan/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	people *controllers.PersonController
}

func NewRouter(people *controllers.PersonController) *Router {
	return &Router{people: people}
}

func (r *Router) Start(port string) {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.GET("/people", r.people.GetPeople)
	router.GET("/people/:id", r.people.GetPersonById)
	router.POST("/people", r.people.AddPeople)
	router.PUT("/people/:id", r.people.UpdatePerson)
	router.DELETE("/people/:id", r.people.DeletePerson)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(port)
}
