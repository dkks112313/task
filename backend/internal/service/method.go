package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dkks112313/task/internal/repository"
	"github.com/dkks112313/task/models"
)

type EventService interface {
	MethodGetForMainRoute(w http.ResponseWriter, r *http.Request)
	MethodPostForMainRoute(w http.ResponseWriter, r *http.Request)
}

type eventService struct {
	repo repository.EventRepository
}

func InitServiceEvents(repo repository.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) MethodGetForMainRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
		if t, err := time.Parse(time.RFC3339, from); err == nil {
			t = t.UTC()
			args = append(args, t)
			conditions = append(conditions, fmt.Sprintf("time_event >= $%d", len(args)))
		}
	}

	if to := query.Get("to"); to != "" {
		if t, err := time.Parse(time.RFC3339, to); err == nil {
			t = t.UTC()
			args = append(args, t)
			conditions = append(conditions, fmt.Sprintf("time_event <= $%d", len(args)))
		}
	}

	sqlQuery := "SELECT id, user_id, action, metadata, time_event FROM events"
	if len(conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	err := s.repo.SelectFromEvents(w, sqlQuery, args)
	if err != nil {
		http.Error(w, "Uncorrect filter data", http.StatusBadRequest)
	}
}

func (s *eventService) MethodPostForMainRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		var event models.Event
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&event)
		if err != nil {
			http.Error(w, "Uncorrect decode json", http.StatusBadRequest)
			log.Println("Error, when decode json")
			return
		}

		err = s.repo.InsertIntoEvents(event)
		if err != nil {
			log.Println("Error, insert data in event table")
			http.Error(w, "Uncorrect insert data in event table", http.StatusBadRequest)
			return
		}

		log.Println("Success insert into event table")
	} else {
		log.Fatalln("Uncorrect content type")
	}
}
