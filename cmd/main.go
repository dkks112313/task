package main

import (
	"net/http"
	"task/internal/handler"
	"task/internal/repository"
	"task/internal/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	repository.InitDB()

	service.BackgroundTaskSaveToDatabase()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/events", handler.HandlerEvents)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
