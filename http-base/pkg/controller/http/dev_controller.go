package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/isutare412/goarch/http-base/pkg/log"
	"github.com/isutare412/goarch/http-base/pkg/pkgerr"
	"github.com/isutare412/goarch/http-base/pkg/tracing"
)

type devController struct{}

func (ctrl *devController) router() chi.Router {
	jsonContent := allowContentType("application/json")

	r := chi.NewRouter()
	r.Get("/", ctrl.handleGet)
	r.Post("/", jsonContent(ctrl.handlePost))

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

	var req handlePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tracing.AutoError(span, err)
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}

	log.L().Debugf("handlePostRequest: %+v", req)

	if req.Name == "foo" {
		err := fmt.Errorf("foo user leads to 500 error")
		tracing.AutoError(span, err)
		responseError(w, r, err)
		return
	} else if req.Name == "bar" {
		err := pkgerr.InvalidRequest{Reason: "bar user leads to 400 error"}
		tracing.AutoError(span, err)
		responseError(w, r, err)
		return
	}

	responseText(w, r, "Post request handled!")
}
