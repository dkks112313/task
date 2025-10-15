package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"task/internal/repository"
	"task/models"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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

				var event models.Event
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

			var conditions []string
			var args []interface{}

			if id := query.Get("user_id"); id != "" {
				if userID, err := strconv.Atoi(id); err == nil {
					args = append(args, userID)
					conditions = append(conditions, fmt.Sprintf("user_id=$%d", len(args)))
				}
			}

			if action := query.Get("action"); action != "" {
				args = append(args, action)
				conditions = append(conditions, fmt.Sprintf("action=$%d", len(args)))
			}

			if metadata := query.Get("metadata"); metadata != "" {
				args = append(args, metadata)
				conditions = append(conditions, fmt.Sprintf("metadata=$%d", len(args)))
			}

			if from := query.Get("from"); from != "" {
				if t, err := time.Parse("2006-01-02", from); err == nil {
					args = append(args, t)
					conditions = append(conditions, fmt.Sprintf("time_event >= $%d", len(args)))
				}
			}

			if to := query.Get("to"); to != "" {
				if t, err := time.Parse("2006-01-02", to); err == nil {
					args = append(args, t)
					conditions = append(conditions, fmt.Sprintf("time_event <= $%d", len(args)))
				}
			}

			sqlQuery := "SELECT id, user_id, action, metadata, time_event FROM events"
			if len(conditions) > 0 {
				sqlQuery += " WHERE " + strings.Join(conditions, " AND ")
			}

			rows, err := db.Query(sqlQuery, args...)
			if err != nil {
				log.Println("Error fetching events:", err)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var event models.Event
				var id int
				if err := rows.Scan(&id, &event.UserID, &event.Action, &event.Metadata.Path, &event.Timestamp); err != nil {
					log.Println("Error scanning row:", err)
					return
				}
				log.Println(event)
			}
		}

		if r.Method == http.MethodPost {
			if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				var event models.Event
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&event)
				if err != nil {
					log.Println("Error, when decode json")
					return
				}

				_, err = db.Exec("INSERT INTO events (user_id, action, metadata, time_event) VALUES ($1, $2, $3, $4);",
					event.UserID, event.Action, event.Metadata.Path, time.Now())
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
