/*package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metadata struct {
	Path string `json:"page"`
}

type Event struct {
	UserID    uint      `json:"user_id"`
	Action    string    `json:"action"`
	Metadata  Metadata  `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
}

var savePlace []*Event

func backgroundSaveToDatabase() {
	ticker := time.NewTicker(4 * time.Hour)

	go func() {
		for range ticker.C {
			fmt.Println("Working")
		}
	}()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			query := r.URL.Query()
			_ = query.Get("user_id")
			_ = query.Get("action")
			_ = query.Get("metadata")
			_ = query.Get("timestamp")
		}

		if r.Method == http.MethodPost {
			if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				var event Event
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&event)
				if err != nil {
					fmt.Errorf("Error, when decode json")
					return
				}

				savePlace = append(savePlace, &event)

				println(event.UserID)
				println(event.Action)
				println(event.Metadata.Path)
			} else {
				fmt.Errorf("Unkorrect content-type")
				return
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
*/

package main

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
