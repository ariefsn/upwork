package delivery

import (
	"net/http"

	"github.com/ariefsn/upwork/helper"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandlers() *chi.Mux {
	r := chi.NewRouter()
	h := &Handler{}

	r.Get("/", h.HealthHandler)
	r.Get("/health", h.HealthHandler)

	return r
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	helper.ResponseJsonSuccess(w, "running")
}
