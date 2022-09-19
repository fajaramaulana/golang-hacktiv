package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var cars = []Car{}

func CreateCar(c *gin.Context) {
	var car Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carId := len(cars) + 1

	newCar := Car{
		ID:    carId,
		Brand: car.Brand,
		Model: car.Model,
		Price: car.Price,
	}

	cars = append(cars, newCar)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create car",
		"data":    newCar,
		// "allData": cars,
	})
}

func UpdateCar(c *gin.Context) {
	carId := c.Param("id")
	carIdInt, err := strconv.Atoi(carId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conditiion := false
	var updateCar Car

	if err := c.ShouldBindJSON(&updateCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, v := range cars {
		if v.ID == carIdInt {
			conditiion = true
			cars[i].ID = carIdInt
			if updateCar.Brand != "" {
				cars[i].Brand = updateCar.Brand
			}
			if updateCar.Model != "" {
				cars[i].Model = updateCar.Model
			}
			if updateCar.Price != 0 {
				cars[i].Price = updateCar.Price
			}

		}
	}

	if !conditiion {
		c.JSON(http.StatusNotFound, gin.H{"error_status": "Data not found", "error_message": fmt.Errorf("Data with id %d not found", carIdInt)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update car",
		"data":    cars[carIdInt-1],
	})
}

func GetAllCars(c *gin.Context) {
	if lenCars := len(cars); lenCars == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found", "message": "Empty Cars Data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get all cars",
		"data":    cars,
	})
}

func GetCarById(c *gin.Context) {
	carId := c.Param("id")
	carIdInt, err := strconv.Atoi(carId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_status": "Your inputed data type is not number", "error_message": err.Error()})
		return
	}

	checkIsNegative := math.Signbit(float64(carIdInt))

	if checkIsNegative {
		c.JSON(http.StatusBadRequest, gin.H{"error_status": "Your inputed data is negative number", "error_message": fmt.Errorf("Your inputed data is negative number")})
		return
	}

	conditiion := false
	var car Car

	for _, v := range cars {
		if v.ID == carIdInt {
			conditiion = true
			car = v
		}
	}

	if !conditiion {
		c.JSON(http.StatusNotFound, gin.H{"error_status": "Data not found", "error_message": fmt.Errorf("Data with id %d not found", carIdInt)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get car by id",
		"data":    car,
	})
}

func DeleteCar(c *gin.Context) {
	carId := c.Param("id")
	carIdInt, err := strconv.Atoi(carId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_status": "Your inputed data type is not number", "error_message": err.Error()})
		return
	}

	checkIsNegative := math.Signbit(float64(carIdInt))

	if checkIsNegative {
		c.JSON(http.StatusBadRequest, gin.H{"error_status": "Your inputed data is negative number", "error_message": fmt.Errorf("Your inputed data is negative number")})
		return
	}

	conditiion := false

	for i, v := range cars {
		if v.ID == carIdInt {
			conditiion = true
			cars = append(cars[:i], cars[i+1:]...)
		}
	}

	if !conditiion {
		c.JSON(http.StatusNotFound, gin.H{"error_status": "Data not found", "error_message": fmt.Errorf("Data with id %d not found", carIdInt)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete car by id",
		"data":    cars,
	})
}
