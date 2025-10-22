package service

import (
	"time"
)

func (s *eventService) BackgroundTaskSaveToDatabase() {
	start_time := time.Now()
	ticker := time.NewTicker(4 * time.Hour)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			s.repo.InsertIntoUserStats(&start_time)
		}
	}()
}
