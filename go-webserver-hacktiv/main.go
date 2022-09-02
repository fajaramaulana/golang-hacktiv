package main

import (
	"go-webserver-hacktiv/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var db []*service.User
	userServices := service.NewUserService(db)
	router := gin.Default()
	port := ":8080"

	router.POST("/register", userServices.RegisterAPI)
	router.GET("/getuser", userServices.GetUserAPI)
	router.GET("/getuser/:id", userServices.GetUserAPI)
	router.PUT("/updateuser/:id", userServices.UpdateUserAPI)
	router.DELETE("/deleteuser/:id", userServices.DeleteUserAPI)

	// http.HandleFunc("/register", userServices.RegisterAPI)
	// http.HandleFunc("/getalluser", userServices.GetUserAPI)
	// http.HandleFunc("/getuserbyid", userServices.GetUserAPIById)
	router.Run(port)
	// http.ListenAndServe(port, nil)
}
