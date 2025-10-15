package service

import (
	"task/internal/repository"
	"time"
)

func BackgroundTaskSaveToDatabase() {
	start_time := time.Now()
	ticker := time.NewTicker(4 * time.Hour)

	go func() {
		for range ticker.C {
			repository.InsertIntoUserStats(&start_time)
		}
	}()
}
