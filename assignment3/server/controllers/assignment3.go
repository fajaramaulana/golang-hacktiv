package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type Cuaca struct {
	Status StatusCuaca `json:"status"`
}

type StatusCuaca struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

const (
	MAX = 100
	MIN = 1
)

func Assignment3(ctx *gin.Context) {
	// byteReturn, err := os.ReadFile("data/data.json")

	Water := rand.Intn(MAX-MIN) + MIN
	Wind := rand.Intn(MAX-MIN) + MIN

	weatherStatus := "DEFAULT"
	backgroundColor := "#00ff00"

	if Water < 5 {
		weatherStatus = "AMAN"
		backgroundColor = "#addfad"
	} else if Water >= 6 && Water == 8 {
		weatherStatus = "SIAGA"
		backgroundColor = "#ff7518"
	} else if Water > 8 {
		weatherStatus = "BAHAYA"
		backgroundColor = "#f6222f"
	} else if Wind < 6 {
		weatherStatus = "AMAN"
		backgroundColor = "#addfad"
	} else if Wind >= 7 && Wind == 15 {
		weatherStatus = "SIAGA"
		backgroundColor = "#ff7518"
	} else if Wind > 15 {
		weatherStatus = "BAHAYA"
		backgroundColor = "#f6222f"
	}

	dataCuaca := StatusCuaca{
		Water: Water,
		Wind:  Wind,
	}

	Status := Cuaca{
		Status: dataCuaca,
	}

	jsonString, err := json.Marshal(Status)

	_, err = os.Stat("data/data.json")

	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll("data", 0755)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("data/data.json", jsonString, os.ModePerm)

		if err != nil {
			fmt.Println(err)
		}
	}

	err = ioutil.WriteFile("data/data.json", jsonString, os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}

	ctx.HTML(http.StatusOK, "pages/assignment3.html", gin.H{
		"title":  "Assignment 3",
		"status": weatherStatus,
		"water":  Water,
		"wind":   Wind,
		"year":   time.Now().Year(),
		"color":  backgroundColor,
	})
}
