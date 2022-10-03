package controllers

import (
	"encoding/json"
	"net/http"
	"sesi10/server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUserGin(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"payload": models.Users,
	})
}

func (u *UserController) RegisterUser(ctx *gin.Context) {
	var req models.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Password = string(hash)
	req.ID = len(models.Users) + 1
	models.Users = append(models.Users, req)

	ctx.JSON(http.StatusCreated, req)
}

func (u *UserController) Login(ctx *gin.Context) {
	var req models.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.FindByEmail(req.Email)

	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"payload": user,
	})
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": "Hello World",
	})
}
