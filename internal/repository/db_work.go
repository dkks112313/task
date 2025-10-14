package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx"
)

const (
	dbType   = "postgres"
	name     = "admin"
	password = "admin"
	host     = "127.0.0.1"
	port     = 3306
	dbname   = "basedb"
)

func InitDB() *sql.DB {
	dbConnect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", name, password, host, port, dbname)
	db, err := sql.Open(dbType, dbConnect)
	if err != nil {
		panic(err)
	}

	return db
}
