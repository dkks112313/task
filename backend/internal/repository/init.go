package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func readEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error, opening env file")
		return
	}
}

func InitDB() {
	readEnv()

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
	}

	if err := dbOpen.Ping(); err != nil {
		log.Fatalln("Error when checking data base connection")
	}

	DB = dbOpen
}
