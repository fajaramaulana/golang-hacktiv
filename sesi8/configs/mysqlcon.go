package configs

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDbMysql() (*sql.DB, error) {

	dsn := "root:root@tcp(127.0.0.1:3306)/kode-hacktv8?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
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
