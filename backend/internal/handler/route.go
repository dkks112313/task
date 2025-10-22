package handler

import (
	"net/http"

	"github.com/dkks112313/task/internal/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Routes struct {
	eventService service.EventService
}

func NewRoutes(eventService service.EventService) *Routes {
	return &Routes{
		eventService: eventService,
	}
}

func (r *Routes) HandlerEvents() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("GET /events", r.eventService.MethodGetForMainRoute)
	mux.HandleFunc("POST /events", r.eventService.MethodPostForMainRoute)

	return mux
}
