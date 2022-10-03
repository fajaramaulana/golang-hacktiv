package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do something before
		fmt.Println("Before")
		next(w, r)
		// do something after
		fmt.Println("After")
	}
}

func MiddlewareGin(ctx *gin.Context) {
	fmt.Println("Before")
	ctx.Next()
	fmt.Println("After")
}
