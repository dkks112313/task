package service

import (
	"log"
	"task/internal/repository"
	"task/models"
	"time"
)

func BackgroundTaskSaveToDatabase() {
	start_time := time.Now()
	ticker := time.NewTicker(4 * time.Hour)

	go func() {
		for range ticker.C {
			row, err := repository.DB.Query("SELECT * FROM events WHERE time_event > $1", start_time)
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
				_, err := repository.DB.Exec("INSERT INTO user_event_stats (user_id, start_time, end_time, event_count) VALUES ($1, $2, $3, $4);", k, start_time, time.Now(), v)
				if err != nil {
					log.Println("Error, insert row into user_event_stats")
					return
				}
			}
			start_time = time.Now()
		}
	}()
}
