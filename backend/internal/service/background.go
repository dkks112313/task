package service

import (
	"context"
	"log"
	"time"
)

func (s *eventService) BackgroundTaskSaveToDatabase() {
	start_time := time.Now()
	ticker := time.NewTicker(4 * time.Hour)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			err := s.repo.InsertIntoUserStats(ctx, &start_time)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}
