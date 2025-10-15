package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbType   = "postgres"
	name     = "user"
	password = "admin"
	host     = "postgres"
	port     = 5432
	dbname   = "basedb"
)

func InitDB() *sql.DB {
	dbConnect := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", name, password, host, port, dbname)
	db, err := sql.Open(dbType, dbConnect)
	if err != nil {
		panic(err)
	}

	return db
}
