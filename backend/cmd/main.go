package main

import (
	"net/http"

	"github.com/dkks112313/task/internal/handler"
	"github.com/dkks112313/task/internal/repository"
	"github.com/dkks112313/task/internal/service"
)

func main() {
	repo := repository.InitRepositoryEvents()
	serv := service.InitServiceEvents(repo)
	hand := handler.NewRoutes(serv)

	mux := hand.HandlerEvents()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
