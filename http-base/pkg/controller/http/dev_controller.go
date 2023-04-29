package http

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/goarch/http-base/pkg/log"
	"github.com/isutare412/goarch/http-base/pkg/tracing"
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
	_, span := tracing.AutoSpan(r.Context())
	defer span.End()

	log.L().Debugf("Hello!")
	responseStatus(w, r, http.StatusOK)
}

func (ctrl *devController) handlePost(w http.ResponseWriter, r *http.Request) {
	_, span := tracing.AutoSpan(r.Context())
	defer span.End()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithOperation("readHTTPBody").Errorf("Failed to read body of HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.L().Debugf("Post dev request body: %s", string(bodyBytes))
	responseBytes(w, r, bodyBytes)
}
