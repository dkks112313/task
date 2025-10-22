package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dkks112313/task/models"
)

type EventRepository interface {
	SelectFromEvents(ctx context.Context, w http.ResponseWriter, sqlQuery string, args []interface{}) error
	InsertIntoEvents(ctx context.Context, event models.Event) error
	InsertIntoUserStats(ctx context.Context, start_time *time.Time) error
}

type eventRepository struct {
	db *sql.DB
}

func InitRepositoryEvents() (EventRepository, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	return &eventRepository{db: db}, nil
}

func (u *eventRepository) SelectFromEvents(ctx context.Context, w http.ResponseWriter, sqlQuery string, args []interface{}) error {
	rows, err := u.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		log.Println("Error fetching events:", err)
		return err
	}
	defer rows.Close()

	var events []models.EventSend
	for rows.Next() {
		var event models.EventSend
		var id int64
		var t time.Time

		if err := rows.Scan(&id, &event.UserID, &event.Action, &event.Metadata.Path, &t); err != nil {
			log.Println("Error while scanning row:", err)
			return err
		}
		event.Timestamp = t.Format(time.RFC3339)
		events = append(events, event)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		log.Println("Error encoding json")
		return err
	}

	return nil
}

func (u *eventRepository) InsertIntoEvents(ctx context.Context, event models.Event) error {
	_, err := u.db.ExecContext(ctx, "INSERT INTO events (user_id, action, metadata, time_event) VALUES ($1, $2, $3, $4);",
		event.UserID, event.Action, event.Metadata.Path, time.Now())
	if err != nil {
		log.Println("Error, insert data in event table")
		return err
	}

	return nil
}

func (u *eventRepository) InsertIntoUserStats(ctx context.Context, start_time *time.Time) error {
	row, err := u.db.QueryContext(ctx, "SELECT * FROM events WHERE time_event >= $1", start_time)
	if err != nil {
		log.Println("Uncorrect select data from time")
		return err
	}
	defer row.Close()

	user_id_count := map[uint]int64{}
	for row.Next() {
		log.Println("New row")

		var event models.Event
		var id int64

		err := row.Scan(&id, &event.UserID, &event.Action, &event.Metadata.Path, &event.Timestamp)
		if err != nil {
			log.Println("Error, reading from event table")
			return err
		}

		user_id_count[event.UserID] += 1
	}

	end_time := time.Now()
	for k, v := range user_id_count {
		_, err := u.db.Exec("INSERT INTO user_event_stats (user_id, start_time, end_time, event_count) VALUES ($1, $2, $3, $4);", k, start_time, end_time, v)
		if err != nil {
			log.Println("Error, insert row into user_event_stats")
			return err
		}
	}

	*start_time = time.Now()

	return nil
}
