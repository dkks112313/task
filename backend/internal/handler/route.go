package handler

import (
	"net/http"

	"github.com/dkks112313/task/internal/service"
)

func HandlerEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		service.MethodPostForMainRoute(w, r)
	case http.MethodPost:
		service.MethodGetForMainRoute(w, r)
	}
}
