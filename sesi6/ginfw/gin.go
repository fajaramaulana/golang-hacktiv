package ginfw

import (
	"sesi6/ginfw/routes"
)

func StartGinFw() {
	r := routes.SetupRouter()
	r.Run(":8082")
}
