package connections

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func ConnectionDBPgSql() (*sql.DB, error) {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// envPath := filepath.Dir(wd) + "\\.env"
	// err = godotenv.Load(envPath)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file1")
	}

	pqsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST_PQSQL"), os.Getenv("PORT_PQSQL"), os.Getenv("USER_PQSQL"), os.Getenv("PASS_PQSQL"), os.Getenv("DBNAME_PQSQL"))

	db, err := sql.Open("postgres", pqsqlInfo)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
