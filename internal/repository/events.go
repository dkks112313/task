package repository

import (
	"log"
	"time"

	"github.com/dkks112313/task/models"
)

func SelectFromEvents(sqlQuery string, args []interface{}) {
	rows, err := DB.Query(sqlQuery, args...)
	if err != nil {
		log.Println("Error fetching events:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		var id int
		if err := rows.Scan(&id, &event.UserID, &event.Action, &event.Metadata.Path, &event.Timestamp); err != nil {
			log.Println("Error while scanning row:", err)
			return
		}
		log.Println(event)
	}
}

func InsertIntoEvents(event models.Event) error {
	_, err := DB.Exec("INSERT INTO events (user_id, action, metadata, time_event) VALUES ($1, $2, $3, $4);",
		event.UserID, event.Action, event.Metadata.Path, time.Now())
	if err != nil {
		log.Println("Error, insert data in event table")
		return err
	}

	log.Println("Success insert into event table")
	return nil
}

func InsertIntoUserStats(start_time *time.Time) {
	row, err := DB.Query("SELECT * FROM events WHERE time_event > $1", start_time)
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
		_, err := DB.Exec("INSERT INTO user_event_stats (user_id, start_time, end_time, event_count) VALUES ($1, $2, $3, $4);", k, start_time, time.Now(), v)
		if err != nil {
			log.Println("Error, insert row into user_event_stats")
			return
		}
	}

	*start_time = time.Now()
}
