package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func initDB() (*sql.DB, error) {
	dbType := os.Getenv("DB_TYPE")
	name := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dbConnect := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", name, password, host, port, dbname)
	dbOpen, err := sql.Open(dbType, dbConnect)
	if err != nil {
		log.Fatalln("Error opening db connection")
		return nil, err
	}

	if err := dbOpen.Ping(); err != nil {
		log.Fatalln("Error when checking data base connection")
		return nil, err
	}

	return dbOpen, nil
}
