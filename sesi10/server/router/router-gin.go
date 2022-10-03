package router

import (
	"sesi10/server"
	"sesi10/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	user *controllers.UserController
}

func NewRouterGin(user *controllers.UserController) *Router {
	return &Router{user: user}
}

func (r *Router) StartGin(port string) {
	router := gin.Default()
	router.GET("/gin/users", server.MiddlewareGin, r.user.GetUserGin)
	router.POST("/gin/users", r.user.RegisterUser)
	router.POST("/gin/users/login", r.user.Login)
	router.Run(port)
}
