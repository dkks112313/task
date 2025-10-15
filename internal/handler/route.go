package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"task/internal/repository"
	"task/models"
	"time"
)

func HandlerEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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

		repository.SelectFromEvents(sqlQuery, args)
	case http.MethodPost:
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			var event models.Event
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&event)
			if err != nil {
				log.Println("Error, when decode json")
				return
			}

			err = repository.InsertIntoEvents(event)
			if err != nil {
				log.Println("Error, insert data in event table")
				return
			}

			log.Println("Success insert into event table")
		} else {
			log.Fatalln("Unсorrect content type")
		}
	}
}
