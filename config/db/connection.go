package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	// Postgres connection lib
	_ "github.com/lib/pq"
)

// InitDb ...
func InitDb() (*sql.DB, error) {
	connectionString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(10)

	return db, err
}
