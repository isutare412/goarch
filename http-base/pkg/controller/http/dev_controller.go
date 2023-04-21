package http

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/goarch/http-base/pkg/log"
)

type devController struct{}

func (ctrl *devController) router() chi.Router {
	jsonContent := middleware.AllowContentType("application/json")

	r := chi.NewRouter()
	r.Get("/", ctrl.handleGet)
	r.With(jsonContent).Group(func(r chi.Router) {
		r.Post("/", ctrl.handlePost)
	})

	return r
}

func (ctrl *devController) handleGet(w http.ResponseWriter, r *http.Request) {
	log.L().Debugf("Hello!")
}

func (ctrl *devController) handlePost(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithOperation("readHTTPBody").Errorf("Failed to read body of HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.L().Debugf("POST dev request body: %s", string(bodyBytes))
	w.Write(bodyBytes)
}
