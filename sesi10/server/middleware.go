package server

import (
	"fmt"
	"net/http"
	"sesi10/server/helper"
	"sesi10/server/models"
	"strings"

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

func CheckAuth(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization")
	tokenArr := strings.Split(tokenHeader, "Bearer ")

	if len(tokenArr) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	tokenStr := tokenArr[1]

	payload, err := helper.ValidateToken(tokenStr)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.Set("email", payload["email"])
	// ctx.Set("role", payload["role"])
	ctx.Next()
}

func AdminRole(ctx *gin.Context) {
	user := models.FindByEmail(ctx.GetString("email"))

	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	isCanAccess := checkRole(user.Role, []string{"admin"})
	if !isCanAccess {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "forbidden",
		})
		return
	}
	ctx.Next()
}

func checkRole(userRole string, needs []string) bool {
	for _, role := range needs {
		if userRole == role {
			return true
		}
	}
	return false
}
