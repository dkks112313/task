package service

import (
	"time"
)

func BackgroundTaskSaveToDatabase() {
	start_time := time.Now()
	ticker := time.NewTicker(4 * time.Hour)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			eventService.repo.InsertIntoUserStats(&start_time)
		}
	}()
}
