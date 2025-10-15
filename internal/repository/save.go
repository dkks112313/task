package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

var db *sql.DB

func init() {
	db = InitDB()
}

type Metadata struct {
	Path string
}

type Event struct {
	UserID    uint
	Action    string
	Metadata  Metadata
	Timestamp time.Time
}

func main() {
	_, err := db.Query("INSERT INTO events (id, user_id, action, metadata, time_event) VALUES (1, 1, 'dds', 'brbr', '2025-10-14 12:30:00');")
	if err != nil {
		panic(err)
	}

	log.Println("All ok")
}
