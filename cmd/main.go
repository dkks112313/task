package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"task/internal/repository"
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

func backgroundSaveToDatabase() {
	start_time := time.Now()
	ticker := time.NewTicker(4 * time.Hour)

	go func() {
		for range ticker.C {
			row, err := db.Query("SELECT * FROM events WHERE time_event > $1", start_time)
			if err != nil {
				log.Println("Uncorrect select data from time")
				return
			}
			defer row.Close()

			user_id_count := map[uint]int{}
			for row.Next() {
				log.Println("New row")

				var event Event
				var id int
				err := row.Scan(&id, &event.UserID, &event.Action, &event.Metadata.Path, &event.Timestamp)
				if err != nil {
					log.Println("Error, reading from event table")
					return
				}

				user_id_count[event.UserID] += 1
			}

			for k, v := range user_id_count {
				_, err := db.Exec("INSERT INTO user_event_stats (user_id, start_time, end_time, event_count) VALUES ($1, $2, $3, $4);", k, start_time, time.Now(), v)
				if err != nil {
					log.Println("Error, insert row into user_event_stats")
					return
				}
			}
			start_time = time.Now()
		}
	}()
}

var db *sql.DB

func init() {
	db = repository.InitDB()
}

func main() {
	backgroundSaveToDatabase()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			query := r.URL.Query()
			id := query.Get("user_id")
			user_id, err := strconv.Atoi(id)
			if err != nil {
				log.Println("Invalid user_id")
				return
			}

			row, err := db.Query("SELECT * FROM events WHERE user_id=$1", user_id)
			if err != nil {
				log.Println("Error, get data from event table")
				return
			}
			defer row.Close()

			for row.Next() {
				log.Println("New row")

				var event Event
				var id int
				err := row.Scan(&id, &event.UserID, &event.Action, &event.Metadata.Path, &event.Timestamp)
				if err != nil {
					log.Println("Error, reading from event table")
					return
				}

				log.Println(event)
			}
			log.Println("Success get data from event table")
		}

		if r.Method == http.MethodPost {
			if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				var event Event
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&event)
				if err != nil {
					log.Println("Error, when decode json")
					return
				}

				_, err = db.Exec("INSERT INTO events (user_id, action, metadata, time_event) VALUES ($1, $2, $3, $4);",
					event.UserID, event.Action, event.Metadata.Path, event.Timestamp)
				if err != nil {
					log.Println("Error, insert data in event table")
					return
				}

				log.Println("Success insert into event table")
			} else {
				log.Fatalln("Unсorrect content type")
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
