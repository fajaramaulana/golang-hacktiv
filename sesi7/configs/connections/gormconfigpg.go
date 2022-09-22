package connections

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionGormPg() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file1")
	}

	pqsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST_PQSQL"), os.Getenv("PORT_PQSQL"), os.Getenv("USER_PQSQL"), os.Getenv("PASS_PQSQL"), os.Getenv("DBNAME_PQSQL"))

	db, err := gorm.Open(postgres.Open(pqsqlInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
