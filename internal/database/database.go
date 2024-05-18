package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //MySQL driver
	"github.com/jmoiron/sqlx"
)

func InitDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}
