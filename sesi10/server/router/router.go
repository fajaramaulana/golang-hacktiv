package router

import (
	"net/http"
	"sesi10/server"
	"sesi10/server/controllers"
)

func Start(port string) {
	router := http.NewServeMux()

	router.Handle("/users", server.Middleware1(controllers.NewUserController().GetUser))
	http.ListenAndServe(port, router)
}
