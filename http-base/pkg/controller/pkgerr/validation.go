package pkgerr

import "net/http"

type InvalidRequest struct {
	Reason string
}

func (v InvalidRequest) Error() string {
	return v.Reason
}

func (v InvalidRequest) StatusCode() int { return http.StatusBadRequest }
