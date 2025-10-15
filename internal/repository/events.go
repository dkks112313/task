package repository

import (
	"log"
	"task/models"
	"time"
)

func SelectFromEventsByUserID() {

}

func SelectFromEventsByUPeriod() {

}

func InsertIntoEvents(event models.Event) {
	_, err := DB.Exec("INSERT INTO events (user_id, action, metadata, time_event) VALUES ($1, $2, $3, $4);",
		event.UserID, event.Action, event.Metadata.Path, time.Now())
	if err != nil {
		log.Println("Error, insert data in event table")
		return
	}

	log.Println("Success insert into event table")
}

func InsertIntoUserStats() {

}
