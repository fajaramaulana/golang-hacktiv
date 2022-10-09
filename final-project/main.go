package main

import (
	"final-project/config/db"
)

func main() {
	_, err := db.ConnectMysqlGorm()

	if err != nil {
		panic(err)
	}

	// router := gin.Default()
}
