package http

import (
	"fmt"

	"github.com/isutare412/goarch/gateway/pkg/pkgerr"
)

type errorResponse struct {
	Msg string `json:"msg"`
}

type pathParameters struct {
	Nickname string `uri:"nickname"`
}

func (p pathParameters) checkNickname() error {
	if p.Nickname == "" {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEmptyField,
			Simple: fmt.Errorf("nickname path parameter is mandatory"),
		}
	}
	return nil
}
