package webserver

import (
	"net/http"
	"sesi6/webserver/controllers"
)

func Start(port string) {
	http.HandleFunc("/users", controllers.GetUserHandler)
	http.ListenAndServe(port, nil)
}
