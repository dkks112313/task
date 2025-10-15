package service

import (
	"time"

	"github.com/dkks112313/task/internal/repository"
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
