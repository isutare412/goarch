package http

import "net/http"

func responseStatus(w http.ResponseWriter, r *http.Request, status int) {
	if id, ok := extractTraceID(r.Context()); ok {
		w.Header().Set("X-B3-Traceid", id.String())
	}
	w.WriteHeader(status)
}

func responseBytes(w http.ResponseWriter, r *http.Request, bytes []byte) {
	if id, ok := extractTraceID(r.Context()); ok {
		w.Header().Set("X-B3-Traceid", id.String())
	}
	w.Write(bytes)
}
