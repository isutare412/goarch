package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/isutare412/goarch/http-base/pkg/log"
)

const (
	headerKeyTraceID = "x-trace-id"
)

type statusCodeError interface {
	error
	StatusCode() int
}

type simpleError interface {
	error
	SimpleError() string
}

type errorResponse struct {
	Message string `json:"message"`
}

func responseStatus(w http.ResponseWriter, r *http.Request, status int) {
	if id, ok := extractTraceID(r.Context()); ok && id.IsValid() {
		w.Header().Set(headerKeyTraceID, id.String())
	}
	w.WriteHeader(status)
}

func responseText(w http.ResponseWriter, r *http.Request, text string) {
	if id, ok := extractTraceID(r.Context()); ok && id.IsValid() {
		w.Header().Set(headerKeyTraceID, id.String())
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(text))
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := http.StatusInternalServerError
	msg := err.Error()

	var codeErr statusCodeError
	if errors.As(err, &codeErr) {
		statusCode = codeErr.StatusCode()
	}
	if statusCode >= http.StatusInternalServerError {
		log.WithOperation("responseError").Error(err.Error())
	}

	var simpleErr simpleError
	if errors.As(err, &simpleErr) {
		msg = simpleErr.SimpleError()
	}

	errResp := errorResponse{Message: msg}
	*r = *r.WithContext(injectErrorResponse(r.Context(), errResp))

	errBytes, err := json.Marshal(&errResp)
	if err != nil {
		panic(fmt.Errorf("marshaling http error response: %w", err))
	}

	if id, ok := extractTraceID(r.Context()); ok && id.IsValid() {
		w.Header().Set(headerKeyTraceID, id.String())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(errBytes)
}
